import matplotlib.pyplot as plt

def schedule(tasks):
    # Сортируем задачи по времени завершения
    sorted_tasks = sorted(tasks, key=lambda x: x[1])

    # Инициализируем пустой список для сохранения выбранных задач
    selected_tasks = []

    # Инициализируем переменную для отслеживания времени окончания последней выбранной задачи
    last_end_time = float('-inf')

    # Проходимся по всем задачам
    for task in sorted_tasks:
        start_time, end_time = task
        # Если время начала текущей задачи больше времени окончания последней задачи,
        # то добавляем текущую задачу в список выбранных задач и обновляем last_end_time
        if start_time >= last_end_time:
            selected_tasks.append(task)
            last_end_time = end_time

    return selected_tasks

def visualize(tasks, selected_tasks):
    # Создаем график
    fig, ax = plt.subplots()

    # Рисуем задачи
    for idx, task in enumerate(tasks):
        start_time, end_time = task
        ax.barh(idx, end_time - start_time, left=start_time, color='lightblue', edgecolor='black')

    # Рисуем выбранные задачи
    for idx, task in enumerate(selected_tasks):
        start_time, end_time = task
        ax.barh(idx, end_time - start_time, left=start_time, color='blue', edgecolor='black')

    ax.set_yticks(range(len(tasks)))
    ax.set_yticklabels([f'Task {i+1}' for i in range(len(tasks))])
    ax.set_xlabel('Time')
    ax.set_ylabel('Tasks')
    ax.set_title('Scheduled Tasks')

    plt.show()

# Пример использования
tasks = [(1, 4), (3, 5), (0, 6), (5, 7), (3, 8), (5, 9), (6, 10), (8, 11), (8, 12), (2, 13), (12, 14), (15,20)]
selected_tasks = schedule(tasks)
print("Максимальное количество непересекающихся задач:", len(selected_tasks))
print("Выбранные задачи:", selected_tasks)

# Визуализируем результат
visualize(tasks, selected_tasks)
