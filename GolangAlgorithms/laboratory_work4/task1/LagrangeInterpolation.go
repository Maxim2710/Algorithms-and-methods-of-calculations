package task1

import (
	"fmt"
	"strings"
)

// LagrangePolynomial вычисляет полином Лагранжа для набора точек
func LagrangePolynomial(points [][]float64) string {
	var terms []string
	for i, p := range points {
		term := fmt.Sprintf("(%.2f)", p[1]) // Формируем начальное значение для одного члена полинома Лагранжа.
		for j, q := range points {
			if i != j {
				term += fmt.Sprintf(" * (x - %.2f) / (%.2f - %.2f)", q[0], p[0], q[0]) // Добавляем члены полинома Лагранжа.
			}
		}
		terms = append(terms, term) // Добавляем сформированный член полинома Лагранжа в общий список членов.
	}
	return strings.Join(terms, " + ") // Объединяем все члены полинома Лагранжа в одну строку с разделителем "+".
}

// LagrangeInterpolation вычисляет значение интерполяционного многочлена Лагранжа в точке x
func LagrangeInterpolation(x float64, points [][]float64) float64 {
	result := 0.0 // Инициализируем результат как 0.
	for i := range points {
		term := points[i][1] // Начинаем с базового значения, которое соответствует y-координате точки.
		for j := range points {
			if j != i {
				term *= (x - points[j][0]) / (points[i][0] - points[j][0]) // Вычисляем следующий член полинома Лагранжа.
			}
		}
		result += term // Добавляем текущий член к результату.
	}
	return result // Возвращаем итоговый результат интерполяции.
}
