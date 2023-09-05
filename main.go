package main

import (
	"flag"
	"fmt"

	"github.com/mushoffa/non-linear-graph-segmentation/coordinate"
	"github.com/mushoffa/non-linear-graph-segmentation/plotter"
)

func main() {
	var signal string
	flag.StringVar(&signal, "s", "signal1", "Signal file name on dataset folder")
	flag.Parse()

	// Read from input file and store graph data in array
	data := read_dataset()

	var points []coordinate.Segment

	// YOUR SEGMENTATION FUNCTION OR PROCESS GOES HERE

	// Plot segmentation point against input signal
	plotter.Plot(signal, data, points)
}

// This is a helper function, you can make your own function
func read_dataset() []coordinate.Signal {
	var i,T uint16

	// This is example dataset array, you can make your custom struct
	data := make([]coordinate.Signal, T)

	fmt.Scanf("%d", &T)

	for i = 0; i < T; i++ {
		var x, y uint16
		fmt.Scanf("%d %d", &x, &y)
		
		data = append(data, coordinate.Signal{X:x, Y:y})
	}

	return data
}