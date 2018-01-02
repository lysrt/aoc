def partOne(val):
    # The bottom right value of each square is an odd power of 2 (9, 25, 49, 81, ...)
    # Find which square the number is in by getting the next upper odd power of 2
    i = 1
    while i*i < val:
        i += 2
    square = i/2

    # The number "square" is at distance "square * 2" from the center
    # All numbers in this square are at distance D such as: square >= D >= 2 * square

    # (i * i) is the value of the bottom right corner
    # We can decrement from (i * i) to get to val
    # The distance D of each step number will go down and up in its bounds
    # A modulo on (i * i) - v will give us the distance difference between d(square * 2) and d(value)
    return (square * 2) - ((i * i - val) % square)

def partTwo(val):
    # Build the spiral
    values = [(0, 0, 1)]
    directions = [(1, 0), (0, 1), (-1, 0), (0, -1)] # Right, Up, Left, Down

    steps = 0
    directionIndex = 0
    while True:
        # Increase the step (segment length) every two segment        
        if directionIndex % 2 == 0:
            steps += 1

        # Compute one segment
        for step in range(0, steps):
            currentX, currentY, _ = values[-1]
            dX, dY = directions[directionIndex % 4]
            
            nextX = currentX + dX
            nextY = currentY + dY

            value = sum([v[2] for v in values if abs(nextX - v[0]) <= 1 and abs(nextY - v[1]) <= 1])

            if value > val:
                return value

            next = (nextX, nextY, value)
            values.append(next)

        # Then change direction
        directionIndex += 1

val = 347991 # answer: 480
print(partOne(val))
print(partTwo(val))