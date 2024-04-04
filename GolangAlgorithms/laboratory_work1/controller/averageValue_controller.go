package controller

import (
	"math"
)

// AverageValueFloat64 - функция, которая вычисляет среднеквадратичное отклонение для массива чисел типа float64
func AverageValueFloat64(a []float64) float64 {
	// Объявление переменных для хранения суммы квадратов, суммы и размера массива
	var res1, res2, sizeMath float64

	// Итерация по всем элементам массива
	for _, i := range a {
		// Добавление квадрата элемента к сумме квадратов
		res1 += i * i
		// Добавление элемента к сумме
		res2 += i
		// Увеличение размера массива на 1
		sizeMath += 1
	}

	// Возвращение среднеквадратичного отклонения
	return math.Sqrt(res1/sizeMath - (res2*res2)/(sizeMath*sizeMath))
}

// AverageValueInt - функция, которая вычисляет среднеквадратичное отклонение для массива чисел типа int
func AverageValueInt(a []int) int {
	// Объявление переменных для хранения суммы квадратов, суммы и размера массива
	var res_1, res_2, size_math int

	// Итерация по всем элементам массива
	for _, i := range a {
		// Добавление квадрата элемента к сумме квадратов
		res_1 += i * i
		// Добавление элемента к сумме
		res_2 += i
		// Увеличение размера массива на 1
		size_math += 1
	}

	// Преобразование переменных суммы квадратов, суммы и размера массива в тип float64
	var res1 float64 = float64(res_1)
	var res2 float64 = float64(res_2)
	var sizeMath float64 = float64(size_math)

	// Возвращение среднеквадратичного отклонения
	return int(math.Sqrt(res1/sizeMath - (res2*res2)/(sizeMath*sizeMath)))
}
