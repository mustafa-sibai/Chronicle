package main

import (
	"context"
	"fmt"
	"github.com/mustafa-sibai/chronicle/backend-lib/collections"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/backend-lib/database/mongodb"
	"github.com/mustafa-sibai/chronicle/backend-lib/database/valkeydb"
	"github.com/mustafa-sibai/chronicle/backend-lib/router"
	"github.com/mustafa-sibai/chronicle/character-service/handlers"
	"log"
	"net/http"
)

func main() {
	common.SetConfig(common.Config{
		ApplicationName: common.ApplicationName_CharacterService,
		EnvironmentType: common.EnvironmentType_Development,
	})

	mongodb.Connect("mongodb://localhost:27017/?replicaSet=rs0")
	valkeydb.Connect("localhost:6379")

	if err := collections.IndexCharacters(context.Background()); err != nil {
		log.Fatal(err)
	}
	if err := collections.IndexInventories(context.Background()); err != nil {
		log.Fatal(err)
	}

	r := router.NewRouter()

	r.Get("/health", handlers.HealthHandler)
	r.Post("/character/create", handlers.CreateHandler)
	r.Post("/character/list", handlers.ListHandler)
	r.Post("/character/delete", handlers.DeleteHandler)
	r.Post("/character/choose", handlers.ChooseHandler)
	r.Post("/inventory/get", handlers.GetInventoryHandler)
	r.Post("/inventory/add", handlers.AddItemHandler)
	r.Post("/quest/accept", handlers.AcceptQuestHandler)
	r.Post("/quest/progress", handlers.UpdateQuestProgressHandler)
	r.Post("/quest/complete", handlers.CompleteQuestHandler)

	fmt.Println("Starting service on :3004")
	log.Fatal(http.ListenAndServe(":3004", r))
}
