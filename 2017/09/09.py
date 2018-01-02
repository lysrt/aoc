with open('input', 'r') as f:
    group_level = 0
    group_sum = 0
    garbage = False
    garbage_count = 0

    while True:
        c = f.read(1)
        if not c:
            break

        if c == '{':
            if not garbage:
                group_level += 1
                group_sum += group_level
            else:
                garbage_count += 1
        elif c == '}':
            if not garbage:
                group_level -= 1
            else:
                garbage_count += 1
        elif c == '<':
            if not garbage:
                garbage = True
            else:
                garbage_count += 1
        elif c == '>':
            garbage = False
        elif c == '!':
            f.read(1)
        else:
            if garbage:
                garbage_count += 1

    print(group_sum)
    print(garbage_count)