package gani

import (
	"fmt"
	"strconv"
	"strings"
)

// Frame defines a single Frame in a Gani
type Frame struct {
	Up         []*PlacedSprite
	Left       []*PlacedSprite
	Down       []*PlacedSprite
	Right      []*PlacedSprite
	Properties []fmt.Stringer
	//Wait      float64
	//PlaySound []*Sound

	isSingleDirection bool
}

// AppendPlacedSprites attempts to append the PlacedSprite for the given FrameDirection
func (f *Frame) AppendPlacedSprites(line string, dir FrameDirection) error {
	for _, placedSpriteStr := range strings.Split(line, ",") {
		// Parse the fields
		fields := strings.Fields(placedSpriteStr)

		x, err := strconv.Atoi(fields[0])
		if err != nil {
			return err
		}

		y, err := strconv.Atoi(fields[1])
		if err != nil {
			return err
		}

		spriteIndex, err := strconv.Atoi(fields[2])
		if err != nil {
			return err
		}

		// Append for the direction
		switch dir {
		case UP:
			f.Up = append(f.Up, NewPlacedSprite(x, y, spriteIndex))
		case LEFT:
			f.Left = append(f.Left, NewPlacedSprite(x, y, spriteIndex))
		case DOWN:
			f.Down = append(f.Down, NewPlacedSprite(x, y, spriteIndex))
		case RIGHT:
			f.Right = append(f.Right, NewPlacedSprite(x, y, spriteIndex))
		}
	}

	return nil
}

func (f *Frame) ParseWaitOrSound(line string) error {
	fields := strings.Split(line, " ")

	if strings.HasPrefix(line, "WAIT") {
		// Parse the WAIT
		duration, err := strconv.Atoi(fields[1])
		if err != nil {
			return err
		}

		f.Properties = append(f.Properties, NewWait(duration))
	} else {
		// Parse the PLAYSOUND
		file := fields[1]

		x, err := strconv.ParseFloat(fields[2], 64)
		if err != nil {
			return err
		}

		y, err := strconv.ParseFloat(fields[3], 64)
		if err != nil {
			return err
		}

		f.Properties = append(f.Properties, NewSound(file, x, y))
	}

	return nil
}

// NewFrame creates a new default Frame
func NewFrame(isSingleDirection bool) *Frame {
	return &Frame{
		Up:                make([]*PlacedSprite, 0),
		Left:              make([]*PlacedSprite, 0),
		Down:              make([]*PlacedSprite, 0),
		Right:             make([]*PlacedSprite, 0),
		Properties:        make([]fmt.Stringer, 0),
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
