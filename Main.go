package main

import (
	"bcurves/graphic"
	"bcurves/interpolator"
	"bcurves/point"
	"fmt"
	"github.com/klauspost/cpuid"
	"os"
)

func main() {
	println("Hello go")
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	arg := os.Args[1]
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
	fmt.Println("CPU supports SSE:", cpuid.CPU.Supports(cpuid.SSE))
	points := make([][]point.Point2D, 1)
	points[0] = make([]point.Point2D, 6)
	points[0][0].Set(0.0, 10.0)
	points[0][1].Set(0.5, 1.0)
	points[0][2].Set(2.0, 2.0)
	points[0][3].Set(3.0, 0.0)
	points[0][4].Set(4.0, 7.0)
	points[0][5].Set(2.1, 5.0)
	//graphic.Graph([][]float64{{0.0,0.5,2.0,3.0,4.0,2.1},{0.0,1.0}},[][]float64{{10.0,1.0,2.0,0.0,7.0,5.0},{0.0,1.0}})
	graphic.GraphP(points)
	graphic.GraphP([][]point.Point2D{interpolator.BezierInterpolate(points[0], 256), points[0]})
}
