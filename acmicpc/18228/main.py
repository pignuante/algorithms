from sys import stdin

input = stdin.readline

N =int(input())
A = list(map(int, input().split()))
idx = A.index(-1)
left, right = A[:idx], A[idx+1:]
print(min(left) + min(right))
