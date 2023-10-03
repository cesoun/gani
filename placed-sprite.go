package gani

// PlacedSprite defines a Sprite that has been placed in a Frame
type PlacedSprite struct {
	X, Y int

	sprite *Sprite
}

// GetSpriteIndex returns the underlying Sprite.Index for the PlacedSprite
func (p *PlacedSprite) GetSpriteIndex() int {
	return p.sprite.Index
}

// NewPlacedSprite creates a new PlacedSprite at the given (x, y) and given Sprite
func NewPlacedSprite(x, y int, sprite *Sprite) *PlacedSprite {
	return &PlacedSprite{
		X:      x,
		Y:      y,
		sprite: sprite,
	}
}
