package coordinate

import (
	"fmt"
)

type Signal struct {
	X, Y uint16
}

// This is just a helper function, you may modified as needed
func (c Signal) String() string {
	return fmt.Sprintf("\n{X: %d, Y: %d}", c.X, c.Y)
}

type Segment struct {
	X, Y uint16
}