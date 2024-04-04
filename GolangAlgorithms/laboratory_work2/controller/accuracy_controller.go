package controller

import (
	"Algorithms/laboratory_work2/helpers"
	"math"
)

// CheckAccuracy - проверка точности
func CheckAccuracy(X1, X2 helpers.Vector, epsilon float64) bool {
	n := len(X1)             // Получаем размер вектора X1
	for i := 0; i < n; i++ { // Начинаем цикл по элементам вектора X1
		if math.Abs(X1[i]-X2[i]) > epsilon { // Проверяем точность
			return false // Если точность не достигнута, возвращаем false
		}
	}
	return true // Если точность достигнута, возвращаем true
}
