mat = [
    [1, 2, 3],
    [1, 2, 3],
    [1, 2, 3],
]
m = 3
n = 3
start = None
for i in range(m):
    if start:
        break
    for j in range(n):
        if mat[i][j] == 1:
            start = (i, j)
            break
print(start)
