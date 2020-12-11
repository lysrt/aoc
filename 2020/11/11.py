from copy import deepcopy
from itertools import product
from functools import reduce

def see(rows, col, row, step, nearby_only):
    x, y = step(col, row)
    while x >= 0 and x < len(rows[0]) and y >= 0 and y < len(rows):
        current = rows[y][x]
        if current == 'L':
            return 0
        if current == '#':
            return 1
        if nearby_only and current == '.':
            return 0
        x, y = step(x, y)
    return 0

def change_seat(rows, col, row, nearby_only, threshold):
    if rows[row][col] == '.':
        return '.'

    angles = [a for a in product([-1, 0, 1], [-1, 0, 1]) if a != (0,0)]
    neighbours_count = reduce(lambda a, b: a+b, [see(rows, col, row, lambda x, y: (x+xi, y+yi), nearby_only) for (xi, yi) in angles])

    if rows[row][col] == 'L':
        return '#' if neighbours_count == 0 else 'L'
    if rows[row][col] == '#':
        return 'L' if neighbours_count >= threshold else '#'

def shuffle_seats(rows, nearby_only, threshold):
    while True:
        new_rows = deepcopy(rows)

        for row in range(len(rows)):
            for col in range(len(rows[0])):
                new_rows[row][col] = change_seat(rows, col, row, nearby_only, threshold)

        if new_rows == rows:
            return sum([row.count('#') for row in new_rows])

        rows = new_rows

def main():
    with open('input', 'r') as f:
        rows = [list(x.strip()) for x in f.readlines()]

    print("Part 1:", shuffle_seats(rows, nearby_only=True, threshold=4))
    print("Part 2:", shuffle_seats(rows, nearby_only=False, threshold=5))

main()