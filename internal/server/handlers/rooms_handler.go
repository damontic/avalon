package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/damontic/avalon/internal/domain/server"
	"github.com/damontic/avalon/internal/server/jsend"
	"github.com/damontic/avalon/internal/server/logger"
)

type RoomsHandler struct {
	avalonLogger   *logger.Logger
	Rooms          map[int]server.Room `json:"rooms"`
	MaxNumberRooms int                 `json:"maxNumberRooms"`
}

func (rh *RoomsHandler) get(w http.ResponseWriter, r *http.Request) {
	pathAfterRooms := strings.TrimPrefix(r.URL.Path, "/rooms")
	idString := strings.TrimPrefix(pathAfterRooms, "/")
	if idString != "" {
		rh.avalonLogger.Debugf("handlers.rooms_handler.get", "idString is %s\n", idString)
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

func (rh *RoomsHandler) put(w http.ResponseWriter, r *http.Request) {
	rh.avalonLogger.Debug("handlers.rooms_handler.put", "Not implemented yet.")
	result := jsend.NewJsendResponseFailure("Not implemented: handlers.rooms_handler.put")
	json.NewEncoder(w).Encode(result)
}

func (rh *RoomsHandler) delete(w http.ResponseWriter, r *http.Request) {
	rh.avalonLogger.Debug("handlers.rooms_handler.delete", "Not implemented yet.")
	result := jsend.NewJsendResponseFailure("Not implemented: handlers.rooms_handler.delete")
	json.NewEncoder(w).Encode(result)
}

func (rh *RoomsHandler) createRoom(room server.Room) jsend.JsendResponse {
	if _, ok := rh.Rooms[room.Id]; ok {
		message := fmt.Sprintf("A room with id %d already exists.", room.Id)
		return jsend.NewJsendResponseFailure(message)
	}
	rh.Rooms[room.Id] = room
	return jsend.NewJsendResponseSuccess()
}

func (rh *RoomsHandler) getLogger() *logger.Logger {
	return rh.avalonLogger
}

func NewRoomsHandler(logger *logger.Logger, maxNumberRooms int) *RoomsHandler {
	return &RoomsHandler{
		logger,
		make(map[int]server.Room),
		maxNumberRooms,
	}
}
