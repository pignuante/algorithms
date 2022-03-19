from sys import stdin

input = stdin.readline

A, B = input().split()

minValue = int(A.replace("6", "5")) + int(B.replace("6", "5"))
maxValue = int(A.replace("5", "6")) + int(B.replace("5", "6"))

print(minValue, maxValue)
