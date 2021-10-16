package avalon

type LoyalServant struct {
	Character
}

func NewLoyalServant() *LoyalServant {
	ls := new(LoyalServant)
	ls.Role = "Loyal Servant of Arthur"
	ls.Description = "Should always vote SUCCESS during quests"
	ls.IsGood = true
	return ls
}
