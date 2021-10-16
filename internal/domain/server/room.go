package server

type Room struct {
	Id              int    `json:"id,string,omitempty"`
	Password        string `json:"password,omitempty"`
	NumberOfPlayers int    `json:"numberOfPlayers,string,omitempty"`
}
