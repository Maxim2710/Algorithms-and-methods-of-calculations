package main

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
	task31 "Algorithms/laboratory_work3/task1"
	task32 "Algorithms/laboratory_work3/task2"
	task33 "Algorithms/laboratory_work3/task3"
	task34 "Algorithms/laboratory_work3/task4"
	"fmt"
	"log"
)

func main() {
	var numLabarotary, numTask int
	fmt.Println()
	fmt.Print("Please, enter the number of the laboratory work: ")
	fmt.Scan(&numLabarotary)

	fmt.Print("Laboratory work number: ")
	fmt.Print(numLabarotary)
	fmt.Print("; Select the task1 number: ")
	fmt.Scan(&numTask)
	fmt.Println("Laboratory work:", numLabarotary, "| Task:", numTask)
	fmt.Println()

	if numLabarotary == 1 && numTask == 1 {

		a := 0.0
		b := 3.0
		var n int
		epsilon := 1e-3

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
		maxIterations := 1000

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
		// Начальные границы отрезка и требуемая точность
		a := -40.0
		b := 40.0
		epsilon := 0.0001

		// Ищем корень методом дихотомии
		root := task31.DichotomyMethod(a, b, epsilon)

		// Выводим результат
		fmt.Printf("Approximate value of the root is: %f\n", root)
	}

	if numLabarotary == 3 && numTask == 2 {
		// Начальное предположение и требуемая точность
		guess := 0.5
		epsilon := 0.0001

		// Вызываем метод Ньютона для поиска корня
		root, n := task32.NewtonMethod(guess, epsilon)

		// Выводим результат
		fmt.Printf("The root of the equation is: %v\n", root)
		fmt.Printf("Number of iterations: %v\n", n)
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
		leftBoundary := -150.0
		rightBoundary := 1000.5
		epsilon := 0.0001

		// Вычисление минимума функции
		minimum := task34.InverseQuadraticInterpolation(leftBoundary, rightBoundary, epsilon)
		fmt.Printf("Minimum of the function on the interval [%v, %v]: %v\n", leftBoundary, rightBoundary, minimum)
	}
}
