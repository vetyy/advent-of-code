data = [int(i) for i in open("input").read().split(",")]


def sum_n(n):
    return n * (n + 1) // 2


def run(part):
    fuels = {}
    for i in range(min(data), max(data) + 1):
        fuels[i] = 0
        for n in data:
            fuels[i] += abs(n - i) if part == 1 else sum_n(abs(n - i))

    print(fuels[min(fuels, key=fuels.get)])


run(1)
run(2)
