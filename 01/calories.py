
def parse_calories(file):
    with open(file) as f:
        for line in f:
            yield line

def main():
    e = []
    b = []
    for line in parse_calories('calories.txt'):
        if line not in ['\r', '', '\n']:
            b.append(int(line.strip()))
        else:
            print(line)
            e.append(b)
            b = []

    elves = {}

    for i,v in enumerate(e):
        elves[f"elf_{i}"] = {"calories": v, "total": sum(v)}


    sorted_elves = sorted(elves.items(), key=lambda x: x[1]["total"], reverse=True)

    top = [sorted_elves[0], sorted_elves[1], sorted_elves[2]]

    print(sum(top[0][1]["calories"]) + sum(top[1][1]["calories"]) + sum(top[2][1]["calories"]))
    
    print(sorted_elves[0])

if __name__ == "__main__":
    main()

