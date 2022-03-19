from sys import stdin
input = stdin.readline

A, B, C, M = list(map(int, input().split()))
hour = 24
tired, work = 0, 0

if A > M:
    print(0)
else:
    for _ in range(hour):
        if tired + A <= M:
            tired += A
            work += B
        else:
            tired = tired - C if tired - C >= 0 else 0
    print(work)
