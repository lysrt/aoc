from collections import Counter

def explode_line(line):
    policy, password = line.split(': ')
    bounds, letter = policy.split(' ')
    low, high = bounds.split('-')
    return int(low), int(high), letter, password

def solve1(lines):
    valid = 0
    for line in lines:
        low, high, letter, password = explode_line(line.strip())
        count = Counter(password)
        occurences = count[letter]
        if occurences < low or occurences > high:
            continue # Invalid
        else:
            valid += 1
    return valid

def solve2(lines):
    valid = 0
    for line in lines:
        low, high, letter, password = explode_line(line.strip())
        letter_low = password[low-1]
        letter_high = password[high-1]
        match = 0
        if letter_low == letter:
            match += 1
        if letter_high == letter:
            match += 1
        if match == 1:
            valid += 1
    return valid

def main():
    with open('input', 'r') as f:
        lines = f.readlines()

    print(solve1(lines))
    print(solve2(lines))

main()