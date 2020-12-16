from itertools import takewhile
from functools import reduce

class Field:
    def __init__(self, raw):
        name, raw_rules = raw.split(': ')
        rules = raw_rules.split(' or ')
        self.name = name
        self.rules = []
        for r in rules:
            low, high = r.split('-')
            self.rules.append((int(low), int(high)))

    def validate(self, value):
        for rule in self.rules:
            if rule[0] <= value <= rule[1]:
                return True
        return False

class Ticket(list):
    def __init__(self, raw, fields):
        self.extend([int(x) for x in raw.split(',')])
        self.valid_fields_by_index = [{f.name for f in fields if f.validate(v)} for v in self]

    def error_rate(self, fields):
        error_values_sum = 0
        for index, ticket_value in enumerate(self):
            if not self.valid_fields_by_index[index]:
                error_values_sum += ticket_value
        return error_values_sum

    def is_valid(self, fields):
        for index in range(len(self)):
            if not self.valid_fields_by_index[index]:
                return False
        return True

def parse(file):
    raw_fields = [line.strip() for line in takewhile(lambda line: line != '\n', file)]
    next(file)
    raw_my_ticket = file.readline().strip()
    next(file)
    next(file)
    raw_tickets = [line.strip() for line in file]

    fields = [Field(raw) for raw in raw_fields]
    my_ticket = Ticket(raw_my_ticket, fields)
    tickets = [Ticket(t, fields) for t in raw_tickets]

    return fields, my_ticket, tickets

def main():
    with open('input', 'r') as f:
        fields, my_ticket, tickets = parse(f)

    # Part 1
    print("Part 1:", sum([t.error_rate(fields) for t in tickets]))

    # Part 2
    valid_tickets = [t for t in tickets if t.is_valid(fields)]

    field_id_candidates = {i: {f.name for f in fields} for i in range(len(fields))}
    for ticket in valid_tickets:
        for index in range(len(ticket)):
            field_id_candidates[index] &= ticket.valid_fields_by_index[index]

    field_names = ['' for _ in fields]

    while '' in field_names:
        single_candidates = {k: v.pop() for k, v in field_id_candidates.items() if len(v) == 1}
        for field_id, field_name in single_candidates.items():
            field_names[field_id] = field_name
            for other_candidate in field_id_candidates.values():
                if field_name in other_candidate:
                    other_candidate.remove(field_name)

    named_ticket = {name: value for name, value in zip(field_names, my_ticket)}
    print("Part 2:", reduce(lambda a, b: a*b, [v for name, v in named_ticket.items() if name.startswith("departure")]))

main()