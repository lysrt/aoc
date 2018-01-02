with open('input', 'r') as f:
    ints = [int(i) for i in f.read().split()]

def partOne():
    pastStates = []
    nb_distributions = 0
    while True:
        if ints in pastStates:
            break
        pastStates.append(ints[:])

        amount = max(ints)
        i = ints.index(amount)
        ints[i] = 0
        
        # Distribute
        j = i+1 # Start after empty cell
        while amount > 0:
            ints[j%len(ints)] += 1
            amount -= 1
            j += 1
        nb_distributions += 1

    return nb_distributions

print(partOne())
print(partOne())