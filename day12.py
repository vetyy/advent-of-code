import collections

data = [i.strip() for i in open("input2")]


caves = collections.defaultdict(list)
for line in data:
    x, y = line.split("-")
    caves[x].append(y)
    caves[y].append(x)


def run(part1):
    visited = set()
    search = [[("start", c), 0] for c in caves["start"]]
    while search:
        path, vcounter = search.pop()
        for c in caves[path[-1]]:
            if c == "start":
                continue

            counter = vcounter
            if c in path and c >= "a":
                if counter > (-1 if part1 else 0):
                    continue
                counter += 1

            new_path = path + (c,)
            if c == "end" and new_path not in visited:
                visited.add(new_path)
                continue

            search.append([new_path, counter])

    print(len(visited))


run(True)
run(False)
