package persistence

import (
	"context"
	"minhdq/internal/model"
	"sync"
)

var (
	_roomChatRepo        RoomChatRepository
	loadRoomChatRepoOnce sync.Once
)

type RoomChatRepository interface {
	FindAll(ctx context.Context) (rooms []*model.Room, err error)
	GetOne(ctx context.Context, ID string) (room *model.Room, err error)
	NewRoom(ctx context.Context, ID string) (err error)
	GetOneHistory(ctx context.Context, ID string) (history []string, err error)
	AddHistory(ctx context.Context, ID string, message string) (err error)
	AddUser(ctx context.Context, ID string, UserName string) (err error)
	RemoveUser(ctx context.Context, ID string, UserName string) (err error)
}

func RoomChat() RoomChatRepository {
	if _roomChatRepo == nil {
		panic("persistence: _roomChatRepo Repository not initiated")
	}

	return _roomChatRepo
}

func LoadRoomChatRepoRespositoryMem(ctx context.Context) (err error) {
	loadRoomChatRepoOnce.Do(func() {
		_roomChatRepo, err = newRoomRepoMem(ctx)
	})
	return
}
