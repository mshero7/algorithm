def solution(arr, divisor):
    return sorted([i for i in arr if i % divisor == 0]) or [-1]

print(solution([5,9,7,10], 5))
print(solution([3,2,6], 10))