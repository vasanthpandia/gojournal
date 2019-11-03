package models

type AuthToken struct {
	Token string `json:"token"`
	ExpiresAt int64 `json:"expiresAt"`
}
