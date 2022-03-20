S = input()
start = "A"
result = 0

for s in S:
    left = ord(start) - ord(s)
    right = ord(s) - ord(start)

    if left < 0:
        left += 26
    elif right < 0:
        right += 26

    result += min(left, right)
    start = s

print(result)
