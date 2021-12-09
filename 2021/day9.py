data = [[int(j) for j in list(i.strip())] for i in open("input")]

part1 = 0
for i in range(len(data)):
    for j in range(len(data[0])):
        lower = True
        for x, y in [[1, 0], [0, 1], [0, -1], [-1, 0]]:
            if 0 <= i + x < len(data) and 0 <= j + y < len(data[0]):
                if data[i + x][j + y] <= data[i][j]:
                    lower = False
        if lower:
            part1 += data[i][j] + 1

print(part1)

basins = []
visited = set()
for i in range(len(data)):
    for j in range(len(data[0])):
        if (i, j) in visited or data[i][j] == 9:
            continue

        size = 0
        search = [(i, j)]
        while search:
            i, j = search.pop()
            if (i, j) in visited:
                continue

            visited.add((i, j))
            size += 1
            for x, y in [[1, 0], [0, 1], [0, -1], [-1, 0]]:
                if 0 <= i + x < len(data) and 0 <= j + y < len(data[0]):
                    if data[i + x][j + y] != 9:
                        search.append((i + x, j + y))
        basins.append(size)

part2 = 1
for i in sorted(basins)[-3:]:
    part2 *= i
print(part2)
