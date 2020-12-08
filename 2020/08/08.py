class VM():
    def __init__(self, instructions):
        self.instructions = instructions
        self.acc = 0
        self.counter = 0
        self.seen = set()
    
    def get_current_instruction(self):
        self.seen.add(self.counter)
        op, arg = self.instructions[self.counter].split(' ')
        return op, int(arg)

    def run(self):
        while True:
            if self.counter >= len(self.instructions):
                return 0
            if self.counter in self.seen:
                return -1 # Infinite loop
            op, arg = self.get_current_instruction()
            if op == 'nop':
                self.counter += 1
            elif op == 'acc':
                self.acc += arg
                self.counter += 1
            elif op == 'jmp':
                self.counter += arg
            else:
                raise IndexError(f"Unknown operation: {op}")


def main():
    with open('input', 'r') as f:
        instructions = [x.strip() for x in f.readlines()]

    vm = VM(instructions)
    vm.run()
    print("Part 1:", vm.acc)

    alternative_instructions = []
    for i, instruction in enumerate(instructions):
        op, arg = instruction.split(' ')
        if op == 'acc':
            continue
        ins = instructions.copy()
        if op == 'nop':
            ins[i] = f'jmp {arg}'
        elif op == 'jmp':
            ins[i] = f'nop {arg}'
        alternative_instructions.append(ins)

    for ins in alternative_instructions:
        vm = VM(ins)
        code = vm.run()
        if code == 0:
            print("Part 2:", vm.acc)
            break

main()