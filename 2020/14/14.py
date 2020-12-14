
class Memory():
    def __init__(self):
        self.mask = ''
        self.mask0 = 0
        self.mask1 = 0
        self.memory = {}

    def parse_mask(self, instruction):
        self.mask = instruction.replace("mask = ", '')
        self.mask0 = int(self.mask.replace('X', '1'), 2)
        self.mask1 = int(self.mask.replace('X', '0'), 2)

    def parse_mem(self, instruction):
        str_addr, str_value = instruction.split(" = ")
        return int(str_addr.replace("mem[", '').replace(']', '')), int(str_value)

    def apply_mask0(self, value):
        return self.mask0 & value

    def apply_mask1(self, value):
        return self.mask1 | value

    def apply_mask(self, value):
        return ''.join([value[i] if self.mask[i] == '0' else self.mask[i] for i in range(len(value))])

    def mask_value(self, addr, value):
        self.memory[addr] = self.apply_mask1(self.apply_mask0(value))

    def mask_address(self, addr, value):
        binary_address = f'{addr:036b}'
        x_masked_address = self.apply_mask(binary_address)

        x_count = x_masked_address.count('X')
        candidates_for_x = [f'{x:0{x_count}b}' for x in range(2 ** x_count)]

        binary_addresses_to_write = []
        for candidate in candidates_for_x:
            address = x_masked_address
            for char in candidate:
                address = address.replace('X', char, 1)
            binary_addresses_to_write.append(address)

        for int_address in [int(x, 2) for x in binary_addresses_to_write]:
            self.memory[int_address] = value

def run(instructions, part):
    mem = Memory()
    for ins in instructions:
        if ins.startswith('mask'):
            mem.parse_mask(ins)
        else:
            if part == 1:
                mem.mask_value(*mem.parse_mem(ins))
            elif part == 2:
                mem.mask_address(*mem.parse_mem(ins))

    return sum(mem.memory.values())

def main():
    with open('input', 'r') as f:
        instructions = [l.strip() for l in f.readlines()]

    print("Part 1:", run(instructions, 1))
    print("Part 2:", run(instructions, 2))

main()