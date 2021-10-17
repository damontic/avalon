package avalon

type Character struct {
	Role        string `json:"Role"`
	Description string `json:"Description"`
	IsGood      bool   `json:"IsGood"`
}

const (
	MINION_ROLE               = "Minion"
	MINION_DESCRIPTION        = "Tries to confuse everyone during quests. Sometimes by voting SUCCESS and sometimes by voting FAIL"
	MINION_IS_GOOD            = false
	LOYAL_SERVANT_ROLE        = "Loyal Servant of Arthur"
	LOYAL_SERVANT_DESCRIPTION = "Should always vote SUCCESS during quests"
	LOYAL_SERVANT_GOOD        = true
)

var MINION Character = Character{
	Role:        MINION_ROLE,
	Description: MINION_DESCRIPTION,
	IsGood:      MINION_IS_GOOD,
}
