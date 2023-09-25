A, B = map(int, input().split())
C = int(input())

cc = B + C
h = (A + cc//60)%24
m = cc%60
print(f"{h} {m}")
