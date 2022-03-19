from sys import stdin
input = stdin.readline

N = int(input())
stores = list(map(int, input().split()))
count = 0
for i in range(N):
    if stores[i] == count % 3:
        count += 1
print(count)
