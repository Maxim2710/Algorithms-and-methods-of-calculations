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
		A := helpers.Matrix{{2, -1, 0}, {0, 5, 2}, {1, -1, 3}}
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
		A := helpers.Matrix{{2, -1, 0}, {0, 5, 2}, {1, -1, 3}}
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
			{0, 5, 2},
			{1, -1, 3},
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
}
