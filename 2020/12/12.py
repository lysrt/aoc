import numpy as np

class Ship():
    def __init__(self, waypoint):
        self.pos = np.array([0, 0])
        self.waypoint = waypoint

    def move(self, action, value):
        if action == 'N':
            self.pos[1] += value
        elif action == 'S':
            self.pos[1] -= value
        elif action == 'E':
            self.pos[0] += value
        elif action == 'W':
            self.pos[0] -= value
        elif action == 'R':
            self.rotate(-value)
        elif action == 'L':
            self.rotate(value)
        elif action == 'F':
            self.forward(value)

    def move2(self, action, value):
        if action == 'N':
            self.waypoint[1] += value
        elif action == 'S':
            self.waypoint[1] -= value
        elif action == 'E':
            self.waypoint[0] += value
        elif action == 'W':
            self.waypoint[0] -= value
        elif action == 'R':
            self.rotate(-value)
        elif action == 'L':
            self.rotate(value)
        elif action == 'F':
            self.forward(value)

    def rotate(self, value):
        r090 = np.array([[0, -1],[1, 0]])
        r180 = np.array([[-1, 0],[0, -1]])
        r270 = np.array([[0, 1],[-1, 0]])

        if value == 90 or value == -270:
            self.waypoint = r090.dot(self.waypoint)
        if value == 180 or value == -180:
            self.waypoint = r180.dot(self.waypoint)
        if value == 270 or value == -90:
            self.waypoint = r270.dot(self.waypoint)

    def forward(self, value):
        self.pos += self.waypoint * value

def part1(instructions):
    waypoint = np.array([1, 0]) # E
    ship = Ship(waypoint)
    for i in instructions:
        ship.move(i[0], i[1])
    return abs(ship.pos[0]) + abs(ship.pos[1])

def part2(instructions):
    waypoint = np.array([10, 1])
    ship = Ship(waypoint)
    for i in instructions:
        ship.move2(i[0], i[1])
    return abs(ship.pos[0]) + abs(ship.pos[1])

def main():
    with open('input', 'r') as f:
        instructions = [(line[0], int(line.strip()[1:])) for line in f.readlines()]

    print("Part 1:", part1(instructions))
    print("Part 2:", part2(instructions))

main()