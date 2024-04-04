from pulp import *

# Создание экземпляра задачи оптимизации
prob = LpProblem("Минимизация целевой функции", LpMinimize)

# Определение переменных решения
x1 = LpVariable("x1", lowBound=0)
x2 = LpVariable("x2", lowBound=0)
x3 = LpVariable("x3", lowBound=0)

# Определение целевой функции
prob += 3*x1 + 2*x2 - 4*x3, "F(X)"

# Добавление ограничений
prob += x1 + x2 - 2*x3 >= 4
prob += 3*x1 + x2 - 4*x3 <= 2

# Решение задачи оптимизации
prob.solve()

# Вывод результатов
print("Оптимальный план:")
print(f"x1 = {value(x1)}, x2 = {value(x2)}, x3 = {value(x3)}")
print(f"F(X) = {value(prob.objective)}")
