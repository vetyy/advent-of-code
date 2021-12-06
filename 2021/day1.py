data = [int(i) for i in open("input")]

def run(n):
    count = 0
    prev = sum(data[:n])
    for i, measurement in enumerate(data):
        sum_ = sum(data[i:i+n])
        if sum_ > prev:
            count += 1
        prev = sum_

    print(count)


run(1)
run(3)
