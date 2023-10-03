package gani

import "fmt"

type Wait struct {
	Duration int
}

func (w *Wait) String() string {
	return fmt.Sprintf("%d", w.Duration)
}

// NewWait creates a new Wait property from the given duration
func NewWait(duration int) *Wait {
	return &Wait{Duration: duration}
}
