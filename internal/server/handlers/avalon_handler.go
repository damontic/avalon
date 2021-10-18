package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/damontic/avalon/internal/server/jsend"
	"github.com/damontic/avalon/internal/server/logger"
)

type AvalonHandler interface {
	get(w http.ResponseWriter, r *http.Request)
	post(w http.ResponseWriter, r *http.Request)
	put(w http.ResponseWriter, r *http.Request)
	delete(w http.ResponseWriter, r *http.Request)
	getLogger() *logger.Logger
}

func DispatchHttpMethod(avalonHandler AvalonHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			avalonHandler.get(w, r)
		case http.MethodPost:
			avalonHandler.post(w, r)
		case http.MethodPut:
			avalonHandler.put(w, r)
		case http.MethodDelete:
			avalonHandler.delete(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			message := fmt.Sprintf("Received an unexpected HTTP method: %s", r.Method)
			avalonHandler.getLogger().Warn("handlers.DispatchHttpMethod", message)
			result := jsend.NewJsendResponseFailure(message)
			json.NewEncoder(w).Encode(result)
		}
	})
}

type AvalonHandlerDecorator func(http.HandlerFunc) http.HandlerFunc

func FilterAccept(handleFunc http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if value := r.Header.Get("Accept"); value != "application/json" && value != "*/*" {
			http.Error(w, http.StatusText(http.StatusUnsupportedMediaType), http.StatusUnsupportedMediaType)
		}
		handleFunc(w, r)
	})
}

func AvalonHeaders(handleFunc http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		handleFunc(w, r)
	})
}
