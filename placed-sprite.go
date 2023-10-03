package gani

import (
	"fmt"
	"strings"
)

// PlacedSprite defines a Sprite that has been placed in a Frame
type PlacedSprite struct {
	X, Y        int
	SpriteIndex int
}

func (p *PlacedSprite) String() string {
	return fmt.Sprintf("%4d %3d %3d", p.X, p.Y, p.SpriteIndex)
}

// NewPlacedSprite creates a new PlacedSprite at the given (x, y) and given Sprite.Index
func NewPlacedSprite(x, y, spriteIndex int) *PlacedSprite {
	return &PlacedSprite{
		X:           x,
		Y:           y,
		SpriteIndex: spriteIndex,
	}
}

// PlacedSpriteSliceToString strings.Builder the slice into an output format
func PlacedSpriteSliceToString(placedSprites []*PlacedSprite) string {
	builder := strings.Builder{}

	var strs []string
	for _, placedSprite := range placedSprites {
		strs = append(strs, placedSprite.String())
	}

	builder.WriteString(strings.Join(strs, ", "))

	return builder.String()
}
