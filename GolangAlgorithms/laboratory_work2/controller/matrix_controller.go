package controller

import "Algorithms/laboratory_work2/helpers"

// MultiplyMatrixVector - умножение матрицы на вектор
func MultiplyMatrixVector(A helpers.Matrix, X helpers.Vector) helpers.Vector {
	n := len(A)                       // Получаем размер матрицы A
	result := make(helpers.Vector, n) // Создаем вектор result длиной n

	for i := 0; i < n; i++ { // Начинаем цикл по элементам вектора result
		sum := 0.0               // Инициализируем переменную sum
		for j := 0; j < n; j++ { // Начинаем цикл по элементам вектора X
			sum += A[i][j] * X[j] // Считаем сумму
		}
		result[i] = sum // Записываем результат в вектор result
	}

	return result // Возвращаем результат
}
