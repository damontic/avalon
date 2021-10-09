package server

import (
	avalon "github.com/damontic/avalon/internal/avalon"
)

type Room struct {
	Id       string        `json:"Id"`
	Password string        `json:"Password"`
	Avalon   avalon.Avalon `json:"Avalon"`
}
