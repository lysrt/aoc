from functools import reduce

class Group():
    def __init__(self, raw):
        self.persons = [set(x) for x in raw.split("\n")]

    def get_answers_count(self, aggregation):
        return len(reduce(aggregation, self.persons))

with open('input', 'r') as f:
    groups = [Group(x) for x in f.read().split("\n\n")]

print("Part 1:", sum([g.get_answers_count(lambda a, b: a|b) for g in groups]))
print("Part 2:", sum([g.get_answers_count(lambda a, b: a&b) for g in groups]))