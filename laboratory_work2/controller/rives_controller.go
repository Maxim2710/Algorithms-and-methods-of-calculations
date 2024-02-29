package controller

import "fmt"

func PrintMatrix(matrix [][]float64) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%.2f ", matrix[i][j])
		}
		fmt.Println()
	}
}

func MatrixMultiply(matrix [][]float64, vector []float64) []float64 {
	result := make([]float64, len(matrix))
	for i := 0; i < len(matrix); i++ {
		sum := 0.0
		for j := 0; j < len(vector); j++ {
			sum += matrix[i][j] * vector[j]
		}
		result[i] = sum
	}
	return result
}

func DotProduct(a []float64, b []float64) float64 {
	result := 0.0
	for i := 0; i < len(a); i++ {
		result += a[i] * b[i]
	}
	return result
}

func SubtractVectors(a []float64, b []float64) []float64 {
	result := make([]float64, len(a))
	for i := 0; i < len(a); i++ {
		result[i] = a[i] - b[i]
	}
	return result
}

func MultiplyVectorByScalar(vector []float64, scalar float64) []float64 {
	result := make([]float64, len(vector))
	for i := 0; i < len(vector); i++ {
		result[i] = vector[i] * scalar
	}
	return result
}
