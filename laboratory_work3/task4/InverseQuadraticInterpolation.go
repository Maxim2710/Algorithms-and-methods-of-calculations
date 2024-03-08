package task4

import (
	"Algorithms/laboratory_work3/controller"
	"math"
)

// Метод обратной квадратичной интерполяции для поиска минимума функции
func InverseQuadraticInterpolation(minX, maxX, epsilon float64) float64 {
	// Выполняем цикл до тех пор, пока разница между значениями maxX и minX не станет меньше epsilon
	for math.Abs(maxX-minX) >= epsilon {
		// Вычисляем значение x, используя формулу обратной квадратичной интерполяции
		x := (minX*controller.QuadraticFunction(maxX) - maxX*controller.QuadraticFunction(minX)) / (controller.QuadraticFunction(maxX) - controller.QuadraticFunction(minX))
		// Проверяем, является ли значение функции в точке x меньше, чем значения функции в точках minX и maxX
		if controller.QuadraticFunction(x) < controller.QuadraticFunction(minX) && controller.QuadraticFunction(x) < controller.QuadraticFunction(maxX) {
			// Если значение функции в точке x меньше, обновляем maxX значением x
			maxX = x
		} else {
			// Иначе обновляем minX значением x
			minX = x
		}
	}
	// Когда разница между maxX и minX станет меньше epsilon, возвращаем minX, предполагая, что это значение ближе всего к минимуму функции на заданном интервале
	return minX
}
