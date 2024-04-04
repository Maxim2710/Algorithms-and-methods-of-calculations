import numpy as np
import matplotlib.pyplot as plt


def fft(a):
    # Базовый случай: если длина массива равна 1, возвращаем его
    n = len(a)
    if n == 1:
        return a

    # Разбиваем массив на две части
    p = fft(a[::2])  # Чётные элементы
    q = fft(a[1::2])  # Нечётные элементы

    # Рассчитываем первообразный корень из единицы
    epsilon = np.exp(2j * np.pi / n)

    # Выполняем FFT
    b = np.zeros(n, dtype=np.complex128)
    for i in range(n // 2):
        b[i] = p[i] + q[i] * epsilon ** i
        b[i + n // 2] = p[i] + q[i] * epsilon ** (i + n // 2)

    return b


# Создание входного сигнала
t = np.linspace(0, 1, 300, endpoint=False)
a = np.sin(2 * np.pi * 10 * t) + 0.5 * np.sin(2 * np.pi * 20 * t)

if __name__ == "__main__":
    # Применение FFT
    result = fft(a)

    # Вычисление амплитуд и фаз выходных данных
    amplitudes = np.abs(result)
    phases = np.angle(result)

    # Построение графиков
    plt.figure(figsize=(12, 6))

    plt.subplot(2, 1, 1)
    plt.plot(amplitudes)
    plt.title('Амплитуды выходных данных после применения FFT')
    plt.xlabel('Частота')
    plt.ylabel('Амплитуда')

    plt.subplot(2, 1, 2)
    plt.plot(phases)
    plt.title('Фазы выходных данных после применения FFT')
    plt.xlabel('Частота')
    plt.ylabel('Фаза')

    plt.tight_layout()
    plt.show()
