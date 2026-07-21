package leagueMath

func KillParticipation(kills int, assists int) float64 {
	totalKills := kills + assists

	killParticipation := (float64(kills) + float64(assists)) / float64(totalKills) * 100.0
	return killParticipation
}
