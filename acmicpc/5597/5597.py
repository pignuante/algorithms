students = set([])
for s in range(28):
    n = int(input())
    students.update([n])
print(*(students ^ set(range(1, 31))))
