# import math
# import matplotlib.pyplot as plt
# from integral import integral
# from generator import generator_values
#
# FOURIER_EPS = math.pow(10, -4)
# N = 2
# pi = math.pi
#
#
# def func(x):
#     return math.sin(x) + math.cos(2 * x) + x * x / 10.0
#
#
# def fourier(func, N, x, T):
#     a0 = 2 / T * integral(func, 0, T)
#     f = a0 / 2
#     for k in range(0, N + 1):
#         a = 2 / T * integral(lambda first: func(first) * math.cos(k * first), 0, T)
#         b = 2 / T * integral(lambda second: func(second) * math.sin(k * second), 0, T)
#         f += a * math.cos(k * x) + b * math.sin(k * x)
#
#     return f
#
#
# if __name__ == '__main__':
#     x_coords = generator_values(-2, 2, 0.025)
#     y_coords = [fourier(func, N, i, 2 * pi) for i in x_coords]
#
#     plt.grid()
#     plt.plot(x_coords, y_coords, marker='o', markersize=2)
#     plt.show()

import math
import matplotlib.pyplot as plt

from generator import generator_values
from integral import integral

FOURIER_EPS = math.pow(10, -4)
N = 50
pi = math.pi

def func(x):
    return math.sin(x) + math.cos(2 * x) + x * x / 10.0


def fourier(func, N, x, T):
    a0 = 2 / T * integral(func, 0, T)
    f = a0 / 2
    for k in range(1, N + 1):
        a = 2 / T * integral(lambda first: func(first) * math.cos(k * first), 0, T)
        b = 2 / T * integral(lambda second: func(second) * math.sin(k * second), 0, T)
        f += a * math.cos(k * x) + b * math.sin(k * x)

    return f


if __name__ == '__main__':
    x_coords = generator_values(0, 7, 0.025)
    y_coords = [fourier(func, N, i, 2 * pi) for i in x_coords]

    plt.grid()
    plt.plot(x_coords, y_coords, marker='o', markersize=2)
    plt.show()
