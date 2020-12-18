import numpy as np
import itertools

def iterate(lines, dimensions):
    active_cells = set()
    for y in range(len(lines)):
        for x in range(len(lines[0])):
            if lines[y][x] == '#':
                active_cells.add(tuple([x, y] + (dimensions-2) * [0]))

    for _ in range(6):
        active_cells = step(active_cells)

    return len(active_cells)

def step(active_cells):
    result = set(active_cells)

    full_board = set(active_cells)
    for t in active_cells:
        full_board |= get_all_neighbours(t)

    for t in full_board:
        active_neighbours = len([n for n in get_all_neighbours(t) if n in active_cells])
        if t in active_cells:
            if not active_neighbours in [2, 3]:
                result.remove(tuple(t))
        else:
            if active_neighbours == 3:
                result.add(tuple(t))
    return result

def get_all_neighbours(cell):
    dimensions = len(cell)
    perms = set(itertools.combinations([-1,0,1] * dimensions, dimensions))
    perms.remove(tuple([0] * dimensions))
    return set([tuple(a) for a in (cell + np.array(list(perms)))])

def main():
    with open('input', 'r') as f:
        lines = [list(x.strip()) for x in f.readlines()]

    print("Part 1:", iterate(lines, 3))
    print("Part 2:", iterate(lines, 4))

main()