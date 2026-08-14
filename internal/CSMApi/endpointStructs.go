package csmapi

type Champion struct {
	ChampionId              int
	ChampionIcon            string
	RecommendedPerks        []int
	RecommendedSpells       []int    // example [1, 2]
	RecommendedAbilityOrder []string // example ["Q","W","E"]
	RecommendedStartItems   []int    // should consist of a list of start item ids
	RecommendedItems        []int    // should consist of a list of item ids
	BestMatchups            []int    // should consist of a list of champions ids
	WorstMatchups           []int    // should consist of a list of champions ids
}
