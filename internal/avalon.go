package avalon

type Avalon struct {
	NumberOfPlayers int `json:"NumberOfPlayers"`
	NumberLoyals    int `json:"NumberLoyals"`
	NumberMinions   int `json:"NumberMinions"`
}

func NewAvalon() *Avalon {
	newAvalon := new(Avalon)
	return newAvalon
}

func (a Avalon) GetNumberLoyals() int {
	numberLoyals := 0
	switch a.NumberOfPlayers {
	case 5:
		numberLoyals = 3
	case 6:
		numberLoyals = 4
	case 7:
		numberLoyals = 4
	case 8:
		numberLoyals = 5
	case 9:
		numberLoyals = 6
	case 10:
		numberLoyals = 6
	}
	return numberLoyals
}

func (a Avalon) GetNumberMinions() int {
	numberMinions := 0
	switch a.NumberOfPlayers {
	case 5:
		numberMinions = 2
	case 6:
		numberMinions = 2
	case 7:
		numberMinions = 3
	case 8:
		numberMinions = 3
	case 9:
		numberMinions = 3
	case 10:
		numberMinions = 4
	}
	return numberMinions
}
