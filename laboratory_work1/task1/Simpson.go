package task1

import (
	"Algorithms/laboratory_work1/controller"
	"math"
)

func Simpson(a, b float64, n int) float64 {
	h := (b - a) / float64(n)                                   // исходя из теории, где h = (b - a) / n
	sum_full := controller.Function(a) + controller.Function(b) // для удобства считаем y(0), y(n); так как они без коэффицентов

	for i := 1; i < n; i += 2 {
		sum_full += 4 * controller.Function(a+float64(i)*h)
		/*
			Делаем цикл с счетчиком от 1 до n - 1 включительно,
			причемтак как начиная с y(1) коэфф. чередуются 4,2,4,2,4.....,
			то мы идем с шагом в 2 (вычисление происходит в
			нечетных точках)
		*/
	}
	for i := 2; i < n-1; i += 2 {
		sum_full += 2 * controller.Function(a+float64(i)*h)
		/*
			Аналогичный цикл верхнему. Из-за чередования коэффицентов
			начинаем с 2 и идем с шагом 2 (вычисление происходит в
			четных точках)
		*/

	}

	return sum_full * h / 3
}

func Runge(a, b float64, n int, epsilon float64) (float64, float64, int, float64) {
	h := (b - a) / float64(n)          // исходя из формулы, где h = (b - a) / n
	first_value := Simpson(a, b, n)    // результат для текущих данных
	second_value := Simpson(a, b, n*2) // результат после удвоения кол-ва разбиений

	/*сохраняем результат в отдельную переменную,
	  чтобы потом вывести его в ответ для сравнения
	*/
	itog := 0
	result := 0.0
	for math.Abs(second_value-first_value) > epsilon {
		n = n * 2
		itog = n
		h = h / 2
		first_value = second_value
		second_value = Simpson(a, b, n*2)
		/*
			Это цикл в функции Runge,
			который удваивает количество интервалов разбиения и
			уменьшает шаг h до тех пор, пока разница между двумя
			последовательными вычислениями интеграла не станет
			меньше epsilon.
		*/
		result = math.Abs(second_value-first_value) / 15
	}
	saveBeginValue := first_value

	return saveBeginValue, second_value, itog, result
}
