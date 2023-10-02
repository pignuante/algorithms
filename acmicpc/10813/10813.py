N, M  = map(int, input().split())
l = [x for x in range(1, N+1)]
for _ in range(M):
    i, j = map(int, input().split())
    i, j = i-1, j-1
    l[i], l[j] = l[j], l[i]
print(*l)
