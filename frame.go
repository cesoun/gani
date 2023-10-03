package gani

// Frame defines a single Frame in a Gani
type Frame struct {
	Up        []PlacedSprite
	Left      []PlacedSprite
	Down      []PlacedSprite
	Right     []PlacedSprite
	Wait      float64
	PlaySound string

	isSingleDirection bool
}

// NewFrame creates a new default Frame
func NewFrame(isSingleDirection bool) *Frame {
	return &Frame{
		Up:                make([]PlacedSprite, 0),
		Left:              make([]PlacedSprite, 0),
		Down:              make([]PlacedSprite, 0),
		Right:             make([]PlacedSprite, 0),
		isSingleDirection: isSingleDirection,
	}
}

// FrameDirection defines the direction a Frame is in
type FrameDirection int

const (
	UP FrameDirection = iota
	LEFT
	DOWN
	RIGHT
)
