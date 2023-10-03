package gani

import "fmt"

// Sound defines a Sound found in a Gani
type Sound struct {
	File       string
	Start, End float64
}

// String formats the Sound for output
// todo: https://pkg.go.dev/strconv#FormatFloat for more accurate formatting maybe
func (s *Sound) String() string {
	return fmt.Sprintf("PLAYSOUND %s %.3f %.3f", s.File, s.Start, s.End)
}

// NewSound creates a new Sound from the given file, start and end times
func NewSound(file string, start, end float64) *Sound {
	return &Sound{
		File:  file,
		Start: start,
		End:   end,
	}
}
