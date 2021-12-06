data = [i.strip() for i in open('input')]
numbers = [int(i) for i in data[0].split(',')]

bingo_boards = []
bingo_board = []

for line in data[1:]:
    if line:
        bingo_board.append([int(i) for i in line.split()])
    else:
        if bingo_board:
            bingo_boards.append(bingo_board)
        bingo_board = []
else:
    bingo_boards.append(bingo_board)

for board in bingo_boards:
    columns = []
    for i in range(len(board[0])):
        column = [row[i] for row in board]
        columns.append(column)
    board.extend(columns)


def run(n_wins):
    wins = []
    for n in numbers:
        for i, board in enumerate(bingo_boards):
            if i in wins:
                continue

            for row in board:
                if n in row:
                    row.remove(n)
                if not row:
                    wins.append(i)
                    break

            if len(wins) == n_wins:
                s = 0
                for row in board[:len(board)//2]:
                    s += sum(row)
                print(n * s)
                return

run(1)
run(len(bingo_boards))
