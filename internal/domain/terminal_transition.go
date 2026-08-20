package domain

// terminalTransitionPolicy governs transitions out of terminal states.
// A completed or cancelled case is final: no outgoing transition may
// reopen it. Reaching this with a terminal `from` therefore always denies
// the move (the same-state no-op is short-circuited by ValidateTransition).
func terminalTransitionPolicy(from, to ItemStatus) bool {
	if !from.IsTerminal() {
		return false
	}
	return false
}
