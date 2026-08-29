package add

import (
	"context"
	"errors"

	"github.com/mustafa-sibai/chronicle/backend-lib/collections"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/backend-lib/game"
	"github.com/mustafa-sibai/chronicle/backend-lib/session"
	"github.com/mustafa-sibai/chronicle/backend-lib/types"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func AddItem(ctx context.Context, req AddItemRequest) *AddItemResponse {
	config := common.GetConfig()

	res := &AddItemResponse{
		BaseResponse: common.BaseResponse{
			ApplicationName: config.ApplicationName,
			EnvironmentType: config.EnvironmentType,
			ResponseCode:    common.ResponseCodes_AddItem,
		},
		StatusCode: AddItemCodes_Unknown,
	}

	req.Body.Normalize()

	if err := req.Body.Validate(); err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = AddItemCodes_InvalidInput
		res.Message = err.Error()
		return res
	}

	accountID, ok, err := session.ResolveAccountID(ctx, req.Head.SessionID)
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to add item"
		return res
	}
	if !ok {
		res.HttpCode = common.HttpCodes_Unauthorized
		res.StatusCode = AddItemCodes_Unauthorized
		res.Message = "Invalid or expired session"
		return res
	}

	accountObjectID, err := bson.ObjectIDFromHex(accountID)
	if err != nil {
		res.HttpCode = common.HttpCodes_Unauthorized
		res.StatusCode = AddItemCodes_Unauthorized
		res.Message = "Invalid or expired session"
		return res
	}

	characterObjectID, err := bson.ObjectIDFromHex(req.Body.CharacterID)
	if err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = AddItemCodes_InvalidInput
		res.Message = "characterId is invalid"
		return res
	}

	templateObjectID, err := bson.ObjectIDFromHex(req.Body.TemplateID)
	if err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = AddItemCodes_InvalidInput
		res.Message = "templateId is invalid"
		return res
	}

	err = collections.GetCharactersCollection().FindOne(ctx, bson.M{"_id": characterObjectID, "accountId": accountObjectID}).Err()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			res.HttpCode = common.HttpCodes_BadRequest
			res.StatusCode = AddItemCodes_CharacterNotFound
			res.Message = "Character not found"
			return res
		}
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to add item"
		return res
	}

	var template types.ContainerItemTemplate
	err = collections.GetItemTemplatesCollection().FindOne(ctx, bson.M{"_id": templateObjectID}).Decode(&template)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			res.HttpCode = common.HttpCodes_BadRequest
			res.StatusCode = AddItemCodes_TemplateNotFound
			res.Message = "Item template not found"
			return res
		}
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to add item"
		return res
	}

	var inventory types.Inventory
	err = collections.GetInventoriesCollection().FindOne(ctx, bson.M{"characterId": characterObjectID}).Decode(&inventory)
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to add item"
		return res
	}

	if template.ItemType == game.ItemType_Container {
		remaining := req.Body.Quantity

		// Fill free top-level bag slots first.
		for i := range inventory.Bags {
			if remaining == 0 {
				break
			}
			if inventory.Bags[i] != nil {
				continue
			}
			inventory.Bags[i] = &types.Bag{
				Item: types.Item{
					ID:            bson.NewObjectID(),
					TemplateID:    templateObjectID,
					CurrentStacks: 1,
				},
				Contents: []types.Item{},
			}
			remaining--
		}

		// Once the top-level slots are full, a bag can be stashed inside any
		// currently-empty equipped bag - it just becomes cargo there and stops
		// being usable as a container until something moves it into a real
		// slot later. A bag stops being eligible the moment it holds anything.
		for _, bag := range inventory.Bags {
			if remaining == 0 {
				break
			}
			if bag == nil || len(bag.Contents) != 0 {
				continue
			}
			bag.Contents = append(bag.Contents, types.Item{
				ID:            bson.NewObjectID(),
				TemplateID:    templateObjectID,
				CurrentStacks: 1,
			})
			remaining--
		}

		if remaining > 0 {
			res.HttpCode = common.HttpCodes_Conflict
			res.StatusCode = AddItemCodes_InventoryFull
			res.Message = "No free bag slot"
			return res
		}
	} else {
		bagCapacities, err := loadBagCapacities(ctx, inventory)
		if err != nil {
			res.HttpCode = common.HttpCodes_InternalServerError
			res.Message = "Failed to add item"
			return res
		}

		remaining := req.Body.Quantity

		// First, top up any existing stacks of this item that still have room.
		for _, bag := range inventory.Bags {
			if bag == nil || remaining == 0 {
				continue
			}
			for i := range bag.Contents {
				if remaining == 0 {
					break
				}
				item := &bag.Contents[i]
				if item.TemplateID != templateObjectID {
					continue
				}
				space := template.MaxStacks - item.CurrentStacks
				if space <= 0 {
					continue
				}
				amount := min(space, remaining)
				item.CurrentStacks += amount
				remaining -= amount
			}
		}

		// Then place whatever's left as new stacks in any free slots, splitting
		// across multiple slots (and bags) if a single stack can't hold it all.
		for _, bag := range inventory.Bags {
			if bag == nil || remaining == 0 {
				continue
			}
			capacity := bagCapacities[bag.TemplateID]
			for len(bag.Contents) < capacity && remaining > 0 {
				amount := min(template.MaxStacks, remaining)
				bag.Contents = append(bag.Contents, types.Item{
					ID:            bson.NewObjectID(),
					TemplateID:    templateObjectID,
					CurrentStacks: amount,
				})
				remaining -= amount
			}
		}

		if remaining > 0 {
			res.HttpCode = common.HttpCodes_Conflict
			res.StatusCode = AddItemCodes_InventoryFull
			res.Message = "Inventory full"
			return res
		}
	}

	if _, err := collections.GetInventoriesCollection().ReplaceOne(ctx, bson.M{"_id": inventory.ID}, inventory); err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to add item"
		return res
	}

	res.HttpCode = common.HttpCodes_OK
	res.StatusCode = AddItemCodes_ItemAddedSuccessfully
	res.Message = "Item added successfully"
	return res
}

// loadBagCapacities looks up the Capacity of every currently equipped bag,
// keyed by the bag's own TemplateID.
func loadBagCapacities(ctx context.Context, inventory types.Inventory) (map[bson.ObjectID]int, error) {
	bagTemplateIDs := make([]bson.ObjectID, 0, game.InventoryBagSlots)
	for _, bag := range inventory.Bags {
		if bag != nil {
			bagTemplateIDs = append(bagTemplateIDs, bag.TemplateID)
		}
	}

	capacities := map[bson.ObjectID]int{}
	if len(bagTemplateIDs) == 0 {
		return capacities, nil
	}

	cursor, err := collections.GetItemTemplatesCollection().Find(ctx, bson.M{"_id": bson.M{"$in": bagTemplateIDs}})
	if err != nil {
		return nil, err
	}

	var bagTemplates []types.ContainerItemTemplate
	if err := cursor.All(ctx, &bagTemplates); err != nil {
		return nil, err
	}

	for _, bagTemplate := range bagTemplates {
		capacities[bagTemplate.ID] = bagTemplate.Capacity
	}
	return capacities, nil
}
