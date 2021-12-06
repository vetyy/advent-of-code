data = [i for i in open("input")]


def run(part):
    pos = 0
    depth = 0
    aim = 0
    for item in data:
        a, b = item.split()

        if a == "forward":
            pos += int(b)
            if part == 2:
                depth += int(b) * aim
        elif a == "down":
            if part == 2:
                aim += int(b)
            else:
                depth += int(b)
        elif a == "up":
            if part == 2:
                aim -= int(b)
            else:
                depth -= int(b)

    print(pos * depth)


run(1)
run(2)
