package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/damontic/avalon/internal/domain/room"
	"github.com/damontic/avalon/internal/domain/state"
	"github.com/damontic/avalon/internal/server/handlers"
	"github.com/damontic/avalon/internal/server/logger"
)

type AvalonServer struct {
	Port                 int    `json:"port"`
	IsSsl                bool   `json:"isSsl"`
	IsHttpToHttpsEnabled bool   `json:"isHttpToHttpsEnabled"`
	SslCert              string `json:"sslCert"`
	SslKey               string `json:"sslKey"`
	Domain               string `json:"domain"`
	handlers             map[string]http.HandlerFunc
	logger               *logger.Logger
	State                *state.State `json:"state"`
}

func NewAvalonServer(version string, maxNumberRooms int, port int, isSsl bool, domain string, sslCert string, sslKey string, isHttpToHttpsEnabled bool, verbosity int, useUtc bool) *AvalonServer {
	avalonLogger, err := logger.New(logger.VerbosityLevel(verbosity), useUtc)
	if err != nil {
		log.Fatalln("Could not create the AvalonServer logger.")
	}

	state := &state.State{
		Version:        version,
		Rooms:          make(map[int]room.Room),
		MaxNumberRooms: maxNumberRooms,
	}

	metricsHandler := handlers.DispatchHttpMethod(handlers.NewMetricsHandler(avalonLogger, state))
	roomHandler := handlers.DispatchHttpMethod(handlers.NewRoomsHandler(avalonLogger, state))
	decorators := [2]handlers.AvalonHandlerDecorator{handlers.AvalonHeaders, handlers.FilterAccept}
	handlersMap := map[string]http.HandlerFunc{
		"/metrics": metricsHandler,
		"/rooms/":  roomHandler,
		"/rooms":   roomHandler,
	}
	for key, value := range handlersMap {
		currentHandler := value
		for _, decorator := range decorators {
			currentHandler = decorator(currentHandler)
		}
		handlersMap[key] = currentHandler
	}

	return &AvalonServer{
		Port:                 port,
		IsSsl:                isSsl,
		IsHttpToHttpsEnabled: isHttpToHttpsEnabled,
		SslCert:              sslCert,
		SslKey:               sslKey,
		Domain:               domain,
		handlers:             handlersMap,
		logger:               avalonLogger,
		State:                state,
	}
}

func (s AvalonServer) Run() {
	mux := http.NewServeMux()
	staticHandler := http.FileServer(http.Dir("./assets"))
	mux.Handle("/", staticHandler)

	for key, value := range s.handlers {
		mux.Handle(key, value)
	}

	s.logger.Info("server.server.Run", "Started")

	if s.IsSsl {
		log.Fatal(
			http.ListenAndServeTLS(
				fmt.Sprintf(":%d", s.Port),
				s.SslCert,
				s.SslKey,
				mux,
			),
		)
	} else {
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.Port), mux))
	}

	if s.IsHttpToHttpsEnabled && s.Port != 80 {
		log.Fatal(http.ListenAndServe(":80", http.HandlerFunc(s.redirectTLS)))
	}
}

func (s AvalonServer) String() string {
	res, err := json.Marshal(s)
	if err != nil {
		log.Fatalln("Error Marshalling AvalonServer")
	}
	return string(res[:])
}

func (s AvalonServer) redirectTLS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(
		w,
		r,
		fmt.Sprintf("https://%s%s", s.Domain, r.RequestURI),
		http.StatusMovedPermanently,
	)
}
