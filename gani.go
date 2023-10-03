package gani

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"strings"
)

// TODO: SCRIPT block(s)? [Anything else we are missing?]

// TODO: Read a GANI
// TODO: Write a GANI

// Gani defines the Gani file format
type Gani struct {
	Sprites  []*Sprite
	Settings *Settings
	Frames   []*Frame

	header string
}

// Parse extracts the Gani data from the incoming buffer
func (g *Gani) Parse(b []byte) error {
	scanner := bufio.NewScanner(bytes.NewReader(b))
	scanner.Scan()

	// SPRITE lines
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		sprite, err := NewSprite(scanner.Text())
		if err != nil {
			return err
		}

		g.Sprites = append(g.Sprites, sprite)
	}

	// Settings lines
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		if err := g.Settings.Parse(scanner.Text()); err != nil {
			return err
		}
	}

	// ANI - ANIEND
	for scanner.Scan() {
		if strings.EqualFold(scanner.Text(), "ANI") {
			continue
		}

		//frame := NewFrame(g.Settings.SingleDirection)

		// TODO: Nested looping for the frames.

		if strings.EqualFold(scanner.Text(), "ANIEND") {
			break
		}

		fmt.Println(scanner.Text())
	}

	return errors.New("unimplemented")
}

// NewGani creates a new empty Gani
func NewGani() *Gani {
	return &Gani{
		Sprites:  make([]*Sprite, 0),
		Settings: NewSettings(),
		Frames:   make([]*Frame, 0),
		header:   "GANI0001",
	}
}
