package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
	handlers             map[string]http.Handler
	logger               *logger.Logger
}

func NewAvalonServer(maxNumberRooms int, port int, isSsl bool, domain string, sslCert string, sslKey string, isHttpToHttpsEnabled bool, verbosity int, useUtc bool) *AvalonServer {
	avalonLogger, err := logger.New(logger.VerbosityLevel(verbosity), useUtc)
	if err != nil {
		log.Fatalln("Could not create the AvalonServer logger.")
	}
	metricsHandler := handlers.NewMetricsHandler(avalonLogger)
	roomHandler := handlers.NewRoomsHandler(avalonLogger, maxNumberRooms)
	handlersMap := map[string]http.Handler{
		"/metrics": handlers.DispatchHttpMethod(metricsHandler),
		"/rooms/":  handlers.DispatchHttpMethod(roomHandler),
		"/rooms":   handlers.DispatchHttpMethod(roomHandler),
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
