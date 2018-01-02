class Operation:
    def __init__(self, register, action, action_value, cond_register, cond_symbol, cond_value):
        self.register = register
        self.action = action
        self.action_value = action_value
        self.cond_register = cond_register
        self.cond_symbol = cond_symbol
        self.cond_value = cond_value

    def run(self, registers, all_time_max):
        # Check if condition is fullfilled
        condition_ok = False
        actual = registers[self.cond_register]
        expected = self.cond_value
        
        if self.cond_symbol == '==':
            if actual == expected:
                condition_ok = True
        elif self.cond_symbol == '!=':
            if actual != expected:
                condition_ok = True
        elif self.cond_symbol == '<':
            if actual < expected:
                condition_ok = True
        elif self.cond_symbol == '<=':
            if actual <= expected:
                condition_ok = True
        elif self.cond_symbol == '>':
            if actual > expected:
                condition_ok = True
        elif self.cond_symbol == '>=':
            if actual >= expected:
                condition_ok = True
        else:
            print("Unknown symbol: " + self.cond_symbol)
            raise 
        
        if not condition_ok:
            return all_time_max

        if self.action == 'inc':
            registers[self.register] += self.action_value
        else:
            registers[self.register] -= self.action_value

        if registers[self.register] > all_time_max:
            all_time_max = registers[self.register]
        return all_time_max
        
def parse_line_to_operation(line):
    parts = line.split()

    identifier = parts[0]
    action = parts[1]
    if action != 'inc' and action != 'dec':
        print('Wrong action: ' + action)
        return None
    number = int(parts[2])
    if parts[3] != 'if':
        print('Missing if: ' + parts[3])
        return None
    identifier2 = parts[4]
    symbol = parts[5]
    number2 = int(parts[6])

    return Operation(identifier, action, number, identifier2, symbol, number2)

with open('input', 'r') as f:
    lines = f.read().splitlines()

operations = [parse_line_to_operation(l) for l in lines]

# Initialize registers
registers = {}
for op in operations:
    if not op.register in registers:
        registers[op.register] = 0

# Run all operations
all_time_max = -2**32
for op in operations:
    all_time_max = op.run(registers, all_time_max)

# Display the largest value in the registers
m = 0
for value in registers.values():
    if value > m:
        m = value
print("Max: " + str(m))
print("All time max: " + str(all_time_max))
