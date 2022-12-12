package service

import (
	"github.com/ory/fosite/handler/openid"
	"github.com/ory/fosite/token/jwt"
	"time"
)

func NewSession(user string, nonce string) *openid.DefaultSession {
	return &openid.DefaultSession{
		Claims: &jwt.IDTokenClaims{
			Issuer:      "server",
			Subject:     user,
			Nonce:       nonce,
			Audience:    []string{"server"},
			ExpiresAt:   time.Now().Add(time.Hour * 6),
			IssuedAt:    time.Now(),
			RequestedAt: time.Now(),
			AuthTime:    time.Now(),
		},
		Headers: &jwt.Headers{
			Extra: make(map[string]interface{}),
		},
	}
}
