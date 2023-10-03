package gani

import (
	"fmt"
	"strconv"
	"strings"
)

// Sprite defines a sprite definition in a Gani file
type Sprite struct {
	Index       int
	X, Y        int
	W, H        int
	Image       string
	Description string
}

// String formats the Sprite for output
func (s *Sprite) String() string {
	return fmt.Sprintf("SPRITE %4d %15s %4d %4d %4d %4d %s", s.Index, s.Image, s.X, s.Y, s.W, s.H, s.Description)
}

// NewSprite will attempt to parse the Sprite from the given string
func NewSprite(line string) (*Sprite, error) {
	sprite := &Sprite{}
	fields := strings.Fields(line)

	/*
		SPRITE index image x y w h [description]
		0 	   1 	 2 	   3 4 5 6 7:
	*/

	index, err := strconv.Atoi(fields[1])
	if err != nil {
		return nil, err
	}
	sprite.Index = index
	sprite.Image = fields[2]

	x, err := strconv.Atoi(fields[3])
	if err != nil {
		return nil, err
	}
	sprite.X = x

	y, err := strconv.Atoi(fields[4])
	if err != nil {
		return nil, err
	}
	sprite.Y = y

	w, err := strconv.Atoi(fields[5])
	if err != nil {
		return nil, err
	}
	sprite.W = w

	h, err := strconv.Atoi(fields[6])
	if err != nil {
		return nil, err
	}
	sprite.H = h
	sprite.Description = strings.Join(fields[7:], " ")

	return sprite, nil
}
