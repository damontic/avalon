package server

import (
	"fmt"
)

type Room struct {
	Id              int    `json:"id,string,omitempty"`
	Password        string `json:"password,omitempty"`
	NumberOfPlayers int    `json:"numberOfPlayers,string,omitempty"`
}

type State struct {
	Rooms          map[int]Room `json:"rooms"`
	MaxNumberRooms int          `json:"maxNumberRooms"`
}

func (s *State) CreateRoom(room Room) error {
	if _, ok := s.Rooms[room.Id]; ok {
		return fmt.Errorf("a room with id %d already exists", room.Id)
	}
	s.Rooms[room.Id] = room
	return nil
}
