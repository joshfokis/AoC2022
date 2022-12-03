def main():
    win = {"Rock": "Scissors", "Paper": "Rock", "Scissors": "Paper"}
    lose = {"Rock": "Paper", "Paper": "Scissors", "Scissors": "Rock"}
    draw = {"Rock": "Rock", "Paper": "Paper", "Scissors": "Scissors"}
    legend = {"A": "Rock", "B": "Paper", "C": "Scissors", "X": "Rock", "Y": "Paper", "Z": "Scissors"}
    my_choice = { "Rock":"X", "Paper":"Y", "Scissors":"Z" }
    points = {"win": 6, "lose": 0, "draw": 3, "Rock": 1, "Paper": 2, "Scissors": 3}

    plantally = 0
    allwins = 0
    
    with open("rps.txt", "r") as f:
        for line in f:
            line = line.strip().split()
            
            if line in ["", " ", "\n", "\r", None, []]:
                continue
            if line[1] == "X":
                mc = win[legend[line[0]]]
                score = points["lose"] + points[win[legend[line[0]]]]
                plantally += score
            if line[1] == "Z":
                mc = lose[legend[line[0]]]
                score = points["win"] + points[lose[legend[line[0]]]]
                plantally += score
            if line[1] == "Y":
                score = points["draw"] + points[legend[line[0]]]
                plantally += score

            if win[legend[line[0]]] == legend[line[1]]:
                score = points["lose"] + points[legend[line[1]]]
                allwins += score
            elif lose[legend[line[0]]] == legend[line[1]]:
                score = points["win"] + points[legend[line[1]]]
                allwins += score
            elif draw[legend[line[0]]] == legend[line[1]]:
                score = points["draw"] + points[legend[line[1]]]
                allwins += score
            
    print(allwins)
    print(plantally)

if __name__ == "__main__":
    main()

