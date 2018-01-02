nodes = []

def node_from_name(name):
    return [n for n in nodes if n.name == name][0]

class Node:
    def __init__(self, name, weight):
        self.name = name
        self.weight = weight
        self.children = []
        self.balanced = False
    
    def __str__(self):
        return self.name + "(" + str(self.weight) + ") - " + str(self.children)

    def compute_total_weight(self):
        children_nodes = [node_from_name(child) for child in self.children]
        children_weights = []
        for c in children_nodes:
            cw = c.compute_total_weight()
            children_weights.append(cw)

        if not children_weights:
            # A leaf is balanced
            self.balanced = True
        else:
            # Is this node balanced?
            self.balanced = all(x==children_weights[0] for x in children_weights)
            
        self.t_weight = self.weight + sum(children_weights)
        return self.t_weight

def get_node(t):
    p = t.strip().split(' ')
    node = Node(p[0], int(p[1].strip('()')))
    return node

def get_children(t):
    ps = t.strip().split(',')
    return [p.strip() for p in ps]

def parse_line(l):
    parts = l.split('->')
    node = get_node(parts[0])
    if len(parts) > 1:
        node.children = get_children(parts[1])
    return node

# Part One
def get_root():
    all_children = []
    for node in nodes:
        for child in node.children:
            if not child in all_children:
                all_children.append(child)

    for node in nodes:
        if node.name not in all_children:
            return str(node.name)

def part_two(root_name):
    root = node_from_name(root_name)
    root.compute_total_weight() # Populate t_weight and balanced values for each node

    for n in nodes:
        if not n.balanced:
            child_nodes = [node_from_name(c) for c in n.children]

            one_children_unbalanced = False
            for child in child_nodes:
                if not child.balanced:
                    one_children_unbalanced = True
                    break
            if one_children_unbalanced:
                #One child is not balanced, check another node until we find it
                continue

            # All children are balanced, need to change the weight of one of them
            weights = {}
            for c in child_nodes:
                if not c.t_weight in weights:
                    weights[c.t_weight] = 1
                else:
                    weights[c.t_weight] += 1

            # Take the weight appearing the less
            wrong_weight = min(weights, key=weights.get) # With the min occurence
            correct_weight = max(weights, key=weights.get) # With the max occurence
            
            # Find the wrong child and correct it
            wrong_child = [c for c in child_nodes if c.t_weight == wrong_weight][0]
            corrected_child_weight = wrong_child.weight - (wrong_weight - correct_weight)
            return corrected_child_weight

with open('input', 'r') as f:
    lines = f.read().splitlines()

nodes = [parse_line(l) for l in lines]

root = get_root()
print(root)
print(part_two(root))