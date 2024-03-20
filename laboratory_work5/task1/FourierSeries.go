package task1

import (
	"gonum.org/v1/gonum/integrate" // Импортируем пакет для численного интегрирования
	"math"                         // Импортируем математический пакет
)

// Функция Simpson вычисляет интеграл функции f методом Симпсона на интервале от a до b.
func Simpson(f func(float64) float64, a, b float64) float64 {
	// Создаем пустые срезы для хранения значений x и y
	x := make([]float64, 0)
	y := make([]float64, 0)
	// Заполняем срезы значениями функции f на равномерной сетке от a до b с шагом 0.01
	for i := a; i < b; i += 0.01 {
		x = append(x, i)
		y = append(y, f(i))
	}
	// Добавляем точку b в конец сетки
	x = append(x, b)
	y = append(y, f(b))
	// Вычисляем интеграл методом Симпсона и возвращаем результат
	return integrate.Simpsons(x, y)
}

// Функция FourierSeries вычисляет коэффициенты ряда Фурье для функции g на отрезке [-l, l] с числом гармоник N.
// Затем она вычисляет значения ряда Фурье для набора точек x.
func FourierSeries(f func(float64) float64, l, N float64, x []float64) []float64 {
	// Создаем срез для хранения коэффициентов ряда Фурье
	indexList := make([]float64, 0)
	// Вычисляем коэффициент a0 (среднее значение функции)
	indexList = append(indexList, Simpson(f, -l, l)/l)
	// Вычисляем остальные коэффициенты ряда Фурье
	for i := 1.0; i < N+1; i++ {
		// Функция an(x) для вычисления коэффициентов an
		an := func(x float64) float64 {
			return f(x) * math.Cos((math.Pi*i*x)/l)
		}
		// Функция bn(x) для вычисления коэффициентов bn
		bn := func(x float64) float64 {
			return f(x) * math.Sin((math.Pi*i*x)/l)
		}
		// Вычисляем коэффициенты an и bn и добавляем их в список коэффициентов
		indexList = append(indexList, Simpson(an, -l, l)/l)
		indexList = append(indexList, Simpson(bn, -l, l)/l)
	}
	// Создаем срез для хранения значений ряда Фурье
	result := make([]float64, 0)
	// Вычисляем значения ряда Фурье для каждой точки из x
	for _, val := range x {
		res := indexList[0] / 2 // Начальное значение - среднее значение функции
		for n := 1.0; n < N+1; n++ {
			// Добавляем вклад каждой гармоники ряда Фурье
			res += indexList[2*int(n)-1]*math.Cos((math.Pi*n*val)/l) + indexList[2*int(n)]*math.Sin((math.Pi*n*val)/l)
		}
		// Добавляем значение в результирующий срез
		result = append(result, res)
	}
	// Возвращаем значения ряда Фурье для всех точек из x
	return result
}
