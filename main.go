package main

import (
	"flag"

	avalon_server "github.com/damontic/avalon/internal/server"
)

func main() {
	port := flag.Int("port", 80, "port to bind to")
	maxNumberRooms := flag.Int("rooms", 80, "port to bind to")
	isSsl := flag.Bool("ssl", false, "specifies the max number of rooms")
	isHttpToHttpsEnabled := flag.Bool("httpsredirect", false, "enables 80 to 443 redirect")
	domain := flag.String("domain", "avalon.davidmontaño.com", "domain tied to this app")
	sslCert := flag.String("cert", "cert.pem", "domain tied to this app")
	sslKey := flag.String("key", "privkey.pem", "domain tied to this app")
	flag.Parse()

	server := avalon_server.NewAvalonServer(*maxNumberRooms, *port, *isSsl, *domain, *sslCert, *sslKey, *isHttpToHttpsEnabled)
	server.Run()
}
