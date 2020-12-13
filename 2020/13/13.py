
def part1(earliest, buses):
    int_buses = [int(b) for b in buses.split(',') if b != 'x']

    i = earliest
    candidate = 0
    while candidate == 0:
        for b in int_buses:
            if i % b == 0:
                candidate = b
                break
        i += 1
    
    return (i-earliest) * candidate

def part2(buses):
    indices = {}
    for i, b in enumerate(buses):
        if b == 'x':
            continue
        number = int(b)
        indices[number] = i

    no_x = [int(b) for b in buses if b != 'x']

    prev = no_x[0]
    factor = prev
    for n in no_x[1:]:
        step = 0
        while True:
            x1 = prev + step
            x2 = x1 + indices[n]
            if x2 % n == 0:
                # print("Found:", x1)
                prev = x1
                factor *= n
                break
            step += factor
    return x1

def main():
    with open('input', 'r') as f:
        earliest, buses = [l.strip() for l in f.readlines()]

    print("Part 1:", part1(int(earliest), buses)) # 410
    print("Part 2:", part2(buses.split(','))) # 600691418730595

main()