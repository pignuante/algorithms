from collections import deque

n, m = list(map(int, input().split()))
graph = []
for y in range(n):
    graph.append(list(map(int, input().split())))


def BFS(graph: list, visited: list, start: tuple):
    yx = ((1, 0), (-1, 0), (0, -1), (0, 1))
    visited[start[0]][start[1]] = 0
    queue = deque([start])

    while queue:
        y, x = queue.popleft()
        for yy, xx in yx:
            dy = y + yy
            dx = x + xx

            if 0 <= dy < n and 0 <= dx < m:
                if visited[dy][dx] == -1 and graph[dy][dx] == 1:
                    visited[dy][dx] = visited[y][x] + 1
                    queue.append((dy, dx))


visited = [[-1] * m for y in range(n)]
start = None

for y in range(n):
    for x in range(m):
        if graph[y][x] == 2:
            start = (y, x)
        if graph[y][x] == 0:
            visited[y][x] = 0

BFS(graph, visited, start)

for y in range(n):
    for x in range(m):
        print(visited[y][x], end=" ")
    print()
