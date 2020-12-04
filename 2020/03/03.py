
def solve1(board):
    return descend(board, 3, 1)

def solve2(board):
    a = descend(board, 1, 1)
    b = descend(board, 3, 1)
    c = descend(board, 5, 1)
    d = descend(board, 7, 1)
    e = descend(board, 1, 2)

    return a * b * c * d * e

def descend(board, slope_x, slope_y):
    trees_count = 0
    x, y = 0, 0

    while y < len(board):
        tree = access(board, x, y)
        if tree == '#':
            trees_count += 1
        x += slope_x
        y += slope_y
    
    return trees_count

def access(board, x, y):
    width = len(board[0])
    x = x % width
    return board[y][x]

def main():
    with open('input', 'r') as f:
        lines = f.readlines()

    board = [x.strip() for x in lines]
    
    print(solve1(board))
    print(solve2(board))

main()