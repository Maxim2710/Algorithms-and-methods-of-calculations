import random

# Функция для вычисления значения целевой функции F(X)
def objective_function(x):
    return 3*x[0] + 2*x[1] - 4*x[2]

# Функция проверки выполнения ограничений
def constraints_satisfied(x):
    return x[0] + x[1] - 2*x[2] >= 4 and 3*x[0] + x[1] - 4*x[2] <= 2

# Инициализация параметров муравьиного алгоритма
num_ants = 10
num_iterations = 100
alpha = 1.0  # Вес феромона
beta = 2.0   # Вес эвристики
rho = 0.5    # Коэффициент испарения феромона
pheromone = 1.0  # Инициализация феромона

# Инициализация начальных значений переменных
num_variables = 3
variables = [[random.uniform(0, 10) for _ in range(num_variables)] for _ in range(num_ants)]
best_solution = None
best_score = float('inf')

# Запуск муравьиного алгоритма
for _ in range(num_iterations):
    for ant in range(num_ants):
        # Выбор следующей переменной с помощью правила выбора муравья
        next_variable = [0] * num_variables
        for i in range(num_variables):
            if random.uniform(0, 1) < 0.5:
                next_variable[i] = random.uniform(0, 10)
            else:
                next_variable[i] = variables[ant][i] + random.uniform(-1, 1)
        # Проверка выполнения ограничений
        if constraints_satisfied(next_variable):
            score = objective_function(next_variable)
            if best_solution is None or score < best_score:
                best_solution = next_variable
                best_score = score
    # Обновление феромона
    if best_solution is not None:
        for i in range(num_variables):
            for j in range(num_variables):
                pheromone *= rho
                if i == j:
                    continue
                pheromone += alpha / (1 + abs(best_solution[i] - best_solution[j]))
            variables[ant][i] = best_solution[i]

# Вывод результата
print("Лучшее решение:", best_solution)
print("Значение целевой функции F(X):", best_score)

