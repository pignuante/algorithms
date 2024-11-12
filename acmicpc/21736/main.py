from collections import deque


def BFS(graph: list, visited: list, start: tuple):
    yx = ((1, 0), (-1, 0), (0, -1), (0, 1))
    visited[start[0]][start[1]] = True
    queue = deque([start])
    answer = 0
    while queue:
        y, x = queue.popleft()

        for yy, xx in yx:
            dy = y + yy
            dx = x + xx

            if 0 <= dy < N and 0 <= dx < M and visited[dy][dx] == False and graph[dy][dx] != "X":
                if graph[dy][dx] == "P":
                    answer += 1
                visited[dy][dx] = True
                queue.append((dy, dx))
    return answer


N, M = list(map(int, input().split()))
graph = []
for y in range(N):
    graph.append(list(map(str, input())))

visited = [[False for _ in range(M)] for _ in range(N)]
start = None
for y in range(N):
    for x in range(M):
        if graph[y][x] == "I":
            start = (y, x)

a = BFS(graph, visited, start)

if a:
    print(a)
else:
    print("TT")
