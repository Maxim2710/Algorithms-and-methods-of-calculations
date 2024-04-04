package task2

import (
	"Algorithms/laboratory_work3/controller"
	"math"
)

const tolerance = 1e-6 // Погрешность

// Метод Ньютона для нахождения корня уравнения
func newtonMethod(x0 float64) float64 {
	x := x0 // Начальное значение приближения корня
	for {
		// Вычисляем следующее приближение корня по методу Ньютона
		xNext := x - controller.Function(x)/controller.Df(x)
		// Если разница между текущим и следующим приближением меньше погрешности, то возвращаем приближение как корень
		if math.Abs(xNext-x) < tolerance {
			return xNext
		}
		x = xNext // Обновляем текущее приближение
	}
}

// Найти корни уравнения на заданном интервале с заданной точностью
func FindRootsOnInterval(a, b float64) []float64 {
	roots := make(map[float64]bool) // Инициализируем карту для хранения уникальных корней
	// Проходимся по интервалу от a до b с шагом 0.01
	for x := a; x <= b; x += 0.01 {
		// Находим корень на текущем значении x методом Ньютона
		root := newtonMethod(x)
		// Если значение функции в найденном корне близко к нулю, добавляем его в карту уникальных корней
		if math.Abs(controller.Function(root)) < tolerance {
			roots[math.Round(root*1e6)/1e6] = true // Округляем корни до 6 знаков после запятой, чтобы избежать повторений
		}
	}

	var uniqueRoots []float64
	for root := range roots {
		uniqueRoots = append(uniqueRoots, root)
	}
	return uniqueRoots // Возвращаем уникальные корни
}
