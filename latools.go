package latools

import (
	"fmt"
	"math"
)

func Test() {
	fmt.Println("I'm Alive!")
}

// DotProduct calculates the dot product of two n-dimensional vectors
func DotProduct(a, b []float64) (sum float64) {
	sum = 0
	if len(a) != len(b) {
		fmt.Println("These vectors are not the same length")
		return
	}

	for i := range a {
		sum += a[i] * b[i]
	}

	return
}

// Length calculates the length of an n-dimensional vector
func Length(v []float64) (d float64) {
	zero := make([]float64, len(v))
	return EuclideanDistance(zero, v)
}

// EuclideanDistance calculates the distance between to n-dimensional points
func EuclideanDistance(p, q []float64) (d float64) {
	d = 0
	if len(p) != len(q) {
		fmt.Println("These vectors are not the same length")
		return
	}

	for i := range p {
		d += (q[i] - p[i]) * (q[i] - p[i])
	}

	return math.Sqrt(d)
}

// ManhattanDistance calculates the 1-norm distance between to n-dimensional points
func ManhattanDistance(p, q []float64) (d float64) {
	d = 0
	for i := range q {
		d += math.Abs(p[i] - q[i])
	}
	return
}

// Add performs simple vector addition
func Add(p, q []float64) (x []float64) {
	x = make([]float64, len(p))
	for i := range x {
		x[i] = p[i] + q[i]
	}
	return
}
