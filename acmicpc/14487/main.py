from sys import stdin

input = stdin.readline

n = int(input())
v = sorted(list(map(int, input().split())))
print(sum(v) - max(v))
