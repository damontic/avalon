package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/damontic/avalon/internal/domain/server"
)

type RoomsHandler struct {
	Rooms          map[int]server.Room `json:"rooms"`
	MaxNumberRooms int                 `json:"maxNumberRooms"`
}

func (rh RoomsHandler) get(w http.ResponseWriter, r *http.Request) {
	pathAfterRooms := strings.TrimPrefix(r.URL.Path, "/rooms")
	idString := strings.TrimPrefix(pathAfterRooms, "/")
	if idString != "" {
		log.Printf("idString is %s\n", idString)
		id, err := strconv.Atoi(idString)
		if err != nil {
			log.Printf("%s", err.Error())
			return
		}
		json.NewEncoder(w).Encode(rh.Rooms[id])
		return
	}
	json.NewEncoder(w).Encode(rh.Rooms)
}

func (rh *RoomsHandler) post(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var room server.Room
	err := decoder.Decode(&room)
	if err != nil {
		panic(err)
	}
	result := rh.createRoom(room)
	log.Printf("%v\n", rh)
	json.NewEncoder(w).Encode(result)
}

func (rh RoomsHandler) put(w http.ResponseWriter, r *http.Request) {

}

func (rh RoomsHandler) delete(w http.ResponseWriter, r *http.Request) {

}

func (rh *RoomsHandler) createRoom(room server.Room) JsendResponse {
	if _, ok := rh.Rooms[room.Id]; !ok {
		return JsendResponse{
			Success: false,
			Error:   fmt.Sprintf("A room with id %d already exists.", room.Id),
			Code:    1,
		}
	}
	rh.Rooms[room.Id] = room
	return JsendResponse{
		Success: true,
	}
}

func NewRoomsHandler(maxNumberRooms int) *RoomsHandler {
	return &RoomsHandler{
		make(map[int]server.Room),
		maxNumberRooms,
	}
}
