package state

import (
	"fmt"

	"github.com/damontic/avalon/internal/domain/room"
)

type State struct {
	Version        string            `json:"version"`
	Rooms          map[int]room.Room `json:"rooms"`
	MaxNumberRooms int               `json:"maxNumberRooms"`
}

func (s *State) CreateRoom(room room.Room) error {
	if _, ok := s.Rooms[room.Id]; ok {
		return fmt.Errorf("a room with id %d already exists", room.Id)
	}
	s.Rooms[room.Id] = room
	return nil
}
