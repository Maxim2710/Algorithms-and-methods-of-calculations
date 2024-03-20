package controller

import "math" // Импортируем пакет math

// MistakeCounter - функция для подсчета средней ошибки между двумя массивами arr1 и arr2
func MistakeCounter(arr1, arr2 []float64) float64 {
	// Проверяем, имеют ли массивы одинаковую длину
	if len(arr1) != len(arr2) {
		panic("Arrays have different lengths") // Если массивы имеют разные длины, вызываем панику
	}

	var sumError float64             // Переменная для хранения суммарной ошибки
	for i := 0; i < len(arr1); i++ { // Проходим по элементам массивов
		sumError += math.Abs(arr1[i] - arr2[i]) // Добавляем абсолютное значение разности элементов массивов к суммарной ошибке
	}

	return sumError / float64(len(arr1)) // Возвращаем среднюю ошибку, разделив суммарную ошибку на длину массива
}
