package task2

import (
	"Algorithms/laboratory_work1/controller"
	"fmt"
	"image/png"
	"math"
	"math/rand"
	"os"
)

// MonteKarlo - функция, которая реализует метод Монте-Карло для вычисления интеграла функции inFunc на интервале [a, b] с использованием n точек и iter итераций
func MonteKarlo(insideFunc func(x float64) float64, a, b, n float64, inaccuracy int) (float64, float64) {
	// Объявление переменных для хранения суммы значений функции, случайного числа и массива результатов
	var res float64
	var randFloat64 float64 = rand.Float64()*(b-a) + a
	arr := make([]float64, 0)

	// Итерация по всем итерациям метода Монте-Карло
	for j := 0; j < inaccuracy; j++ {
		// Итерация по всем точкам на интервале [a, b]
		for i := 1.0; i < n; i++ {
			// Добавление значения функции в точке randFloat64 к сумме значений функции
			res += insideFunc(randFloat64)
			// Генерация нового случайного числа randFloat64 на интервале [a, b]
			randFloat64 = rand.Float64()*math.Abs(b-a) + a
		}

		// Добавление среднего значения функции на интервале [a, b] к массиву результатов
		arr = append(arr, (res*(b-a))/n)
		// Обнуление суммы значений функции
		res = 0
	}

	// Возвращение первого элемента массива результатов и среднеквадратичного отклонения
	return arr[0], controller.AverageValueFloat64(arr)
}

// MonteKarloImage2 - функция, которая реализует метод Монте-Карло для вычисления площади Симферопольского водохранилища по изображению simferopol_reservoir.png
func MonteKarloImage2() (string, int, string, int) {
	// чтение изображения
	OpenImage, _ := os.Open("laboratory_work1/preview/simferopol_reservoir.png")
	// открытие изображения
	DecodeImage, _ := png.Decode(OpenImage)
	//массив куда будем записывать все итоги вычисления
	arr := make([]int, 0)

	// кол-во бросков
	var n int
	fmt.Print("Please, enter the number of throws to calculate the approximate area of the Simferopol reservoir: ")
	fmt.Scan(&n)
	fmt.Println()

	// анонимная функция возвращающая площадь в метрах квадратных
	var summator = func() int {
		var counter int // кол-во попаданий
		// rgba код голубого цвета
		var a, b, c, d uint32 = 144, 218, 238, 255

		// Итерация по всем броскам
		for i := 0; i < n; i++ {
			// генерация рандомных x,y
			x := rand.Intn(270)
			y := rand.Intn(357)
			// загрузка значений rgba пикселя с координатами (x,y)
			val1, val2, val3, val4 := DecodeImage.At(x, y).RGBA()

			/*
				В изображении, которое используется для вычисления площади водохранилища, голубой цвет воды представлен
				в формате RGBA (красный, зеленый, синий, альфа). Каждый из этих компонентов представлен в виде числа от 0 до
				255, где 0 - это минимальное значение, а 255 - максимальное.
				Цвет голубой воды Симферопольского водохранилища представлен в формате RGBA (144, 218, 238, 255).
				Таким образом, строка кода if (a == val1%256) && (b == val2%256) && (c == val3%256) && (d == val4%256)
				проверяет, совпадают ли значения красного, зеленого, синего и альфа каналов пикселя с соответствующими
				значениями голубого цвета.
			*/
			if (a == val1%256) && (b == val2%256) && (c == val3%256) && (d == val4%256) {
				counter += 1
			}
		}

		//вычисление площади в пикселях
		counter = (counter * 270 * 357) / n

		// в гугле выяснил площадь - 4,8 км, вырезал прямоугольник с ними
		// ширина изображения 270 px. Длина одной стороны пикселя 4800/270=17,7777...
		// площадь одного пикселя в метрах 17.7777...*17.7777...=317 м кв
		return counter * 475
	}

	// Запускаем алгоритм несколько раз
	for i := 0; i < 10; i++ {
		arr = append(arr, summator())
	}

	// возвращаем площадь и погрешность
	return "Square is", arr[0], "| Inaccuracy is", controller.AverageValueInt(arr)
}
