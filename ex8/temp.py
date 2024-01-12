import time


def gcd(a, b):
    while b != 0:
        a, b = b, a % b
    return a

def lcm(numbers):
    if not numbers:
        return 0

    result = numbers[0]
    for num in numbers[1:]:
        result = (result * num) // gcd(result, num)

    return result


start_time = time.perf_counter()
arr = [14681, 21883, 13019, 16897, 16343, 20221]
result1 = lcm(arr)
end_time = time.perf_counter()
elapsed_time = end_time - start_time

print(f"LCM: {result1}, Time: {elapsed_time * 1000:.4f} ms")

