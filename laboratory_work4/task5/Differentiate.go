package task5

// Численное дифференцирование
func Differentiate(f func(float64) float64, x float64, h float64) float64 {
	// Функция Differentiate принимает на вход функцию f, точку x и шаг h и возвращает приближенное значение производной функции f в точке x с использованием односторонней разностной схемы.
	return (f(x+h) - f(x)) / h // Вычисление разности значений функции в точке x+h и x, а затем деление на шаг h.
}