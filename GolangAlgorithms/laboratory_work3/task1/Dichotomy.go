package task1

import (
	"Algorithms/laboratory_work3/controller"
	"math"
)

// Метод дихотомии для поиска корней уравнения f(x) = 0
func BisectionMethod(a, b, epsilon float64) float64 {
	// Пока разница между b и a больше или равна эпсилону
	for math.Abs(b-a) >= epsilon {
		// Находим середину отрезка
		c := (a + b) / 2
		// Если функция на концах отрезка имеет разные знаки,
		// то корень находится между ними
		if controller.Function(a)*controller.Function(c) < 0 {
			b = c // Сдвигаем правую границу к середине
		} else {
			a = c // Сдвигаем левую границу к середине
		}
	}
	// Возвращаем середину последнего интервала как приближенное значение корня
	return (a + b) / 2
}

// Функция для нахождения всех корней уравнения f(x) = 0 на заданном интервале [a, b]
func FindAllRoots(a, b, step, epsilon float64) []float64 {
	var roots []float64 // Инициализируем массив для хранения корней
	// Проходимся по интервалу от a до b с шагом step
	for x := a; x <= b; x += step {
		// Если функция на текущем отрезке имеет разные знаки на его концах,
		// то на этом отрезке существует корень
		if controller.Function(x)*controller.Function(x+step) < 0 {
			// Находим корень на текущем отрезке методом дихотомии
			root := BisectionMethod(x, x+step, epsilon)
			// Добавляем найденный корень в массив
			roots = append(roots, root)
		}
	}
	// Возвращаем массив корней
	return roots
}
