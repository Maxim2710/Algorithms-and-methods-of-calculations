package task1

import (
	"Algorithms/laboratory_work3/controller"
	"fmt"
	"math"
)

// Метод дихотомии
func DichotomyMethod(a, b, epsilon float64) float64 {
	// Проверка, что функция меняет знак на отрезке [a, b]
	if controller.Function(a)*controller.Function(b) > 0 {
		fmt.Println("Error: the function does not change the sign on this segment")
		return math.NaN() // Если функция не меняет знак, возвращаем NaN (Not-a-Number)
	}

	// Количество итераций
	n := int(math.Log((b-a)/epsilon) / math.Log(2)) // Определяем количество итераций по формуле

	for i := 0; i < n; i++ {
		// Находим середину отрезка
		c := (a + b) / 2 // Середина отрезка

		// Проверяем, меняет ли функция знак на отрезке [a, c]
		if controller.Function(a)*controller.Function(c) < 0 {
			b = c // Если меняет, сдвигаем правую границу отрезка
		} else {
			a = c // Иначе сдвигаем левую границу отрезка
		}
	}

	// Возвращаем середину отрезка как приближенное значение корня
	return (a + b) / 2 // Возвращаем середину отрезка
}
