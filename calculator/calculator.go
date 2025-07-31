package calculator

import (
	"errors"
	"math"
)

type Calculator struct{}

func New() *Calculator {
	return &Calculator{}
}

func (c *Calculator) Add(a, b float64) float64 {
	return a + b
}

func (c *Calculator) Subtract(a, b float64) float64 {
	return a - b
}

func (c *Calculator) Multiply(a, b float64) float64 {
	return a * b
}

func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

func (c *Calculator) Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("cannot calculate square root of negative number")
	}
	return math.Sqrt(a), nil
}

func (c *Calculator) Pow(a, b float64) float64 {
	return math.Pow(a, b)
}

func (c *Calculator) Sin(degrees float64) float64 {
	radians := degrees * math.Pi / 180
	return math.Sin(radians)
}

func (c *Calculator) Cos(degrees float64) float64 {
	radians := degrees * math.Pi / 180
	return math.Cos(radians)
}

func (c *Calculator) String() string {
	return "Basic calculator v1.0"
}
