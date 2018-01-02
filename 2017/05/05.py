import time

def jump(lines, update):
    current, nb_jumps = 0, 0
    while current >= 0 and current < len(lines):
        steps = lines[current]
        update(lines, current)
        # Jump
        current += steps
        nb_jumps += 1
    return nb_jumps

def partOne(lines, current):
    lines[current] += 1

def partTwo(lines, current):
    lines[current] += -1 if lines[current] >= 3 else 1

with open('input', 'r') as f:
    lines = [int(line) for line in f.read().splitlines()]
    # f.readlines() = list(f) = [l for l in f] # Keeps \n at the end of lines

    start_time = time.time()
    print(jump(lines[:], partOne))
    print("1: %s ms" % float((time.time() - start_time)*1000))

    start_time = time.time()
    print(jump(lines[:], partTwo))
    print("2: %s ms" % float((time.time() - start_time)*1000))
