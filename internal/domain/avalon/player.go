package avalon

type Player struct {
	Nickname string `json:"Nickname"`
}

func NewPlayer(nickname string) *Player {
	p := new(Player)
	p.Nickname = nickname
	return p
}
