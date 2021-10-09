package avalon

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type AvalonServer struct {
	Rooms          []Room
	MaxNumberRooms int
}

func NewAvalonServer() *AvalonServer {
	avalon_server := new(AvalonServer)
	avalon_server.MaxNumberRooms = 50
	return avalon_server
}

func (s AvalonServer) GetRoomsHandler() http.HandlerFunc {
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
			result := s.CreateRoom(room)
			json.NewEncoder(w).Encode(result)
		}
	}
}

func (s *AvalonServer) CreateRoom(room Room) JsendResponse {
	s.Rooms = append(s.Rooms, room)
	return JsendResponse{
		"",
		"",
		"",
		0,
	}
}

func (s AvalonServer) redirectTLS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://avalon.xn--davidmontao-beb.com"+r.RequestURI, http.StatusMovedPermanently)
}

func (s AvalonServer) Run() {
	log.Fatal(http.ListenAndServe(":80", http.HandlerFunc(s.redirectTLS)))

	mux := http.NewServeMux()
	staticHandler := http.FileServer(http.Dir("./html"))
	mux.Handle("/", staticHandler)
	mux.HandleFunc("/rooms", s.GetRoomsHandler())

	log.Fatal(
		http.ListenAndServeTLS(
			fmt.Sprintf(":%s", os.Getenv("PORT")),
			os.Getenv("CERTFILE"),
			os.Getenv("KEYFILE"),
			mux,
		),
	)
}
