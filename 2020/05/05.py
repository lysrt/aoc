def get_id(line):
    row = int(line[:7].replace('F', '0').replace('B', '1'), base=2)
    column = int(line [-3:].replace('L', '0').replace('R', '1'), base=2)
    return row * 8 + column

def main():
    with open('input', 'r') as f:
        lines = [x.strip() for x in f.readlines()]

    ids = sorted([get_id(x) for x in lines])
    for x, y in zip(ids, ids[1:]):
        if y - x == 2:
            seat = x + 1
            break

    print("Part 1:", max(ids))
    print("Part 2:", seat)

main()