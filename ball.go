package ball

import (
	"fmt"
)

// Ball is a generic ball object
type Ball struct {
	color string
}

// New creates a new ball
func New() *Ball {
	b := &Ball{
		color: "black",
	}

	return b
}

// Color returns the color of the ball
func (b *Ball) SetColor(color string) {
	b.color = color
}

// Color returns the color of the ball
func (b *Ball) Color() string {
	return fmt.Sprintf("The ball is %s.", b.color)
}
