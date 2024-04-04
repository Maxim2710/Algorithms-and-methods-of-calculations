import numpy as np
import matplotlib.pyplot as plt

# Функция для поиска всех локальных минимумов заданной функции с помощью наискорейшего спуска
def find_local_minima(func, grad_func, initial_points, learning_rate=0.01, tolerance=1e-6, max_iterations=10000):
    min_points = set()  # используем set для хранения уникальных точек минимума
    for initial_point in initial_points:
        current_point = initial_point
        iterations = 0

        while iterations < max_iterations:
            gradient = grad_func(*current_point)
            if np.linalg.norm(gradient) < tolerance:
                min_points.add(tuple(np.round(current_point, decimals=6)))  # добавляем точку минимума в set
                break

            # Определяем длину шага по направлению антиградиента
            step_size = dichotomy_search(func, grad_func, current_point, gradient)

            # Двигаемся вдоль антиградиента с учетом выбранной длины шага
            current_point = current_point - learning_rate * step_size * gradient
            iterations += 1

    return list(min_points)  # преобразуем set обратно в список для удобства использования

# Метод дихотомии для поиска оптимальной длины шага
def dichotomy_search(func, grad_func, current_point, gradient, tol=1e-6):
    # Определение функции, которую нужно минимизировать
    def line_search_function(alpha):
        return func(*(current_point - alpha * gradient))

    # Начальные значения
    a, b = 0, 1
    delta = tol / 10

    # Итеративный процесс поиска оптимальной длины шага
    while (b - a) > tol:
        # Выбор двух промежуточных точек
        x1 = (a + b) / 2 - delta
        x2 = (a + b) / 2 + delta

        # Оценка значений функции в выбранных точках
        f1 = line_search_function(x1)
        f2 = line_search_function(x2)

        # Обновление интервала
        if f1 < f2:
            b = x2
        else:
            a = x1

    return (a + b) / 2

# Определим функцию Химмельблау для демонстрации
def himmelblau(x, y):
    return (x ** 2 + y - 11) ** 2 + (x + y ** 2 - 7) ** 2

# Определим градиент функции Химмельблау
def himmelblau_gradient(x, y):
    dx = 4 * x * (x ** 2 + y - 11) + 2 * (x + y ** 2 - 7)
    dy = 2 * (x ** 2 + y - 11) + 4 * y * (x + y ** 2 - 7)
    return np.array([dx, dy])

# Начальные точки для запуска наискорейшего спуска
initial_points = [(x, y) for x in np.linspace(-6, 6, 10) for y in np.linspace(-6, 6, 10)]

# Найдем все локальные минимумы функции Химмельблау с помощью наискорейшего спуска
min_points = find_local_minima(himmelblau, himmelblau_gradient, initial_points)

print("Найденные локальные минимумы:", min_points)

# Рисуем график
x = np.linspace(-6, 6, 400)
y = np.linspace(-6, 6, 400)
X, Y = np.meshgrid(x, y)
Z = himmelblau(X, Y)

plt.figure(figsize=(8, 6))
plt.contour(X, Y, Z, levels=np.logspace(0, 5, 35))
plt.scatter(*zip(*min_points), color='red', label='Локальные минимумы')
plt.title('Локальные минимумы функции Химмельблау с помощью наискорейшего спуска (метод дихотомии)')
plt.xlabel('x')
plt.ylabel('y')
plt.legend()
plt.grid(True)
plt.show()




