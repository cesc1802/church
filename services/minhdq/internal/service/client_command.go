package service

import (
	"context"
	"minhdq/internal/model"
	"minhdq/internal/persistence"
)

func FindALlClients(ctx context.Context) (clients []*model.Client, err error) {
	return persistence.Client().FindAll(ctx)
}

type ClientCommand struct {
	ID             string   `json:"id"`
	Secret         []byte   `json:"client_secret,omitempty"`
	RotatedSecrets [][]byte `json:"rotated_secrets,omitempty"`
	RedirectURIs   []string `json:"redirect_uris"`
	GrantTypes     []string `json:"grant_types"`
	ResponseTypes  []string `json:"response_types"`
	Scopes         []string `json:"scopes"`
	Audience       []string `json:"audience"`
	Public         bool     `json:"public"`
}

func (cmd *ClientCommand) RegisterClient(ctx context.Context) (secret string, err error) {
	return persistence.Client().CreateByID(ctx, cmd.ID, cmd.RedirectURIs, cmd.ResponseTypes, cmd.Scopes, cmd.GrantTypes, cmd.Public, cmd.Audience)
}

func (cmd *ClientCommand) FindClientById(ctx context.Context) (client *model.Client, err error) {
	return persistence.Client().FindOneByID(ctx, cmd.ID)
}
