class Dominoes:
    import random

    def __init__(self):
        self.stock, self.player, self.computer, self.snake = [], [], [], []
        self.status = ""

    def start_game(self):
        while True:
            self.set_table()
            self.draw_dominoes()
            for x in [self.player, self. computer]:
                self.find_double(x)
            if self.status != "":
                break
            else:
                for x in [self.stock, self.player, self.computer, self.snake]:
                    x.clear()
        while True:
            self.print_table()

    def set_table(self):
        for i in range(0, 7):
            for j in range(i, 7):
                self.stock.append([i, j])

    def draw_dominoes(self):
        for i in range(0, 7):
            self.random.shuffle(self.stock)
            self.player.append(self.stock.pop())
            self.computer.append(self.stock.pop())

    def find_double(self, contestant):
        for i in range(6, 0, -1):
            if [i, i] in contestant:
                if len(self.snake) != 0:
                    if len(self.player) < len(self.computer):
                        self.player.append(self.snake.pop())
                    else:
                        self.computer.append(self.snake.pop())
                index = contestant.index([i, i])
                self.snake.append(contestant.pop(index))
                self.status = "computer" if contestant == self.player else "player"
                break

    def choose_domino(self, choice):
        while True:
            if abs(choice) in range(1, len(self.player) + 1):
                return self.add_domino(choice, self.player)
            elif choice == 0:
                self.player.append(self.stock.pop())
                return True
            else:
                print("Invalid input. Please try again.")
                return False

    def add_domino(self, choice, contestant):
        while True:
            if choice > 0:
                if self.snake[-1][1] in contestant[choice - 1]:
                    if self.snake[-1][1] == contestant[choice - 1][1]:
                        contestant[choice - 1].reverse()
                    self.snake.append(contestant.pop(choice - 1))
                    return True
                else:
                    if contestant == self.player:
                        print("Illegal move. Please try again.")
                    return False
            else:
                if self.snake[0][0] in contestant[abs(choice) - 1]:
                    if self.snake[0][0] == contestant[abs(choice) - 1][0]:
                        contestant[abs(choice) - 1].reverse()
                    self.snake.insert(0, contestant.pop(abs(choice) - 1))
                    return True
                else:
                    if contestant == self.player:
                        print("Illegal move. Please try again.")
                    return False

    def computer_move(self):
        occurrences = {0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0}
        scores = list()
        for i in range(7):
            for j in range(len(self.snake)):
                occurrences[i] += self.snake[j].count(i)
        for i in range(len(self.computer)):
            left = occurrences[self.computer[i][0]]
            right = occurrences[self.computer[i][1]]
            scores.append(left + right)
        while True:
            if scores.count(0) == len(self.computer):
                if len(self.stock) != 0:
                    self.computer.append(self.stock.pop())
                    break
                else:
                    break
            best_choice = scores.index(max(scores))
            if self.add_domino(best_choice + 1, self.computer):
                break
            elif self.add_domino(-best_choice + 1, self.computer):
                break
            else:
                if len(scores) != 0:
                    scores[best_choice] = 0

    def print_table(self):
        print(f'{70 * "="}\n'
              f'Stock size: {len(self.stock)}\n'
              f'Computer pieces: {len(self.computer)}\n\n')
        if len(self.snake) < 7:
            print(*self.snake[:])
        else:
            short = self.snake[:3] + ['...'] + self.snake[-3:]
            print(*short, sep="")
        print('Your pieces:')
        for x in range(0, len(self.player)):
            print(f'{x + 1}:{self.player[x]}')
        if self.end_check():
            exit()
        if self.status == "player":
            print("\nStatus: It's your turn to make a move. Enter your command.")
            while True:
                try:
                    if self.choose_domino(int(input())):
                        self.status = "computer"
                        break
                    else:
                        continue
                except ValueError:
                    print("Invalid input. Please try again.")
        else:
            print("\nStatus: Computer is about to make a move. Press Enter to continue...")
            self.computer_move()
            self.status = "player"
            input()

    def end_check(self):
        for x in [self.player, self.computer]:
            if len(x) == 0:
                winner = "You" if x == self.player else "The computer"
                print(f"Status: The game is over. {winner} won!")
                return True
        for x in range(0, 7):
            occurrence = 0
            if x in self.snake[0] and x in self.snake[-1]:
                for i in range(len(self.snake)):
                    for j in range(2):
                        if x == self.snake[i][j]:
                            occurrence += 1
            if occurrence == 8:
                print("Status: The game is over. It's a draw!")
                return True
        return False


new_game = Dominoes()
new_game.start_game()
