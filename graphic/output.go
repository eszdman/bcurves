package graphic

import (
	"bytes"
	"github.com/wcharczuk/go-chart"
	"os"
)

func Graph(x,y [][]float64){
	charts := make([]chart.Series, len(x))
	for i := 0; i<len(x);i++ {
		charts[i] = chart.ContinuousSeries{
			XValues: x[i],
			YValues: y[i],
		}
	}
	graph := chart.Chart{
		Series: charts,
		Width:384,
		Height:256,
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