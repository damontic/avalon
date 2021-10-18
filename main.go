package main

import (
	"flag"
	"fmt"
	"os"

	avalon_server "github.com/damontic/avalon/internal/server"
)

func main() {
	version := "0.0.0-beta.0"

	port := flag.Int("port", 80, "port to bind to.")
	maxNumberRooms := flag.Int("rooms", 10, "port to bind to.")
	isSsl := flag.Bool("ssl", false, "specifies the max number of rooms.")
	isHttpToHttpsEnabled := flag.Bool("httpsredirect", false, "enables 80 to 443 redirect.")
	domain := flag.String("domain", "avalon.davidmontaño.com", "domain tied to this app.")
	sslCert := flag.String("cert", "cert.pem", "domain tied to this app.")
	sslKey := flag.String("key", "privkey.pem", "domain tied to this app.")
	verbosity := flag.String("verbosity", "vvvv", "verbosity level from empty up to vvvvvv.")
	utc := flag.Bool("utc", false, "either use UTC for log time or not.")
	printVersion := flag.Bool("version", false, "prints avalon version")
	flag.Parse()

	if *printVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	verbosityLevel := len(*verbosity)
	server := avalon_server.NewAvalonServer(version, *maxNumberRooms, *port, *isSsl, *domain, *sslCert, *sslKey, *isHttpToHttpsEnabled, verbosityLevel, *utc)
	server.Run()
}
