empty_sign = "_"
impossible_text = "Impossible"
field_lines = [[empty_sign, empty_sign, empty_sign],
               [empty_sign, empty_sign, empty_sign],
               [empty_sign, empty_sign, empty_sign]]


def print_field(lines):
    vert_line = "---------"
    hor_line = "|"

    print(vert_line)

    for line in lines:
        print(f"{hor_line} {line[0]} {line[1]} {line[2]} {hor_line}")

    print(vert_line)


def find_winner(lines):
    winner = []

    if lines[0][0] == lines[0][1] == lines[0][2]:
        if lines[0][0] != empty_sign:
            winner.append(lines[0][0])
    if lines[1][0] == lines[1][1] == lines[1][2]:
        if lines[1][0] != empty_sign:
            winner.append(lines[1][0])
    if lines[2][0] == lines[2][1] == lines[2][2]:
        if lines[2][0] != empty_sign:
            winner.append(lines[2][0])
    if lines[0][0] == lines[1][0] == lines[2][0]:
        if lines[0][0] != empty_sign:
            winner.append(lines[0][0])
    if lines[0][1] == lines[1][1] == lines[2][1]:
        if lines[0][1] != empty_sign:
            winner.append(lines[0][1])
    if lines[0][2] == lines[1][2] == lines[2][2]:
        if lines[0][2] != empty_sign:
            winner.append(lines[0][2])
    if lines[0][0] == lines[1][1] == lines[2][2]:
        if lines[0][0] != empty_sign:
            winner.append(lines[0][0])
    if lines[0][2] == lines[1][1] == lines[2][0]:
        if lines[0][2] != empty_sign:
            winner.append(lines[0][2])

    if len(winner) > 1:
        return impossible_text
    elif len(winner) == 1:
        return winner[0]
    else:
        return ""


def check_input(lines, x, y):
    try:
        x = int(x) - 1
        y = int(y) - 1
    except ValueError:
        return "You should enter numbers!"

    if not 0 <= x < 3 or not 0 <= y < 3:
        return "Coordinates should be from 1 to 3!"

    if lines[x][y] != empty_sign:
        return "This cell is occupied! Choose another one!"

    return ""


def set_move(x, y, sign):
    field_lines[x][y] = sign


print_field(field_lines)

user = "O"
turn = 0
max_turns = 9
while True:
    result = find_winner(field_lines)

    if result == "X" or result == "O":
        print(f"{result} wins")
        break
    elif result == impossible_text:
        print(result)
    elif result == "" and turn < max_turns:
        while True:
            print_field(field_lines)
            coord_x, coord_y = input().split()

            check_input_result = check_input(field_lines, coord_x, coord_y)
            if check_input_result != "":
                print(check_input_result)
                continue
            else:
                set_move(int(coord_x) - 1, int(coord_y) - 1, user)
                turn += 1
                break

        print_field(field_lines)
    else:
        print("Draw")
        break

    user = f"{'O' if user == 'X' else 'X'}"
