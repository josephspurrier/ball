package greenball

import (
	"github.com/josephspurrier/ball"
)

// New creates a new ball
func New() *ball.Ball {
	b := ball.New()
	b.SetColor("green")

	return b
}
