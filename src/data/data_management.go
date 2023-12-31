package data

var (
	weeks          = make(map[string]Week)
	contractorGoal Goal
)

func addWeek(weekNumber string, week Week) {
	weeks[weekNumber] = week
}

func setContractorGoal(goal Goal) {
	contractorGoal = goal
}
