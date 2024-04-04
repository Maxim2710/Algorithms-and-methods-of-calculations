import random
from tabulate import tabulate

# Преобразование решения в удобный формат таблицы
def solution_to_table(solution):
    items = []
    for i in range(len(solution)):
        if solution[i] == 1:
            items.append((f'Предмет {i+1}', values[i], weights[i]))
    return items

# Отображение решения в виде таблицы
def display_solution_table(solution):
    table = solution_to_table(solution)
    print(tabulate(table, headers=['Предмет', 'Стоимость', 'Вес'], tablefmt='grid', stralign='center'))

# Входные данные
values = [60, 100, 120, 90, 75, 50, 30]
weights = [10, 20, 30, 15, 10, 5, 5]
capacity = 50

# Генерация начальной популяции
def generate_population(size):
    population = []
    for _ in range(size):
        chromosome = [random.randint(0, 1) for _ in range(len(values))]
        population.append(chromosome)
    return population

# Оценка приспособленности хромосомы
def fitness(chromosome):
    total_weight = sum(weights[i] for i in range(len(chromosome)) if chromosome[i] == 1)
    total_value = sum(values[i] for i in range(len(chromosome)) if chromosome[i] == 1)
    if total_weight > capacity:
        return 0
    else:
        return total_value

# Выбор лучших особей
def selection(population, k=10):
    sorted_population = sorted(population, key=lambda x: fitness(x), reverse=True)
    return sorted_population[:k]

# Скрещивание особей
def crossover(parent1, parent2):
    crossover_point = random.randint(1, len(parent1) - 1)
    child1 = parent1[:crossover_point] + parent2[crossover_point:]
    child2 = parent2[:crossover_point] + parent1[crossover_point:]
    return child1, child2

# Мутация
def mutation(chromosome, probability=0.1):
    mutated_chromosome = chromosome[:]
    for i in range(len(mutated_chromosome)):
        if random.random() < probability:
            mutated_chromosome[i] = 1 - mutated_chromosome[i]
    return mutated_chromosome

# Генетический алгоритм
def genetic_algorithm(population_size=100, generations=1000):
    population = generate_population(population_size)
    for _ in range(generations):
        new_population = []
        for _ in range(population_size // 2):
            parent1, parent2 = random.choices(population, k=2)
            child1, child2 = crossover(parent1, parent2)
            child1 = mutation(child1)
            child2 = mutation(child2)
            new_population.extend([child1, child2])
        population = selection(population + new_population, k=population_size)
    return population

# Ручная проверка решения
def manual_check(solution):
    total_weight = sum(weights[i] for i in range(len(solution)) if solution[i] == 1)
    total_value = sum(values[i] for i in range(len(solution)) if solution[i] == 1)
    print("Total Weight:", total_weight)
    print("Total Value:", total_value)

# Запуск генетического алгоритма
best_solution = genetic_algorithm()[0]
print("Best Solution:", best_solution)
display_solution_table(best_solution)
manual_check(best_solution)
