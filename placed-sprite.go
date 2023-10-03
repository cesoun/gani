package gani

// PlacedSprite defines a Sprite that has been placed in a Frame
type PlacedSprite struct {
	X, Y        int
	SpriteIndex int
}

// NewPlacedSprite creates a new PlacedSprite at the given (x, y) and given Sprite.Index
func NewPlacedSprite(x, y, spriteIndex int) *PlacedSprite {
	return &PlacedSprite{
		X:           x,
		Y:           y,
		SpriteIndex: spriteIndex,
	}
}
