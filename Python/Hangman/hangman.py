# Write your code here
import random

welcome_msg = "H A N G M A N\n"
menu_msg = '"Type "play" to play the game, "results" to show the scoreboard, and "exit" to quit:"'
get_char_msg = "Input a letter:"
single_letter_msg = "Please, input a single letter."
lower_letter_msg = "Please, enter a lowercase letter from the English alphabet."
wrong_char_msg = "That letter doesn't appear in the word."
already_revealed_msg = "You've already guessed this letter."
correct_guess_msg = "You guessed the word "
win_msg = "You survived!"
lose_msg = "You lost!"
empty_char = "-"
win_counter = 0
lost_counter = 0

char_list = ["a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
             "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"]
words_list = ['python', 'java', 'swift', 'javascript']


def check_input(char, revealed):
    if char == "":
        return single_letter_msg
    elif not char.islower():
        return lower_letter_msg
    elif char.isnumeric():
        return lower_letter_msg
    elif not char.isalpha():
        return lower_letter_msg
    elif len(char) != 1:
        return single_letter_msg
    elif char not in char_list:
        return lower_letter_msg
    elif char in revealed:
        return already_revealed_msg
    return ""


def results(w_count, l_count):
    print(f"You won: {w_count} times")
    print(f"You lost: {l_count} times")


def game():
    correct_word = random.choice(words_list)
    curr_revealed_word = list(empty_char * len(correct_word))
    user_chars = set()
    wrong_attempts_left = 8

    while True:
        print("".join(curr_revealed_word))
        user_letter = input(get_char_msg)
        char_check_result = check_input(user_letter, user_chars)

        if char_check_result != "":
            print(char_check_result)
            continue
        elif user_letter not in correct_word:
            print(wrong_char_msg)
            wrong_attempts_left -= 1
        else:
            for index, letter_cur in enumerate(correct_word):
                if letter_cur == user_letter:
                    curr_revealed_word[index] = user_letter

        user_chars.add(user_letter)

        if wrong_attempts_left <= 0:
            print(lose_msg)
            global lost_counter
            lost_counter += 1
            break
        elif "".join(curr_revealed_word) == correct_word:
            print(f"{correct_guess_msg}{correct_word}!")
            print(win_msg)
            global win_counter
            win_counter += 1
            break


def menu():
    while True:
        print(menu_msg)

        user_input = input()
        if user_input == "play":
            game()
        elif user_input == "results":
            results(win_counter, lost_counter)
        elif user_input == "exit":
            break


print(welcome_msg)
menu()
