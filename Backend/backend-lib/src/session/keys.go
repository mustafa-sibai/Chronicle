package session

import "time"

const (
	sessionExchangeCodeKeyPrefix = "exchange:"
	sessionKeyPrefix             = "session:"

	SessionExchangeCodeTTL = 60 * time.Second
	SessionTTL             = 24 * time.Hour
)

// SessionExchangeCodeKey is the Valkey key holding a single-use session exchange code -> accountId.
func SessionExchangeCodeKey(code string) string {
	return sessionExchangeCodeKeyPrefix + code
}

// SessionKey is the Valkey key holding a session id -> accountId.
func SessionKey(id string) string {
	return sessionKeyPrefix + id
}
