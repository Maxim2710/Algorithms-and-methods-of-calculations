import numpy as np
import matplotlib.pyplot as plt

def himmelblau(x):
    return (x[0]**2 + x[1] - 11)**2 + (x[0] + x[1]**2 - 7)**2

def himmelblau_gradient(x):
    return np.array([
        4 * x[0] * (x[0]**2 + x[1] - 11) + 2 * (x[0] + x[1]**2 - 7),
        2 * (x[0]**2 + x[1] - 11) + 4 * x[1] * (x[0] + x[1]**2 - 7)
    ])

def himmelblau_hessian(x):
    return np.array([
        [12 * x[0]**2 + 4 * x[1] - 42, 4 * (x[0] + x[1])],
        [4 * (x[0] + x[1]), 4 * x[0] + 12 * x[1]**2 - 26]
    ])

def newton_method(f, f_prime, f_double_prime, x0, tol=1e-5, max_iter=100):
    x = x0
    for _ in range(max_iter):
        grad = f_prime(x)
        hess_inv = np.linalg.inv(f_double_prime(x))
        delta_x = -np.dot(hess_inv, grad)
        x = x + delta_x
        if np.linalg.norm(delta_x) < tol:
            break
    return x

def plot_himmelblau():
    x = np.linspace(-6, 6, 100)
    y = np.linspace(-6, 6, 100)
    X, Y = np.meshgrid(x, y)
    Z = himmelblau([X, Y])

    plt.figure(figsize=(8, 6))
    plt.contour(X, Y, Z, levels=np.logspace(0, 5, 35), cmap='viridis', alpha=0.8)
    plt.colorbar(label='log scale')
    plt.xlabel('x')
    plt.ylabel('y')

def main():
    plot_himmelblau()
    starting_points = [
        [-4, -4],
        [-4, 4],
        [4, -4],
        [4, 4]
    ]
    for point in starting_points:
        minimum = newton_method(himmelblau, himmelblau_gradient, himmelblau_hessian, point)
        print(f"Минимум в точке: {minimum}")
        plt.scatter(minimum[0], minimum[1], color='red', marker='o')
    plt.show()

if __name__ == "__main__":
    main()
