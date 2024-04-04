from tabulate import tabulate


def knapsack(values, weights, capacity):
    n = len(values)
    K = [[0 for x in range(capacity + 1)] for x in range(n + 1)]

    for i in range(n + 1):
        for w in range(capacity + 1):
            if i == 0 or w == 0:
                K[i][w] = 0
            elif weights[i - 1] <= w:
                K[i][w] = max(values[i - 1] + K[i - 1][w - weights[i - 1]], K[i - 1][w])
            else:
                K[i][w] = K[i - 1][w]

    result = K[n][capacity]
    print(f"Максимальная стоимость: {result}")

    w = capacity
    taken = [0] * n
    for i in range(n, 0, -1):
        if result <= 0:
            break
        if result == K[i - 1][w]:
            continue
        else:
            taken[i - 1] = 1
            result -= values[i - 1]
            w -= weights[i - 1]

    items_taken = []
    for i in range(n):
        if taken[i]:
            items_taken.append([f"Предмет {i + 1}", weights[i], values[i]])

    print(tabulate(items_taken, headers=["Название", "Вес", "Стоимость"], tablefmt="pretty"))


values = [60, 100, 120, 90, 75, 50, 30]
weights = [10, 20, 30, 15, 10, 5, 5]
capacity = 50

knapsack(values, weights, capacity)
