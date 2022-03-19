from sys import stdin
input = stdin.readline

T = int(input())
for _ in range(T):
    C = int(input())
    quarter = C//25
    dime = (C % 25)//10
    nickel = ((C % 25) % 10)//5
    penny = (((C % 25) % 10) % 5)//1

    print(quarter, dime, nickel, penny)
