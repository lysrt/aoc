class Node:
    def __init__(self, node):
        self.name = node
        self.children = []
        self.parent = None
        self.orbits = 0

    def num_orbits(self, n=0):
        self.orbits = n
        if not self.children:
            return
        for c in self.children:
            c.num_orbits(n+1)
        
def create_orbits(lines):
    orbits = {}

    for left, right in [line.split(')') for line in lines]:
        if not left in orbits:
            orbits[left] = Node(left)
        if not right in orbits:
            orbits[right] = Node(right)
        orbits[left].children.append(orbits[right])
        orbits[right].parent = orbits[left]
    return orbits

with open('input', 'r') as f:
    lines = f.read().splitlines()

orbits = create_orbits(lines)

print("Part one")
orbits['COM'].num_orbits()
part1 = sum(o.orbits for o in orbits.values())
print(part1)

def get_parents(name):
    parents = set()
    current = orbits[name]
    while current.parent != None:
        parents.add(current.parent.name)
        current = current.parent
    return parents

yp = get_parents('YOU')
sp = get_parents('SAN')
intersect = yp.intersection(sp)

print("Part two")
part2 = len(yp.difference(intersect)) + len(sp.difference(intersect))
print(part2)