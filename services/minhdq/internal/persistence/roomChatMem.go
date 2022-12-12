package persistence

import (
	"context"
	"errors"
	"minhdq/internal/model"
	"sync"
)

type RoomMem struct {
	Rooms map[string]struct {
		room *model.Room
		mu   *sync.RWMutex
	}
}

func (r *RoomMem) NewRoom(ctx context.Context, ID string) (err error) {
	if _, ok := r.Rooms[ID]; ok {
		err = errors.New("Room with " + ID + " already exist")
		return
	}

	r.Rooms[ID] = struct {
		room *model.Room
		mu   *sync.RWMutex
	}{room: &model.Room{
		ID: ID,
	}, mu: &sync.RWMutex{}}

	return
}

func (r *RoomMem) FindAll(ctx context.Context) (rooms []*model.Room, err error) {
	for _, t := range r.Rooms {
		t.mu.Lock()
		rooms = append(rooms, t.room)
		t.mu.Unlock()
	}
	return
}

func (r *RoomMem) GetOne(ctx context.Context, ID string) (room *model.Room, err error) {
	if r, ok := r.Rooms[ID]; ok {
		room = r.room
		return
	}

	err = errors.New("No room with " + ID + " can be found")
	return
}

func (r *RoomMem) GetOneHistory(ctx context.Context, ID string) (history []string, err error) {
	if r, ok := r.Rooms[ID]; ok {
		r.mu.Lock()
		history = r.room.History
		r.mu.Unlock()
		return
	}

	err = errors.New("No room with " + ID + " can be found")
	return
}

func (r *RoomMem) AddHistory(ctx context.Context, ID string, message string) (err error) {
	if r, ok := r.Rooms[ID]; ok {
		r.mu.Lock()
		r.room.History = append(r.room.History, message)
		r.mu.Unlock()
		return
	}

	err = errors.New("No room with " + ID + " can be found")
	return
}

func (r *RoomMem) AddUser(ctx context.Context, ID string, UserName string) (err error) {
	if r, ok := r.Rooms[ID]; ok {
		r.mu.Lock()
		r.room.User = append(r.room.User, UserName)
		r.mu.Unlock()
		return
	}

	err = errors.New("No room with " + ID + " can be found")
	return
}

func (r *RoomMem) RemoveUser(ctx context.Context, ID string, UserName string) (err error) {
	if r, ok := r.Rooms[ID]; ok {
		r.mu.Lock()
		for i, user := range r.room.User {
			if user == UserName {
				r.room.User = append(r.room.User[:i], r.room.User[i+1:]...)
				break
			}
		}
		r.mu.Unlock()
		return
	}

	err = errors.New("No room with " + ID + " can be found")
	return
}

func newRoomRepoMem(ctx context.Context) (repo *RoomMem, err error) {
	return &RoomMem{map[string]struct {
		room *model.Room
		mu   *sync.RWMutex
	}{}}, nil
}
