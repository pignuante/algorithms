T = int(input())
for _ in range(T):
    x1, y1, r1, x2, y2, r2 = list(map(int, input().split()))
    dist = (x2-x1) * (x2-x1)+(y2-y1) * (y2-y1)
    rd = (r2-r1) * (r2-r1)
    rs = (r2+r1) * (r2+r1)
    C = "0"
    if dist == 0 and rd == 0:
        C = "-1"
    else:
        if dist in (rd, rs):
            C = "1"
        elif rd < dist < rs:
            C = "2"
    print(C)
