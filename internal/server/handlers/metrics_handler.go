package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/damontic/avalon/internal/domain/server"
	"github.com/damontic/avalon/internal/server/jsend"
	"github.com/damontic/avalon/internal/server/logger"
)

type MetricsHandler struct {
	avalonLogger *logger.Logger
	avalonState  *server.State
}

func (mh *MetricsHandler) get(w http.ResponseWriter, r *http.Request) {
	mh.avalonLogger.Debug("handlers.metrics_handler.get", "start")
	response := jsend.NewJsendResponseSuccessData(mh.avalonState)
	responseJson, err := json.Marshal(response)
	if err != nil {
		mh.avalonLogger.Errorf("handlers.rooms_handler.post", "%s", err.Error())
		return
	}
	w.Write(responseJson)

}

func (mh *MetricsHandler) post(w http.ResponseWriter, r *http.Request) {
	mh.avalonLogger.Debug("handlers.metrics_handler.post", "Not implemented yet.")
	result := jsend.NewJsendResponseFailure("Not implemented: metrics_handler.post")
	json.NewEncoder(w).Encode(result)
}

func (mh *MetricsHandler) put(w http.ResponseWriter, r *http.Request) {
	mh.avalonLogger.Debug("handlers.metrics_handler.put", "Not implemented yet.")
	result := jsend.NewJsendResponseFailure("Not implemented: metrics_handler.put")
	json.NewEncoder(w).Encode(result)
}

func (mh MetricsHandler) delete(w http.ResponseWriter, r *http.Request) {
	mh.avalonLogger.Debug("handlers.metrics_handler.delete", "Not implemented yet.")
	result := jsend.NewJsendResponseFailure("Not implemented: metrics_handler.delete")
	json.NewEncoder(w).Encode(result)
}

func (mh *MetricsHandler) getLogger() *logger.Logger {
	return mh.avalonLogger
}

func NewMetricsHandler(avalonLogger *logger.Logger, avalonState *server.State) *MetricsHandler {
	return &MetricsHandler{
		avalonLogger,
		avalonState,
	}
}
