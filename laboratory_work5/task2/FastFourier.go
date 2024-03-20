package task2

import (
	"math"
	"math/cmplx"
)

// FastFourier - функция для быстрого преобразования Фурье
func FastFourier(x []complex128) []complex128 {
	// Получаем длину входного массива
	N := len(x)
	// Преобразуем длину в комплексное число
	NComplex := complex(float64(N), 0)

	// Если длина массива меньше или равна 1, возвращаем его как есть
	if N <= 1 {
		return x
	}

	// Создаем массивы для хранения четных и нечетных элементов
	even := make([]complex128, 0)
	odd := make([]complex128, 0)

	// Разделяем входной массив на четные и нечетные элементы
	for i := 0; i < len(x); i++ {
		if i%2 == 0 {
			even = append(even, x[i])
		} else {
			odd = append(odd, x[i])
		}
	}

	result := make([]complex128, 0)

	// Рекурсивно вызываем FastFourier для четных и нечетных элементов
	evenFFT := FastFourier(even)
	oddFFT := FastFourier(odd)

	// Создаем массив для хранения множителей
	factor := make([]complex128, 0)

	// Проходим по половине массива и вычисляем множители
	for k := 0; k < N/2; k++ {
		// Преобразуем текущий индекс в комплексное число
		kComplex := complex(float64(k), 0)
		// Вычисляем экспоненту
		exponent := cmplx.Exp(complex(0, -2) * complex(math.Pi, 0) * kComplex / NComplex)
		// Добавляем множитель в массив
		factor = append(factor, exponent*oddFFT[k])
		// Добавляем результат преобразования Фурье для текущего элемента
		result = append(result, evenFFT[k]+factor[k])
	}

	// Добавляем оставшуюся часть результата преобразования Фурье
	for k := 0; k < N/2; k++ {
		result = append(result, evenFFT[k]-factor[k])
	}

	// Возвращаем результат
	return result
}
