package task3

import (
	"Algorithms/laboratory_work1/controller"
	"math"
)

func AdaptiveSimpson(insideFunc func(float64) float64, a, b, n, e float64) (float64, float64) {

	summator := func(a, b, h float64) float64 {
		// Инициализация переменной result значением функции insideFunc в точке a
		result := insideFunc(a)
		// Инициализация переменной counter значением 0
		var counter int = 0

		// Цикл for, который итерируется, пока a меньше b
		for a < b {
			// Увеличение a на h
			a += h
			// Проверка условия, если counter делится на 2 без остатка
			if counter%2 == 0 {
				// Добавление к result значения функции insideFunc в точке a, умноженного на 4
				result += insideFunc(a) * 4
			} else {
				// Добавление к result значения функции insideFunc в точке a, умноженного на 2
				result += insideFunc(a) * 2
			}
			// Увеличение counter на 1
			counter += 1
		}

		// Добавление к result значения функции insideFunc в точке b
		result += insideFunc(b)
		// Возвращение приближенного значения интеграла методом Симпсона
		return result * h / 3
	}

	// Вызов функции summator для вычисления приближенного значения интеграла методом Симпсона с шагом (a + b) / n
	x := summator(a, b, (a+b)/n)
	// Вызов функции summator для вычисления приближенного значения интеграла методом Симпсона с шагом (a + b) / (n * 2)
	y := summator(a, b, (a+b)/(n*2))

	// Проверка условия, если разница между x и y больше 15 * e
	if !(math.Abs(x-y) < 15*e) {
		// Рекурсивный вызов функции AdaptiveSimpson для левой половины интервала
		x, _ = AdaptiveSimpson(insideFunc, a, (a+b)/2, n, e/2)
		// Рекурсивный вызов функции AdaptiveSimpson для правой половины интервала
		y, _ = AdaptiveSimpson(insideFunc, (a+b)/2, b, n, e/2)
		// Возвращение суммы значений x и y и нулевой оценки погрешности
		return x + y, 0.0
	} else {
		// Возвращение значения y и оценки погрешности, вычисленной с помощью функции RungeForAdaptSimpson из пакета controller
		return y, controller.RungeForAdaptSimpson(x, y)
	}
}
