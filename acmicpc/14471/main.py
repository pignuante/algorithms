from sys import stdin

input = stdin.readline

N, M = list(map(int, input().split()))
wins = []
loses = []
for _ in range(M):
    a, b = list(map(int, input().split()))
    wins.append(a)
    loses.append(b)

wins = sorted(wins, reverse=True)
result = 0

for i in range(M-1):
    if (N-wins[i] > 0):
        result += (N - wins[i])
        pass
    pass
print(result)
