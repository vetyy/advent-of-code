import collections

data = [i.strip() for i in open("input")]

template = data[0]
pairs = dict([i.split(" -> ") for i in data[2:]])


def run(iters):
    counter = collections.defaultdict(int)
    pcount = collections.defaultdict(int)
    for i in range(len(template) - 1):
        pcount[template[i] + template[i + 1]] += 1
        counter[template[i]] += 1

    for _ in range(iters):
        new_pcount = collections.defaultdict(int)
        for o in pcount:
            new_pcount[o[0] + pairs[o]] += pcount[o]
            new_pcount[pairs[o] + o[1]] += pcount[o]
            counter[pairs[o]] += pcount[o]
        pcount = new_pcount

    print(max(counter.values()) - min(counter.values()))


run(10)
run(40)
