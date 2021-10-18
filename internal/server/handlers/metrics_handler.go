package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/damontic/avalon/internal/server/jsend"
	"github.com/damontic/avalon/internal/server/logger"
)

type MetricsHandler struct {
	avalonLogger *logger.Logger
}

func (mh *MetricsHandler) get(w http.ResponseWriter, r *http.Request) {
	mh.avalonLogger.Debug("handlers.metrics_handler.get", "Not implemented yet.")
	result := jsend.NewJsendResponseFailure("Not implemented: handlers.metrics_handler.get")
	json.NewEncoder(w).Encode(result)
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

func NewMetricsHandler(avalonLogger *logger.Logger) *MetricsHandler {
	return &MetricsHandler{
		avalonLogger,
	}
}
