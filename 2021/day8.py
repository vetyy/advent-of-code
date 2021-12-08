data = [i for i in open("input")]


def run(part):
    digit_sum = 0
    for line in data:
        signals, digits = line.split("|")

        def get_key(val):
            for k, v in sig_map.items():
                if val == v:
                    return k

        sig_map = {}
        for sig in signals.split():
            if len(sig) == 2:
                sig_map[sig] = 1
            elif len(sig) == 4:
                sig_map[sig] = 4
            elif len(sig) == 3:
                sig_map[sig] = 7
            elif len(sig) == 7:
                sig_map[sig] = 8

        if part == 2:
            for sig in signals.split():
                sig_left = set(sig) - set(get_key(1))
                if len(sig) == 5 and len(sig_left) == 3:
                    sig_map[sig] = 3
                elif len(sig) == 6 and len(sig_left) == 5:
                    sig_map[sig] = 6

            for sig in signals.split():
                sig_left = set(sig) - set(get_key(4))
                if len(sig) == 5 and sig != get_key(3):
                    sig_map[sig] = 5 if len(sig_left) == 2 else 2

                elif len(sig) == 6 and sig != get_key(6):
                    sig_map[sig] = 9 if len(sig_left) == 2 else 0

        sorted_sig_map = {}
        for k, v in sig_map.items():
            sorted_sig_map["".join(sorted(k))] = v

        if part == 1:
            for dig in digits.split():
                if "".join(sorted(dig)) in sorted_sig_map:
                    digit_sum += 1
        else:
            full_dig = ""
            for dig in digits.split():
                full_dig += str(sorted_sig_map["".join(sorted(dig))])
            digit_sum += int(full_dig)

    print(digit_sum)


run(1)
run(2)
