package service

import (
	"context"
	"github.com/ory/fosite"
	"github.com/ory/fosite/compose"
)

var (
	oauth2 fosite.OAuth2Provider
)

func InitProvider(ctx context.Context, config *fosite.Config, storage interface{}, key interface{}) {
	oauth2 = compose.ComposeAllEnabled(config, storage, key)
}

func GetProvider() fosite.OAuth2Provider {
	return oauth2
}
