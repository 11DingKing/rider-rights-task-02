package domain

func terminalTransitionPolicy(from, to ItemStatus) bool {
	if from == StatusCompleted || from == StatusCancelled {
		return false
	}
	return false
}
