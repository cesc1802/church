package ulti

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
)

func PrivateKey(ctx context.Context) (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}
