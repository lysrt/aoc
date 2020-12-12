import numpy as np

class Ship():
    def __init__(self, waypoint):
        self.pos = np.array([0, 0])
        self.waypoint = np.array(waypoint)

    def navigate(self, instructions, move_pos):
        for action, value in instructions:
            self.move(action, value, move_pos)
        return abs(self.pos[0]) + abs(self.pos[1])

    def move(self, action, value, move_pos):
        moved_attribute = self.pos if move_pos else self.waypoint
        if action == 'N':
            moved_attribute[1] += value
        elif action == 'S':
            moved_attribute[1] -= value
        elif action == 'E':
            moved_attribute[0] += value
        elif action == 'W':
            moved_attribute[0] -= value
        elif action == 'R':
            self.rotate(-value)
        elif action == 'L':
            self.rotate(value)
        elif action == 'F':
            self.pos += self.waypoint * value # Forward

    def rotate(self, value):
        if value == 90 or value == -270:
            self.waypoint = np.array([[0, -1],[1, 0]]).dot(self.waypoint)
        if value == 180 or value == -180:
            self.waypoint = np.array([[-1, 0],[0, -1]]).dot(self.waypoint)
        if value == 270 or value == -90:
            self.waypoint = np.array([[0, 1],[-1, 0]]).dot(self.waypoint)

def main():
    with open('input', 'r') as f:
        instructions = [(line[0], int(line.strip()[1:])) for line in f.readlines()]

    print("Part 1:", Ship([1, 0]).navigate(instructions, move_pos=True))
    print("Part 2:", Ship([10, 1]).navigate(instructions, move_pos=False))

main()