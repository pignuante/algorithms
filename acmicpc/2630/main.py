def counter(x, y, n):
    color = paper[y][x]
    for xx in range(x, x + n):
        for yy in range(y, y + n):
            if color != paper[yy][xx]:
                counter(x, y, n // 2)
                counter(x, y + n // 2, n // 2)
                counter(x + n // 2, y, n // 2)
                counter(x + n // 2, y + n // 2, n // 2)
                return

    if color == 0:
        result.append(0)
    else:
        result.append(1)


n = int(input())
paper = []
result = []
for _ in range(n):
    paper.append(list(map(int, input().split())))

counter(0, 0, n)
print(result.count(0))
print(result.count(1))
