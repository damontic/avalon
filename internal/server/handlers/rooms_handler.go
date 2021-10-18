package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/damontic/avalon/internal/domain/server"
	"github.com/damontic/avalon/internal/server/jsend"
	"github.com/damontic/avalon/internal/server/logger"
)

type RoomsHandler struct {
	avalonLogger *logger.Logger
	avalonState  *server.State
}

func (rh *RoomsHandler) get(w http.ResponseWriter, r *http.Request) {
	rh.avalonLogger.Debug("handlers.rooms_handler.get", "start")
	var result interface{}
	var err error
	pathAfterRooms := strings.TrimPrefix(r.URL.Path, "/rooms")
	idString := strings.TrimPrefix(pathAfterRooms, "/")
	if idString != "" {
		rh.avalonLogger.Debugf("handlers.rooms_handler.get", "idString is: %s", idString)
		id, err := strconv.Atoi(idString)
		if err != nil {
			rh.avalonLogger.Errorf("handlers.rooms_handler.get", "%s", err.Error())
			return
		}
		result = rh.avalonState.Rooms[id]
	} else {
		rh.avalonLogger.Debug("handlers.rooms_handler.get", "idString is empty")
		result = rh.avalonState.Rooms
	}
	response, err := json.Marshal(jsend.NewJsendResponseSuccessData(result))
	if err != nil {
		rh.avalonLogger.Errorf("handlers.rooms_handler.get", "%s", err.Error())
		return
	}
	w.Write(response)
}

func (rh *RoomsHandler) post(w http.ResponseWriter, r *http.Request) {
	rh.avalonLogger.Debug("handlers.rooms_handler.post", "start")
	decoder := json.NewDecoder(r.Body)
	var room server.Room
	err := decoder.Decode(&room)
	if err != nil {
		rh.avalonLogger.Errorf("handlers.rooms_handler.post", "%s", err.Error())
		return
	}
	err = rh.avalonState.CreateRoom(room)
	if err != nil {
		rh.avalonLogger.Errorf("handlers.rooms_handler.post", "%s", err.Error())
		return
	}
	response, err := json.Marshal(jsend.NewJsendResponseSuccess())
	if err != nil {
		rh.avalonLogger.Errorf("handlers.rooms_handler.post", "%s", err.Error())
		return
	}
	w.Write(response)
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

func (rh *RoomsHandler) getLogger() *logger.Logger {
	return rh.avalonLogger
}

func NewRoomsHandler(logger *logger.Logger, avalonState *server.State) *RoomsHandler {
	return &RoomsHandler{
		logger,
		avalonState,
	}
}
