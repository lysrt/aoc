from collections import Counter

def lastindex(numbers, n):
    return len(numbers) - 1 - numbers[::-1].index(n)

def using_list(numbers, stop):
    said = numbers.copy()
    for i in range(len(numbers), stop):
        previous = said[i-1]
        if previous in said[:-1]:
            said.append(i-lastindex(said[:-1], previous)-1)
        else:
            said.append(0)
    return said[-1]

def using_dict(numbers, stop):
    said = {x: i for i, x in enumerate(numbers)}
    last_said = None

    for current_index in range(len(numbers), stop):
        if last_said is None:
            inserting = 0
        else:
            inserting = current_index-last_said-1

        if inserting in said:
            last_said = said[inserting]
        else:
            last_said = None

        said[inserting] = current_index

        # Pretty progress bar
        if current_index % 3000000 == 0:
            length = 20
            progress = int(current_index / stop * length)
            line = '[' + '-' * progress + ' ' * (length - progress) + ']'
            print(line)

    return inserting

def main():
    starting_numbers = '12,20,0,6,1,17,7'
    numbers = [int(x) for x in starting_numbers.split(',')]

    print("Part 1:", using_list(numbers, 2020)) # Too slow for part 2
    print("Part 1:", using_dict(numbers, 2020))
    print("Part 2:", using_dict(numbers, 30000000))

main()