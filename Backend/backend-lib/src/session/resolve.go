package session

import (
	"context"

	"github.com/mustafa-sibai/chronicle/backend-lib/database/valkeydb"

	"github.com/valkey-io/valkey-go"
)

// ResolveAccountID returns the accountId bound to a session id. ok is false when
// the session does not exist or has expired.
func ResolveAccountID(ctx context.Context, sessionID string) (accountID string, ok bool, err error) {
	return hgetSessionField(ctx, sessionID, "accountId")
}

// ResolveCharacterID returns the character currently being played on a session,
// if any. ok is false when the session doesn't exist or no character is
// currently selected.
func ResolveCharacterID(ctx context.Context, sessionID string) (characterID string, ok bool, err error) {
	return hgetSessionField(ctx, sessionID, "characterId")
}

func hgetSessionField(ctx context.Context, sessionID string, field string) (value string, ok bool, err error) {
	if sessionID == "" {
		return "", false, nil
	}

	valkeyClient := valkeydb.GetValkeyClient()
	res, err := valkeyClient.Do(ctx, valkeyClient.B().Hget().Key(SessionKey(sessionID)).Field(field).Build()).ToString()
	if err != nil {
		if valkey.IsValkeyNil(err) {
			return "", false, nil
		}
		return "", false, err
	}
	return res, true, nil
}
