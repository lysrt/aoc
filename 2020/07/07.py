
def parse_line(line):
    bag, raw_content = line.rstrip('.').split(" bags contain ")
    if raw_content == "no other bags":
        return bag, None

    parsed_content = [parse_component(x) for x in raw_content.split(",")]
    content = {bag: count for (bag, count) in parsed_content}
    return bag, content

def parse_component(raw):
    raw = right_replace(raw, "bags", '')
    raw = right_replace(raw, "bag", '').strip()
    count, bag = raw.split(' ', 1)
    return (bag, int(count))

def right_replace(string, old, new, max_replace=1):
    return new.join(string.rsplit(old, max_replace))

def search(bags, first_bag, string):
    if not bags[first_bag]:
        return False
    if string in bags[first_bag]:
        return True
    else:
        for b, _ in bags[first_bag].items():
            if search(bags, b, string):
                return True
    return False

def search_count(bags, string):
    if not string in bags:
        return 0
    if not bags[string]:
        return 0
    else:
        total_count = 0
        for b, c in bags[string].items():
            total_count += c
            total_count += (c * search_count(bags, b))

    return total_count

def main():
    with open('input', 'r') as f:
        lines = [parse_line(x.strip()) for x in f.readlines()]

    bags = {bag: content for bag, content in lines}

    count = sum([1 for x in bags.keys() if search(bags, x, "shiny gold")])
    print("Part 1:", count)

    count = search_count(bags, "shiny gold")
    print("Part 2:", count)

main()