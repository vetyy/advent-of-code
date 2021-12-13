dots = set()
folds = []
for line in open("input"):
    if not line.strip():
        continue
    if line.startswith("fold"):
        folds.append(line.split()[2].split("="))
    else:
        dots.add(tuple(map(int, line.split(","))))


part1_done = False
for fold, n in folds:
    rows = max((i[0] for i in dots)) + 1
    columns = max((i[1] for i in dots)) + 1
    for i in range(rows):
        for j in range(columns):
            y = j
            x = i

            if (x, y) in dots:
                dots.remove((x, y))
                if fold == "x":
                    x = int(n) - (i - int(n))
                else:
                    y = int(n) - (j - int(n))
                dots.add((x, y))

    if not part1_done:
        part1_done = True
        print(len(dots))


for j in range(columns):
    for i in range(rows):
        print("#" if (i, j) in dots else ".", end="")
    print()
