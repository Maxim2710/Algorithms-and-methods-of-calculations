package task4

import "Algorithms/laboratory_work4/controller"

// Вычисление точки на кривой Безье
func BezierCurve(P []controller.Point, t float64) controller.Point {
	n := len(P) - 1            // Определение степени кривой Безье (на один меньше количества контрольных точек)
	var point controller.Point // Инициализация переменной для хранения результирующей точки

	for i := 0; i <= n; i++ {
		// Вычисление базисной функции Бернштейна для i-го контрольного полинома на момент времени t
		bern := controller.Bernstein(n, i, t)
		// Приращение к точке: добавление i-го контрольного полинома, взвешенного соответствующей базисной функцией, к результирующей точке
		point.X += P[i].X * bern
		point.Y += P[i].Y * bern
	}

	return point // Возврат результирующей точки на кривой Безье
}
