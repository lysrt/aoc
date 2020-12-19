
def match(rules, rule_id, message, skip):
    if rules[rule_id][0] == '"':
        return {skip + 1} if skip < len(message) and rules[rule_id][1] == message[skip] else set()
    else:
        matching = set()
        for sub_rule in rules[rule_id]:
            buffer = {skip}
            for part in sub_rule:
                temp = set()
                for loc in buffer:
                    temp |= match(rules, part, message, loc)
                buffer = temp
            matching |= buffer
        return matching

def main():
    with open('input', 'r') as f:
        raw_rules, messages = [x.splitlines() for x in f.read().split("\n\n")]
        rules = {}
        for text in raw_rules:
            index, rule = text.split(': ')
            if '"' in rule:
                rules[index] = rule
            else:
                rules[index] = [seq.split(' ') for seq in rule.split (' | ')]

    results = [len(m) in match(rules, '0', m, 0) for m in messages]
    print("Part 1:", results.count(True))

    rules['8'] = [['42'], ['42', '8']]
    rules['11'] = [['42', '31'], ['42', '11', '31']]

    results = [len(m) in match(rules, '0', m, 0) for m in messages]
    print("Part 2:", results.count(True))

main()