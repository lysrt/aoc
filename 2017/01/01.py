with open('input', 'r') as f:
    ints = list(f.read())
    r1 = sum([int(c) for i, c in enumerate(ints) if c == ints[(i+1) % len(ints)]])
    r2 = sum([int(c) for i, c in enumerate(ints) if c == ints[(i+len(ints)/2) % len(ints)]])
    print(r1, r2)