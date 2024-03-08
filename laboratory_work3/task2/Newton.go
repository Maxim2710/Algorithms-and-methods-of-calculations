package task2

import (
	"Algorithms/laboratory_work3/controller"
	"math"
)

// Метод Ньютона для поиска корня уравнения f(x) = 0
func NewtonMethod(guess float64, epsilon float64) (float64, int) {
	// Начальное предположение для корня
	x := guess
	// Счетчик итераций
	iterations := 0
	// Начало бесконечного цикла
	for {
		// Вычисляем изменение x (приращение) с помощью метода Ньютона
		dx := -controller.Function(x) / controller.DerivativeFunction(x)
		// Корректируем значение x
		x += dx
		// Увеличиваем счетчик итераций
		iterations++
		// Проверяем условие остановки: если приращение меньше заданной точности, выходим из цикла
		if math.Abs(dx) < epsilon {
			break
		}
	}
	// Возвращаем найденный корень и количество итераций
	return x, iterations
}
