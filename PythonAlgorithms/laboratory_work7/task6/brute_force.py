def calculate_F(x1, x2, x3):
    return 3 * x1 + 2 * x2 - 4 * x3

def brute_force():
    best_solution = None
    best_value = float('-inf')
    for x1 in range(10):  # Перебираем x1 от 0 до 9
        for x2 in range(10):  # Перебираем x2 от 0 до 9
            for x3 in range(10):  # Перебираем x3 от 0 до 9
                if x1 + x2 - 2 * x3 >= 4 and 3 * x1 + x2 - 4 * x3 <= 2:
                    current_value = calculate_F(x1, x2, x3)
                    if current_value > best_value:
                        best_solution = (x1, x2, x3)
                        best_value = current_value
                        return best_solution  # Прерываем цикл после нахождения первого решения
    return best_solution

solution = brute_force()
print("Optimal solution:")
print("x1 =", solution[0])
print("x2 =", solution[1])
print("x3 =", solution[2])
print("F(X) =", calculate_F(*solution))
