import collections

data = [i.strip() for i in open('input')]


def run(part):
    final = collections.defaultdict(int)
    for line in data:
        left, right = line.split(' -> ')

        x1,y1 = left.split(',')
        x2,y2 = right.split(',')
        x1, x2 = int(x1), int(x2)
        y1, y2 = int(y1), int(y2)
        dx = -1 if x1 > x2 else 1
        dy = -1 if y1 > y2 else 1

        if x1 == x2:
            while y1 != y2:
                final[f'{x1},{y1}'] += 1
                y1 += dy
            final[f'{x1},{y1}'] += 1

        elif y1 == y2:
            while x1 != x2:
                final[f'{x1},{y1}'] += 1
                x1 += dx
            final[f'{x1},{y1}'] += 1

        elif part == 2:
            while x1 != x2:
                final[f'{x1},{y1}'] += 1
                x1 += dx
                y1 += dy
            final[f'{x1},{y1}'] += 1

    count = 0
    for value in final.values():
        if value > 1:
            count += 1

    print(count)


run(1)
run(2)
