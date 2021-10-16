package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/damontic/avalon/internal/server/handlers"
)

type AvalonServer struct {
	port                 int
	isSsl                bool
	isHttpToHttpsEnabled bool
	sslCert              string
	sslKey               string
	domain               string
	handlers             map[string]http.Handler
}

func NewAvalonServer(maxNumberRooms int, port int, isSsl bool, domain string, sslCert string, sslKey string, isHttpToHttpsEnabled bool) *AvalonServer {
	roomHandler := handlers.NewRoomsHandler()
	handlersMap := map[string]http.Handler{
		"/rooms/": handlers.DispatchHttpMethod(roomHandler),
		"/rooms":  handlers.DispatchHttpMethod(roomHandler),
	}
	return &AvalonServer{
		port:                 port,
		isSsl:                isSsl,
		isHttpToHttpsEnabled: isHttpToHttpsEnabled,
		sslCert:              sslCert,
		sslKey:               sslKey,
		domain:               domain,
		handlers:             handlersMap,
	}
}

func (s AvalonServer) Run() {
	mux := http.NewServeMux()
	staticHandler := http.FileServer(http.Dir("./assets"))
	mux.Handle("/", staticHandler)

	for key, value := range s.handlers {
		mux.Handle(key, value)
	}

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

func (s AvalonServer) redirectTLS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(
		w,
		r,
		fmt.Sprintf("https://%s%s", s.domain, r.RequestURI),
		http.StatusMovedPermanently,
	)
}
