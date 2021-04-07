package point

type Point2D struct {
	X, Y float64
}
type MathP interface {
	Set(x, y float64)
	Get() (x, y float64)
	Mpy2d(in Point2D) Point2D
	Mpy(t float64) Point2D
}

func (in Point2D) Get() (x, y float64) {
	return in.X, in.Y
}
func (in *Point2D) Set(x, y float64) {
	in.X = x
	in.Y = y
}
func (in *Point2D) Mpy(t float64) {
	in.X *= t
	in.Y *= t
}
