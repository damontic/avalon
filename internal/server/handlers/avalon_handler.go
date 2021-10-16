package handlers

import (
	"fmt"
	"net/http"
)

type AvalonHandler interface {
	get(w http.ResponseWriter, r *http.Request)
	post(w http.ResponseWriter, r *http.Request)
	put(w http.ResponseWriter, r *http.Request)
	delete(w http.ResponseWriter, r *http.Request)
}

func DispatchHttpMethod(avalonHandler AvalonHandler) http.Handler {
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
			fmt.Fprintf(w, "This verb is not supported, %s", r.Method)
		}
	})
}
