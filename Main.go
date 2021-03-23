package main

import (
	"bcurves/graphic"
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
	fmt.Println("CPU supports SSE:",cpuid.CPU.Supports(cpuid.SSE))
	graphic.Graph([][]float64{{0.0,0.5,2.0,3.0,4.0,2.1},{0.0,1.0}},[][]float64{{10.0,1.0,2.0,0.0,7.0,5.0},{0.0,1.0}})
}