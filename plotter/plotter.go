package plotter

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/mushoffa/non-linear-graph-segmentation/coordinate"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

const(
	// Plot function will create output folder and store generated graph inside it
	outputPath = "./output"
)

func Plot(filename string, data []coordinate.Signal, segmentPoints []coordinate.Segment) {

	err := os.MkdirAll(outputPath, os.ModePerm)
	if err != nil {
		panic(err)
	}

	filePath := fmt.Sprintf("%s/segmented_%s.png",outputPath,filename)

	dataLength := len(data)
	segmentLength := len(segmentPoints)

	graphData := make(plotter.XYs, dataLength)
	graphSegment := make(plotter.XYs, segmentLength)

	for i := 0; i < dataLength; i++ {
		coordinate := data[i]
		graphData[i].X = float64(coordinate.X)
		graphData[i].Y = float64(coordinate.Y)

		if i < segmentLength {
			segment := segmentPoints[i]
			graphSegment[i].X = float64(segment.X) 
			graphSegment[i].Y = float64(segment.Y) 
		}
	}

	p := plot.New()
	p.Title.Text = "Signal"
	p.X.Label.Text = "X-Axis"
	p.Y.Label.Text = "Y-Axis"

	p.Add(plotter.NewGrid())
	line, _, err := plotter.NewLinePoints(graphData)
	if err != nil {
		log.Panic(err)
	}

	_, points, err := plotter.NewLinePoints(graphSegment)
	if err != nil {
		log.Panic(err)
	}

	line.Color = color.RGBA{B: 255, A: 255}
	points.Shape = draw.CircleGlyph{}
	points.Color = color.RGBA{R: 255, A: 255}
	p.Add(line,points)

	if err := p.Save(30*vg.Inch, 5*vg.Inch, filePath); err != nil {
		fmt.Println("Error saving plot:", err)
	}
}