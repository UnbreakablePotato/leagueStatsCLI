package leagueMath

func CSPerMin(totalCS int, gameDuration int) float64 {
	totalMin := float64(gameDuration) / 60.0
	csPerMin := float64(totalCS) / totalMin
	return csPerMin
}
