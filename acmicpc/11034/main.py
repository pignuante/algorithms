from sys import stdin
input = stdin.readline

while True:
    try:
        A, B, C = list(map(int, input().split()))
    except ValueError:
        break
    else:
        print(max(B-A, C-B)-1)
