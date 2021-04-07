package graphic

import (
	"bcurves/point"
	"bytes"
	"github.com/wcharczuk/go-chart"
	"os"
)

func Graph(x, y [][]float64) {
	charts := make([]chart.Series, len(x))
	for i := 0; i < len(x); i++ {
		charts[i] = chart.ContinuousSeries{
			XValues: x[i],
			YValues: y[i],
		}
	}
	graph := chart.Chart{
		Series: charts,
		Width:  384,
		Height: 256,
	}
	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	file, err := os.Create("output.png")
	if file == nil || err != nil {
		return
	}
	_, _ = file.Write(buffer.Bytes())
	err = file.Close()
}
func GraphCharts(charts []chart.Series) {
	graph := chart.Chart{
		Series: charts,
		Width:  384,
		Height: 256,
	}
	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	file, err := os.Create("output.png")
	if file == nil || err != nil {
		return
	}
	_, _ = file.Write(buffer.Bytes())
	err = file.Close()
}
func GraphP(points [][]point.Point2D) {
	charts := make([]chart.Series, len(points))
	for i := 0; i < len(points); i++ {
		xPoints := make([]float64, len(points[i]))
		yPoints := make([]float64, len(points[i]))
		for j := 0; j < len(points[i]); j++ {
			xPoints[j] = points[i][j].X
			yPoints[j] = points[i][j].Y
		}
		charts[i] = chart.ContinuousSeries{
			XValues: xPoints,
			YValues: yPoints,
		}
	}
	graph := chart.Chart{
		Series: charts,
		Width:  384,
		Height: 256,
	}
	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	file, err := os.Create("output.png")
	if file == nil || err != nil {
		return
	}
	_, _ = file.Write(buffer.Bytes())
	err = file.Close()
}
