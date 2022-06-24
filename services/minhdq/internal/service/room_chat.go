package service

import (
	"context"
	"minhdq/internal/model"
	"minhdq/internal/persistence"
)

type RoomChatIDCommand struct {
	ID string `json:"IO"`
}

type RoomChatIDPayloadCommand struct {
	ID      string `json:"IO"`
	Payload string `json:"Payload"`
}

func RoomChatGetAll(ctx context.Context) (rooms []*model.Room, err error) {
	return persistence.RoomChat().FindAll(ctx)
}

func (cmd *RoomChatIDCommand) RoomChatGetOne(ctx context.Context) (room *model.Room, err error) {
	return persistence.RoomChat().GetOne(ctx, cmd.ID)
}

func (cmd *RoomChatIDCommand) RoomChatGetHistory(ctx context.Context) (history []string, err error) {
	return persistence.RoomChat().GetOneHistory(ctx, cmd.ID)
}

func (cmd *RoomChatIDCommand) NewRoom(ctx context.Context) (err error) {
	return persistence.RoomChat().NewRoom(ctx, cmd.ID)
}

func (cmd *RoomChatIDPayloadCommand) RoomChatAddHistory(ctx context.Context) error {
	return persistence.RoomChat().AddHistory(ctx, cmd.ID, cmd.Payload)
}

func (cmd *RoomChatIDPayloadCommand) RoomChatAddUser(ctx context.Context) error {
	return persistence.RoomChat().AddUser(ctx, cmd.ID, cmd.Payload)
}

func (cmd *RoomChatIDPayloadCommand) RoomChatDeleteUser(ctx context.Context) error {
	return persistence.RoomChat().RemoveUser(ctx, cmd.ID, cmd.Payload)
}
