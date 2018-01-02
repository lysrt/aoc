with open('input', 'r') as f:
    content = f.read()

    # One liner
    print(sum([max(r) - min(r) for r in [[int(i) for i in l.split('\t')] for l in content.splitlines()]]))

    # Part one
    s = 0
    for line in content.splitlines():
        ints = [int(i) for i in line.split('\t')]
        s += max(ints) - min(ints)
    print(s)

    # Part two
    s = 0
    for line in content.splitlines():
        ints = [int(i) for i in line.split('\t')]
        s += [(x / y) for x in ints for y in ints if x != y and x > y and x % y == 0][0]
    print(s)