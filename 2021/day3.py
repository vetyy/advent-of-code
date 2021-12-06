import collections

data = [i.strip() for i in open('input')]


x = {}
for item in data:
    for j, i in enumerate(item):
        if j not in x:
            x[j] = collections.defaultdict(int)
        x[j][i] += 1


gamma = ''
epsilon = ''
for key, value in x.items():
    if value['0'] > value['1']:
        gamma += '0'
        epsilon += '1'
    else:
        gamma += '1'
        epsilon += '0'


print(int(gamma, 2) * int(epsilon, 2))


one_data = list(data)
zero_data = list(data)

for i in range(len(data[0])):
    if len(one_data) > 1:
        zero_len, one_len = 0, 0
        for data in one_data:
            if data[i] == '0':
                zero_len += 1
            else:
                one_len += 1
        if one_len >= zero_len:
            one_data = [data for data in one_data if data[i] == '1']
        else:
            one_data = [data for data in one_data if data[i] == '0']

    if len(zero_data) > 1:
        zero_len, one_len = 0, 0
        for data in zero_data:
            if data[i] == '0':
                zero_len += 1
            else:
                one_len += 1
        if one_len >= zero_len:
            zero_data = [data for data in zero_data if data[i] == '0']
        else:
            zero_data = [data for data in zero_data if data[i] == '1']


print(int(zero_data[0],2)*int(one_data[0],2))
