package handlers

import (
	"fmt"
	"net/http"

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
			avalonHandler.getLogger().Warnf("handlers.DispatchHttpMethod", "Received an unexpected HTTP method: %s", r.Method)
		}
	})
}

func FilterAccept(handleFunc http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("r.Header.Get: %s\n", r.Header.Get("Accept"))
		if value := r.Header.Get("Accept"); value != "application/json" && value != "*/*" {
			http.Error(w, http.StatusText(http.StatusUnsupportedMediaType), http.StatusUnsupportedMediaType)
		}
		handleFunc(w, r)
	})
}
