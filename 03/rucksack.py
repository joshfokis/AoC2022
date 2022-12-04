def inventory_priority(inv):
    scale = {"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24, "y": 25, "z": 26, "A": 27, "B": 28, "C": 29, "D": 30, "E": 31, "F": 32, "G": 33, "H": 34, "I": 35, "J": 36, "K": 37, "L": 38, "M": 39, "N": 40, "O": 41, "P": 42, "Q": 43, "R": 44, "S": 45, "T": 46, "U": 47, "V": 48, "W": 49, "X": 50, "Y": 51, "Z": 52}
    if len(inv) == 2:
        for i in inv[0]:
            if i in inv[1]:
                return scale[i]
    elif len(inv) == 3:
        for i in inv[0]:
            if i in inv[1] and i in inv[2]:
                return scale[i]

def main():
    inv = []
    with open("rucksack.txt") as f:
        for line in f:
            inv.append(line.strip())

    part1 = sum(inventory_priority([set(i[:len(i)//2]),set(i[len(i)//2:])]) for i in inv)

    print("Part 1:", part1)

    new_inv = []

    for i in range(0, len(inv), 3):
        new_inv.append(inventory_priority([inv[i], inv[i+1],inv[i+2]]))
       
    part2 = sum(new_inv)
    print("Part 2:", part2)

if __name__ == "__main__":
    main()
