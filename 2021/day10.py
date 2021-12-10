data = [i.strip() for i in open("input")]

lookup_map = {
    "{": "}",
    "[": "]",
    "(": ")",
    "<": ">",
}

p1_score = {
    ")": 3,
    "]": 57,
    "}": 1197,
    ">": 25137,
}

p2_score = {
    "(": 1,
    "[": 2,
    "{": 3,
    "<": 4,
}

part1 = 0
part2 = []
for line in data:
    stack = []
    corrupted = False
    for ch in line:
        if ch in lookup_map.keys():
            stack.append(ch)
        else:
            ech = stack.pop()
            if lookup_map[ech] != ch:
                corrupted = True
                part1 += p1_score[ch]

    if not corrupted:
        score = 0
        for ch in stack[::-1]:
            score = score * 5 + p2_score[ch]
        part2.append(score)


print(part1)
print(sorted(part2)[len(part2) // 2])
