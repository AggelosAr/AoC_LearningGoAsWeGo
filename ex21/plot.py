import matplotlib.pyplot as plt
import re
import numpy as np


file_path = "output.txt"

steps = []
cells = []

with open(file_path, 'r') as file:
    for line in file:

        parts = line.split('->')

        step = parts[0]
        cell = parts[1]

        step = re.findall(r'\d+',step)[0]
        cell = re.findall(r'\d+',cell)[0]

        steps.append(int(step))
        cells.append(int(cell))



#Extrapolate pattern for steps 65, 196, 327.
#https://www.reddit.com/r/adventofcode/comments/18nevo3/2023_day_21_solutions/
steps = [steps[65], steps[196], steps[327]]
cells = [cells[65], cells[196], cells[327]]


steps = np.array(steps)
cells = np.array(cells)


degree = 2
coefficients = np.polyfit(steps, cells, degree)
poly_function = np.poly1d(coefficients)

"""
wSteps = [6, 10, 50, 100, 500, 1000, 5000]
correct = [16, 50, 1594, 6536, 167004, 668697, 16733044]
for idx, step in enumerate(wSteps):
    val = int(poly_function(step))
    print(f"For step {step}, value is: {val} - Should be: {correct[idx]}")
"""

print(f"For step {26501365}, value is: {int(poly_function(26501365))}")
#632880972219596 if put all
#632421652138917


plt.scatter(steps, cells, label='Original Data', color='blue')

x_values = np.linspace(min(steps), 5000, 5000)
y_values = poly_function(x_values)
plt.plot(x_values, y_values, label=f'Polynomial Fit (degree={degree})', linestyle='-', color='red')

plt.xlabel('Step')
plt.ylabel('Cells')
plt.title('Visualization of Steps and Cells with Polynomial Fit')
plt.legend()
plt.show()


