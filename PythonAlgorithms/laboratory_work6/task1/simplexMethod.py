from scipy.optimize import linprog

# Коэффициенты целевой функции (13 * x1 + 8 * x2 + 4 * x3), коэффициенты домножены на -1 для решения задачи о максимизации
objective_function = [-13, -8, -4]

# Матрица коэффициентов левой части ограничений (A * x <= b)
constraints_coefficients = [[2, 5, 0],
                            [3, 1, 1],
                            [13, 4, 4]]

# Правая часть ограничений
constraints_rhs = [80, 50, 210]

# Границы переменных (x >= 0)
x1_bounds = (0, None)  # Bound for x1
x2_bounds = (0, None)  # Bound for x2
x3_bounds = (0, None)  # Bound for x3

# Решение задачи линейного программирования
result = linprog(objective_function, A_ub=constraints_coefficients, b_ub=constraints_rhs,
                 bounds=[x1_bounds, x2_bounds, x3_bounds], method='highs')

# Вывод результатов
print("Optimal plan:")
# Вывод значений переменных оптимального плана
for i, value in enumerate(result.x):
    print(f"x{i+1} =", value)
# Вывод значения целевой функции для оптимального плана (противоположный знак, так как решаем задачу о максимизации)
print("Value of the objective function F(X) =", -result.fun)
