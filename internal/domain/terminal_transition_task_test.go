package domain

import "testing"

func TestTerminalCaseCannotReopen(t *testing.T) {
	item := &RightsCase{Status: StatusCompleted}
	if err := item.TransitionTo(StatusInProgress); err == nil {
		t.Fatal("completed case was reopened")
	}
	if item.Status != StatusCompleted {
		t.Fatalf("status changed after rejected transition: %s", item.Status)
	}
}
