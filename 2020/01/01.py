
def solve1(ints):
    for i in range(len(ints)):
        j = i + 1
        while j < len(ints):
            if ints[i] + ints[j] == 2020:
                return ints[i] * ints[j]
            j += 1

def solve2(ints):
    for i in range(len(ints)):
        j = i + 1
        while j < len(ints):
            k = j + 1
            while k < len(ints):
                if ints[i] + ints[j] + ints[k] == 2020:
                    return ints[i] * ints[j] * ints[k]
                k += 1
            j += 1

def main():
    with open('input', 'r') as f:
        ints = [int(x) for x in f.readlines()]

    print(solve1(ints))
    print(solve2(ints))

main()