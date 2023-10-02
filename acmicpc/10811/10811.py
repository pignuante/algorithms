N, M  = map(int, input().split())
l = [x for x in range(N+1)]
for _ in range(M):
    i, j = map(int, input().split())
    l[i:j+1] = l[j:i-1:-1]
print(*l[1:])
