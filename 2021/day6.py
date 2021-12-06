import collections

data = [int(i) for i in open('input').read().split(',')]


spawn = collections.defaultdict(int)
for n in data:
    spawn[n] += 1


def run(n)
    for _ in range(80):
        new_spawn = collections.defaultdict(int)
        for i in range(8, -1, -1):
            if i == 0:
                new_spawn[8] += spawn[i]
                new_spawn[6] += spawn[0]
            else:
                new_spawn[i-1] = spawn[i]
        spawn = new_spawn

    print(sum(spawn.values()))


run(80)
run(256)
