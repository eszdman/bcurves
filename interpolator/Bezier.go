package interpolator

import (
	"bcurves/point"
)

func mixP(in1, in2 point.Point2D, t float64) point.Point2D {
	return point.Point2D{in1.X*(1.0-t) + in2.X*t, in1.Y*(1.0-t) + in2.Y*t}
}
func BezierInterpolate(in []point.Point2D, resolution int) []point.Point2D {
	output := make([]point.Point2D, resolution)
	temp := make([]point.Point2D, len(in))
	copy(temp, in)
	for i := 0; i < resolution; i++ {
		t := float64(i) / float64(resolution)
		for j := 1; j < len(in); j++ {
			for k := 0; k < len(in)-j; k++ {
				temp[k] = mixP(temp[k], temp[k+1], t)
			}
		}
		output[i] = temp[0]
	}
	return output
}
