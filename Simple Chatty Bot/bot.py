def greet(bot_name, birth_year):
    print('Hello! My name is ' + bot_name + '.')
    print('I was created in ' + birth_year + '.')


def remind_name():
    print('Please, remind me your name.')
    name = input()
    print('What a great name you have, ' + name + '!')


def guess_age():
    print('Let me guess your age.')
    print('Enter remainders of dividing your age by 3, 5 and 7.')

    rem_3 = int(input())
    rem_5 = int(input())
    rem_7 = int(input())
    age = (rem_3 * 70 + rem_5 * 21 + rem_7 * 15) % 105

    print("Your age is " + str(age) + "; that's a good time to start programming!")


def count():
    print('Now I will prove to you that I can count to any number you want.')

    num = int(input())
    curr = 0
    while curr <= num:
        print(curr, '!')
        curr = curr + 1


def test():
    print("Let's test your programming knowledge.")
    # write your code here
    # just for this purpose
    question_tmp = "Why do we use methods?"
    answer_1 = "To repeat a statement multiple times."
    answer_2 = "To decompose a program into several small subroutines."
    answer_3 = "To determine the execution time of a program."
    answer_4 = "To interrupt the execution of a program."
    answers_tmp = [answer_1, answer_2, answer_3, answer_4]
    index_correct_answer = 1

    def print_test_question(question, answers):
        print(question)
        index = 1
        for answer in answers:
            print(f"{index}. {answer}")
            index += 1

    print_test_question(question_tmp, answers_tmp)

    while True:
        user_input_answer = int(input())
        if user_input_answer == index_correct_answer:
            break
        else:
            print("Please, try again.")


def end():
    print('Congratulations, have a nice day!')


greet('Chatty', '2022')  # change it as you need
remind_name()
guess_age()
count()
test()
end()
