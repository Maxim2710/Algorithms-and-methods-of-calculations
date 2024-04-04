# import math
#
# INTEGRAL_EPS = math.pow(10, -4)
#
# def integral(f, a: float, b: float, eps=INTEGRAL_EPS) -> float:
#     def simpson_rule(f, a, b, n) -> float:
#         h = (b - a) / n
#         sum = f(a) + f(b)
#         for i in range(1, n):
#             k = 2 + 2 * (i % 2)
#             sum += k * f(a + i * h)
#         sum *= h / 3
#         return sum
#
#     num_intervals = 4
#     difference, integral_value, integral_value2 = 1, 1, 1
#
#     while abs(difference) >= eps:
#         integral_value = simpson_rule(f, a, b, num_intervals)
#         num_intervals *= 2
#         integral_value2 = simpson_rule(f, a, b, num_intervals)
#         difference = (integral_value - integral_value2) / 15
#
#     return integral_value2 + difference
import math

from scipy.integrate import quad

INTEGRAL_EPS = math.pow(10, -4)


def integral(f, a: float, b: float, eps=INTEGRAL_EPS) -> float:
    result, _ = quad(f, a, b)
    return result
