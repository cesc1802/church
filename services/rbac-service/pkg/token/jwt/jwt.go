package jwt

import "time"

type jwt struct {
}

type Payload struct {
	userID uint64
	roleID uint64
}

type payload interface {
	UserID() uint64
	RoleID() uint64
}

type Provider interface {
	Generate(pl payload) (*Token, error)
	Inspect(token string)
	String() string
}

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  time.Time `json:"expiry"`
}
