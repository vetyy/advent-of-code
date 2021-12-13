data = [list(map(int, i.strip())) for i in open("input")]

directions = [
    [0, 1],
    [1, 0],
    [1, 1],
    [0, -1],
    [-1, 0],
    [-1, -1],
    [-1, 1],
    [1, -1],
]

flashes = 0
steps = 0
while True:
    flashed = set()
    for i in range(len(data)):
        for j in range(len(data[0])):
            search = [(i, j)]
            while search:
                x, y = search.pop()
                if (x, y) in flashed:
                    continue

                if data[x][y] != 9:
                    data[x][y] += 1
                    continue

                data[x][y] = 0
                flashed.add((x, y))
                if steps < 100:
                    flashes += 1

                for d in range(8):
                    xx = x + directions[d][0]
                    yy = y + directions[d][1]
                    if 0 <= xx < len(data) and 0 <= yy < len(data[0]):
                        search.append((xx, yy))

    steps += 1
    if len(flashed) == len(data) * len(data[0]):
        break


print(flashes)
print(steps)
