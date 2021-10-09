package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type AvalonServer struct {
	Rooms                []Room
	maxNumberRooms       int
	port                 int
	isSsl                bool
	isHttpToHttpsEnabled bool
	sslCert              string
	sslKey               string
	domain               string
}

func (s AvalonServer) getRoomsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			json.NewEncoder(w).Encode(s)
		}

		if r.Method == http.MethodPost {
			decoder := json.NewDecoder(r.Body)
			var room Room
			err := decoder.Decode(&room)
			if err != nil {
				panic(err)
			}
			result := s.createRoom(room)
			json.NewEncoder(w).Encode(result)
		}
	}
}

func (s *AvalonServer) createRoom(room Room) JsendResponse {
	s.Rooms = append(s.Rooms, room)
	return JsendResponse{
		"",
		"",
		"",
		0,
	}
}

func (s AvalonServer) redirectTLS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(
		w,
		r,
		fmt.Sprintf("https://%s%s", s.domain, r.RequestURI),
		http.StatusMovedPermanently,
	)
}

func NewAvalonServer(maxNumberRooms int, port int, isSsl bool, domain string, sslCert string, sslKey string, isHttpToHttpsEnabled bool) *AvalonServer {
	avalon_server := new(AvalonServer)
	avalon_server.isSsl = isSsl
	avalon_server.isHttpToHttpsEnabled = isHttpToHttpsEnabled
	avalon_server.port = port
	avalon_server.maxNumberRooms = maxNumberRooms
	return avalon_server
}

func (s AvalonServer) Run() {
	mux := http.NewServeMux()
	staticHandler := http.FileServer(http.Dir("./assets/html"))
	mux.Handle("/", staticHandler)
	mux.HandleFunc("/rooms", s.getRoomsHandler())

	if s.isSsl {
		log.Fatal(
			http.ListenAndServeTLS(
				fmt.Sprintf(":%d", s.port),
				s.sslCert,
				s.sslKey,
				mux,
			),
		)
	} else {
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.port), mux))
	}

	if s.isHttpToHttpsEnabled && s.port != 80 {
		log.Fatal(http.ListenAndServe(":80", http.HandlerFunc(s.redirectTLS)))
	}
}
