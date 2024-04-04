package controller

// Функция abs возвращает абсолютное значение числа
func Abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
