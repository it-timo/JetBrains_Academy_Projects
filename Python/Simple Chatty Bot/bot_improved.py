from typing import NamedTuple


def greet(bot_name, birth_year):
    print(f"Hello! My name is {bot_name}.")
    print(f"I was created in {birth_year}.")


def remind_name():
    print("Please, remind me your name.")
    name = input()
    print(f"What a great name you have, {name}!")


def guess_age():
    print("Let me guess your age.")
    print("Enter remainders of dividing your age by 3, 5 and 7.")

    rem_3 = int(input())
    rem_5 = int(input())
    rem_7 = int(input())
    age = (rem_3 * 70 + rem_5 * 21 + rem_7 * 15) % 105

    print(f"Your age is {age}; that's a good time to start programming!")


def count():
    print("Now I will prove to you that I can count to any number you want.")

    num = int(input())
    curr = 0
    while curr <= num:
        print(curr, "!")
        curr = curr + 1


class QuestionStruct(NamedTuple):
    question: str
    answers: list
    index_correct_answer: int


def print_test_question(question):
    print(question.question)
    index = 1
    for answer in question.answers:
        print(f"{index}. {answer}")
        index += 1


def test(question):
    print("Let's test your programming knowledge.")
    # write your code here
    print_test_question(question)

    while True:
        user_input_answer = int(input())
        if user_input_answer == question.index_correct_answer:
            break
        else:
            print("Please, try again.")


def end():
    print('Congratulations, have a nice day!')


greet('Chatty', '2022')  # change it as you need
remind_name()
guess_age()
count()

# just for this purpose
question_tmp = "Why do we use methods?"
answer_1 = "To repeat a statement multiple times."
answer_2 = "To decompose a program into several small subroutines."
answer_3 = "To determine the execution time of a program."
answer_4 = "To interrupt the execution of a program."
answers_tmp = [answer_1, answer_2, answer_3, answer_4]
index_correct_answer = 1

q = QuestionStruct(question_tmp, answers_tmp, index_correct_answer)
test(q)

end()
