from collections import Counter
from functools import reduce

def part1(numbers):
    s = sorted(numbers + [0, max(numbers) + 3])
    c = Counter([b - a for a, b in zip(s, s[1:])])
    return c[1] * c[3]

def ways_to(n, numbers):
    parents = [x for x in numbers if 1 <= n-x <= 3]
    return reduce(lambda a, b: a+b, [ways_to(x, numbers) for x in parents]) if parents else 1

def memoize(func):
    memo = {}
    def stub(x, n):
        if not x in memo:
            memo[x] = func(x, n)
        return memo[x]
    return stub

ways_to = memoize(ways_to)

def main():
    with open('input', 'r') as f:
        numbers = [int(x.strip()) for x in f.readlines()]

    print("Part 1:", part1(numbers))
    print("Part 2:", ways_to(max(numbers) + 3, numbers + [0]))
main()