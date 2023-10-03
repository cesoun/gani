package gani

import (
	"fmt"
	"strconv"
)

// Sound defines a Sound found in a Gani
type Sound struct {
	File string
	X, Y float64
}

// String formats the Sound for output
func (s *Sound) String() string {
	x := strconv.FormatFloat(s.X, 'f', -1, 64)
	y := strconv.FormatFloat(s.Y, 'f', -1, 64)

	return fmt.Sprintf("PLAYSOUND %s %s %s", s.File, x, y)
}

// NewSound creates a new Sound from the given file, start and end times
func NewSound(file string, start, end float64) *Sound {
	return &Sound{
		File: file,
		X:    start,
		Y:    end,
	}
}
