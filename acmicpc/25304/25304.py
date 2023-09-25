X = int(input())
N = int(input())
for _ in range(N):
    a, b = map(int, input().split())
    X -= a * b
    
print("No" if X else "Yes")
