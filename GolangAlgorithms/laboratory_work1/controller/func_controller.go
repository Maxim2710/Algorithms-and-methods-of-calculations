package controller

import (
	"math"
)

func Function(x float64) float64 {
	return math.Pow(math.Sin(x), 2) / (13 - 12*math.Cos(x))
}
