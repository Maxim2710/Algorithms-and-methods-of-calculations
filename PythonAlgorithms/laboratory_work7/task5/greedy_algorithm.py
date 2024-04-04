from tabulate import tabulate

def objective_function(x1, x2, x3):
    return 3 * x1 + 2 * x2 - 4 * x3

def satisfy_constraints(x1, x2, x3):
    return x1 + x2 - 2 * x3 >= 4 and 3 * x1 + x2 - 4 * x3 <= 2

def greedy_algorithm():
    best_solution = {'x1': None, 'x2': None, 'x3': None}
    best_value = float('inf')

    # Перебираем возможные значения x1, x2 и x3
    for x1 in range(100):  # Просто выбираем большое число для диапазона, так как они ограничены условиями
        for x2 in range(100):
            for x3 in range(100):
                # Проверяем удовлетворение ограничениям
                if satisfy_constraints(x1, x2, x3):
                    # Вычисляем значение целевой функции
                    value = objective_function(x1, x2, x3)

                    # Обновляем лучшее решение, если нашли лучшее значение
                    if value < best_value:
                        best_solution['x1'] = x1
                        best_solution['x2'] = x2
                        best_solution['x3'] = x3
                        best_value = value

    # Получаем лучшее решение
    best_solution['F(X)'] = best_value
    return best_solution

# Запускаем алгоритм
solution = greedy_algorithm()

# Формируем таблицу
table = [["Variable", "Value"],
         ["x1", solution['x1']],
         ["x2", solution['x2']],
         ["x3", solution['x3']],
         ["F(X)", solution['F(X)']]]

# Выводим таблицу с использованием стиля "pretty"
print(tabulate(table, tablefmt="pretty"))
