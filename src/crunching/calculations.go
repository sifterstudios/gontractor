package crunching

import data "github.com/sifterstudios/gontractor/src/data"

func GetTotalHours(weeks map[string]data.Week) float64 {
	totalHours := 0.0
	for _, week := range weeks {
		totalHours += week.NormalHours + week.OvertimeHours
	}
	return totalHours
}

func GetPercentComplete(weeks map[string]data.Week, goal data.Goal) float64 {
	totalHours := GetTotalHours(weeks)
	return totalHours / goal.TotalContractHours
}

func GetTimeLeft(weeks map[string]data.Week, goal data.Goal) (int, float64) {
	totalHours := GetTotalHours(weeks)
	hoursLeft := goal.TotalContractHours - totalHours
	daysLeft := 8
	return daysLeft, hoursLeft
}
