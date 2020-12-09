def is_sum_of_2(n, numbers):
    d = set()
    for i in numbers:
        if i in d:
            return True
        else:
            d.add(n-i)
    return False

def find_problematic_number(numbers, preamble_count):
    for i in range(len(numbers)-preamble_count):
        preamble = numbers[i:i+preamble_count]
        number = numbers[i+preamble_count]
        if not is_sum_of_2(number,  preamble):
            return number

def find_sum_window(n, numbers):
    for length in range(2, len(numbers)):
        for start in range(len(numbers)-length):
            window = numbers[start:start+length]
            if n == sum(window):
                return min(window) + max(window)

def main():
    with open('input', 'r') as f:
        numbers = [int(x.strip()) for x in f.readlines()]

    wrong_number = find_problematic_number(numbers, preamble_count=25)
    print("Part 1:", wrong_number)
    print("Part 2:", find_sum_window(wrong_number, numbers))

main()