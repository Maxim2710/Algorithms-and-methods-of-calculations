package controller

import (
	controller1 "Algorithms/laboratory_work1/controller"
	task11 "Algorithms/laboratory_work1/task1"
	task12 "Algorithms/laboratory_work1/task2"
	task13 "Algorithms/laboratory_work1/task3"
	controller2 "Algorithms/laboratory_work2/controller"
	"Algorithms/laboratory_work2/helpers"
	task21 "Algorithms/laboratory_work2/task1"
	task22 "Algorithms/laboratory_work2/task2"
	task23 "Algorithms/laboratory_work2/task3"
	"Algorithms/laboratory_work3/controller"
	task31 "Algorithms/laboratory_work3/task1"
	task32 "Algorithms/laboratory_work3/task2"
	task33 "Algorithms/laboratory_work3/task3"
	task34 "Algorithms/laboratory_work3/task4"
	controller4 "Algorithms/laboratory_work4/controller"
	task41 "Algorithms/laboratory_work4/task1"
	task42 "Algorithms/laboratory_work4/task2"
	task43 "Algorithms/laboratory_work4/task3"
	task44 "Algorithms/laboratory_work4/task4"
	task45 "Algorithms/laboratory_work4/task5"
	"fmt"
	"log"
	"math"
)

func IODistribution(numLabarotary, numTask int) {
	if numLabarotary == 1 && numTask == 1 {

		a := 0.0
		b := 1.6
		var n int
		epsilon := 1e-4

		fmt.Print("Please, enter the number of split intervals: ")
		fmt.Scan(&n)
		fmt.Println()

		resultWithoutRunge, resultWithRunge, itog, result := task11.Runge(a, b, n, epsilon)
		fmt.Println("Result without Runge is", resultWithoutRunge)
		fmt.Println("Result with Runge is", resultWithRunge)
		fmt.Println(itog)
		fmt.Println(result)

	}

	if numLabarotary == 1 && numTask == 2 {

		a := 0.0
		b := 2.0
		var n float64
		var inaccuracy int

		fmt.Print("Please, enter the number of throws: ")
		fmt.Scan(&n)
		fmt.Print("Please, enter the number of repetitions to identify the inaccuracy: ")
		fmt.Scan(&inaccuracy)
		fmt.Println()

		result, deviation := task12.MonteKarlo(controller1.Function, a, b, n, inaccuracy)
		fmt.Println("Result is", result, "| Deviation is", deviation)
		fmt.Println()
		fmt.Println(task12.MonteKarloImage2())

	}

	if numLabarotary == 1 && numTask == 3 {

		a := 0.0
		b := 7.0
		var inaccuracy float64

		fmt.Print("Please, enter the number of repetitions to identify the inaccuracy: ")
		fmt.Scan(&inaccuracy)
		fmt.Println()

		result, deviation := task13.AdaptiveSimpson(controller1.Function, a, b, inaccuracy, 0.01)
		fmt.Println("Result is", result, "| Deviation is", deviation)

	}

	if numLabarotary == 2 && numTask == 1 {

		// Задаем матрицу A и вектор B, X
		A := helpers.Matrix{{2, -1, 0}, {10, 5, 2}, {100, -1, 3}}
		B := helpers.Vector{3, 7, 4}
		X := helpers.Vector{0, 0, 0}

		/*Создаем копии матриц, так как после преобразований внутри функции по прежним адресам
		будут измененные значения*/
		ASave := helpers.Matrix{{2, -1, 0}, {0, 5, 2}, {1, -1, 3}}
		BSave := helpers.Vector{3, 7, 4}

		// Решаем систему уравнений методом Гаусса-Жордана с выбором ведущего элемента
		X, err := task21.GaussJordan(A, B)
		if err != nil {
			log.Fatal(err)
		}

		// Проверяем, что матрица A умноженная на X дает B
		ResComputed := controller2.MultiplyMatrixVector(ASave, X)
		for i := range BSave {
			fmt.Println("A*X[", i, "] =", ResComputed[i], "B[", i, "] =", BSave[i])
		}

		fmt.Println("-----------------------")

		// Выводим результат
		fmt.Println("Result is:", X)

	}

	if numLabarotary == 2 && numTask == 2 {

		// Заданные матрица и вектор
		A := helpers.Matrix{{2, -1, 0}, {15, 5, 2}, {1, -1, 3}}
		B := helpers.Vector{3, 7, 4}

		// Начальное приближение
		X := helpers.Vector{0, 0, 0}

		// Точность
		epsilon := 0.00001

		// Максимальное количество итераций
		maxIterations := 30

		// Итерационный метод Зейделя
		X = task22.SeidelMethod(A, B, epsilon, maxIterations)

		// Проверяем, что матрица A умноженная на X дает B
		ResComputed := controller2.MultiplyMatrixVector(A, X)
		for i := range B {
			fmt.Println("A*X[", i, "] =", ResComputed[i], "B[", i, "] =", B[i])
		}

		fmt.Println("--------------------------------------------------------------------")

		// Выводим результат
		fmt.Println("Result is:", X)

	}

	if numLabarotary == 2 && numTask == 3 {
		// Матрица A
		A := [][]float64{
			{2, -1, 0},
			{1, 5, 2},
			{3, -1, 3},
		}

		// Вектор B
		B := []float64{3, 7, 4}

		// Решение системы уравнений
		X := task23.FletcherRivesMethod(A, B)

		// Вывод результатов
		fmt.Println("Matrix A:")
		controller2.PrintMatrix(A)

		fmt.Println("------------------------------")

		fmt.Println("Vector B:")
		fmt.Println(B)

		fmt.Println("------------------------------")

		fmt.Println("Solving a system of equations:")
		fmt.Println(X)
	}

	if numLabarotary == 3 && numTask == 1 {
		// Задаем начальные значения интервала [a, b]
		a := 0.0
		b := 12.0

		// Задаем размер шага и критерий сходимости (малое значение)
		step := 0.01
		epsilon := 0.0001

		// Находим все корни уравнения f(x) = 0 в заданном интервале
		roots := task31.FindAllRoots(a, b, step, epsilon)

		// Выводим результат
		if len(roots) == 0 {
			fmt.Println("В заданном интервале корни не найдены")
		} else {
			fmt.Println("Корни уравнения f(x):")
			for _, root := range roots {

				if root-1/math.Tan(root) <= 0.5 && root-1/math.Tan(root) >= -0.5 {
					fmt.Printf("%f\n", root)
				}

			}
		}
	}

	if numLabarotary == 3 && numTask == 2 {
		// Интервал, на котором нужно найти корни
		a := -30.0
		b := 40.0

		// Найти корни уравнения на заданном интервале
		roots := task32.FindRootsOnInterval(a, b)

		// Вывести найденные корни
		fmt.Println("Found unique roots of equations:")
		for _, root := range roots {
			fmt.Printf("Root: %f, Function value at root: %f\n", root, controller.Function(root))
		}
	}

	if numLabarotary == 3 && numTask == 3 {
		// Начальные предположения
		x := 1.0
		y := 1.0
		tolerance := 0.00001

		// Вызываем метод Ньютона-Рафсона для решения системы уравнений
		solutionX, solutionY := task33.NewtonRaphson(x, y, tolerance)

		// Выводим результат
		fmt.Printf("Solving a system of equations:\nx = %v\ny = %v\n", solutionX, solutionY)
	}

	if numLabarotary == 3 && numTask == 4 {
		// Инициализация границ интервала, на котором ищем минимум, и критерия останова
		leftBoundary := 10.1
		rightBoundary := 12.0
		epsilon := 0.0001

		// Вычисление минимума функции
		minimum := task34.InverseQuadraticInterpolation(leftBoundary, rightBoundary, epsilon)
		fmt.Printf("Minimum of the function on the interval [%v, %v]: %v\n", leftBoundary, rightBoundary, minimum)
	}

	if numLabarotary == 4 && numTask == 1 {
		// Пример набора точек
		points := [][]float64{{1.0, 5.1}, {2.0, 4.4}, {3.0, 3.2}, {4.0, 2.7}, {5, 2.55}}

		// Вывод полинома Лагранжа
		polynomial := task41.LagrangePolynomial(points)
		fmt.Println("Lagrange polynomial:")
		fmt.Println("L(x) =", polynomial)
		fmt.Println("---------------------------------------------------------------------------------------------------------------------------------------------------------")

		// Проверка интерполяции для некоторых значений x
		xs := []float64{2}
		fmt.Println("Answer is:")
		for _, x := range xs {
			fmt.Printf("L(%.1f) = %.4f\n", x, task41.LagrangeInterpolation(x, points))
		}
	}

	if numLabarotary == 4 && numTask == 2 {
		// Задаем функцию
		functions := []struct {
			name string
			f    func(float64) float64
		}{
			{"e^x", math.Exp},
			{"e^-x", func(x float64) float64 { return math.Exp(-x) }},
			{"sh(x)", func(x float64) float64 { return math.Sinh(x) }},
			{"ch(x)", func(x float64) float64 { return math.Cosh(x) }},
			{"sin(x)", math.Sin},
			{"cos(x)", math.Cos},
			{"ln(x)", math.Log},
		}

		// Задаем интервал и шаг
		start := 1.0
		end := 1.2
		h := 0.04 // Оставляем шаг 0.04

		// Создаем сетку значений x и y для каждой функции с более плотной сеткой
		for _, fn := range functions {
			fmt.Println("Function:", fn.name)
			var x, y []float64
			// Увеличиваем количество точек в 5 раз
			for xi := start; xi <= end; xi += h / 40 {
				x = append(x, xi)
				y = append(y, fn.f(xi))
			}

			// Строим кубический сплайн
			spline := task42.NewCubicSpline(x, y)

			// Находим значения сплайна в указанных точках
			points := []float64{1.18}
			for _, point := range points {
				fmt.Printf("  Spline at x = %.2f: %.6f\n", point, spline.Eval(point))
			}
			fmt.Println("----------------------------------")
		}
	}

	if numLabarotary == 4 && numTask == 3 {
		// Заданный отрезок и шаг
		h := 0.04

		// Вычислить значение сплайна в указанных точках
		points := []float64{1.10}

		// Интерполируемые функции: e^x, e^-x, sh(x), ch(x), sin(x), cos(x), ln(x)
		functions := map[string]func(float64) float64{
			"e^x":    controller4.FExp,
			"e^-x":   controller4.FExpNeg,
			"sh(x)":  controller4.FSinh,
			"ch(x)":  controller4.FCosh,
			"sin(x)": controller4.FSin,
			"cos(x)": controller4.FCos,
			"ln(x)":  controller4.FLn,
		}

		for name, function := range functions {
			// Для каждой функции в наборе функций
			fmt.Println("Function is:", name)
			for _, point := range points {
				// Для каждой точки, в которой нужно вычислить сплайн
				x0 := math.Floor(point/h) * h                                   // Находим ближайшую точку слева от заданной
				x1 := x0 + h                                                    // Находим ближайшую точку справа от заданной
				f0 := function(x0)                                              // Значение функции в точке слева
				f1 := function(x1)                                              // Значение функции в точке справа
				df0 := (function(x0+h) - function(x0-h)) / (2 * h)              // Центральная разностная схема для производной в точке слева
				df1 := (function(x1+h) - function(x1-h)) / (2 * h)              // Центральная разностная схема для производной в точке справа
				result := task43.HermiteSpline(point, x0, x1, f0, f1, df0, df1) // Вычисление сплайна Эрмита для заданной точки
				fmt.Printf("  The spline at the point %v: %v\n", point, result)
			}
			fmt.Println("--------------------------------------")
		}
	}

	if numLabarotary == 4 && numTask == 4 {
		// Опорные точки кривой Безье
		P := []controller4.Point{
			{0, 0},
			{1, 3},
			{2, -1},
			{3, 2},
		}

		// Вычисление точек на кривой Безье с шагом 0.1
		for t := 0.0; t <= 1.0; t += 0.1 {
			point := task44.BezierCurve(P, t)
			fmt.Printf("t = %.1f, Point: (%.2f, %.2f)\n", t, point.X, point.Y)
		}
	}

	if numLabarotary == 4 && numTask == 5 {
		x := 2.0    // Точка, в которой нужно вычислить производную
		h := 0.0001 // Шаг для численного дифференцирования

		derivative := task45.Differentiate(controller4.FunctionDif, x, h)
		fmt.Printf("The value of the derivative at the point x = %v: %v\n", x, derivative)
	}
}
