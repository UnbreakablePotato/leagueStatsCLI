package leagueMath

func GoldPerMin(totalGold int, duration int) float64 {
	totalMin := float64(duration) / 60.0
	goldPerMin := float64(totalGold) / totalMin
	return goldPerMin
}
