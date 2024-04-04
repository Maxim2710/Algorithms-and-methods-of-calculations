import itertools
import networkx as nx
import matplotlib.pyplot as plt

def calculate_distance(path, distances):
    total_distance = 0
    for i in range(len(path) - 1):
        total_distance += distances[path[i]][path[i+1]]
    total_distance += distances[path[-1]][path[0]]
    return total_distance

def traveling_salesman_brute_force(distances):
    num_cities = len(distances)
    min_distance = float('inf')
    min_path = None

    for path in itertools.permutations(range(num_cities)):
        distance = calculate_distance(path, distances)
        if distance < min_distance:
            min_distance = distance
            min_path = path

    return min_path, min_distance

# Создаем граф
G = nx.DiGraph()

# Добавляем узлы (города)
num_cities = 5
city_names = ["Москва", "Санкт-Петербург", "Новосибирск", "Екатеринбург", "Красноярск"]
for i in range(num_cities):
    G.add_node(i, label=city_names[i])

# Добавляем ребра (пути между городами)
distances = [
    [0, 10, 15, 20, 25],
    [10, 0, 35, 25, 30],
    [15, 35, 0, 30, 20],
    [20, 25, 30, 0, 10],
    [25, 30, 20, 10, 0]
]
for i in range(num_cities):
    for j in range(num_cities):
        if i != j:
            G.add_edge(i, j, weight=distances[i][j])

# Определяем оптимальный маршрут
shortest_path, shortest_distance = traveling_salesman_brute_force(distances)

# Отрисовываем граф с оптимальным маршрутом
pos = nx.circular_layout(G)
edge_labels = {(i, j): distances[i][j] for i, j in G.edges()}
nx.draw(G, pos, with_labels=True, node_size=1000, node_color='lightblue', font_size=10, font_weight='bold')
nx.draw_networkx_edge_labels(G, pos, edge_labels=edge_labels, font_color='red')
plt.title('Traveling Salesman Problem')

# Рисуем оптимальный маршрут
for i in range(len(shortest_path) - 1):
    plt.arrow(pos[shortest_path[i]][0], pos[shortest_path[i]][1],
              pos[shortest_path[i+1]][0] - pos[shortest_path[i]][0],
              pos[shortest_path[i+1]][1] - pos[shortest_path[i]][1],
              shape='full', lw=1, length_includes_head=True, head_width=0.1, color='blue')
plt.arrow(pos[shortest_path[-1]][0], pos[shortest_path[-1]][1],
          pos[shortest_path[0]][0] - pos[shortest_path[-1]][0],
          pos[shortest_path[0]][1] - pos[shortest_path[-1]][1],
          shape='full', lw=1, length_includes_head=True, head_width=0.1, color='blue')

plt.show()

# Выводим лучший путь и расстояние
print("Лучший путь:")
for city_index in shortest_path:
    print(city_names[city_index], end=" -> ")
print(city_names[shortest_path[0]])  # Чтобы замкнуть путь

print("Расстояние:", shortest_distance)
