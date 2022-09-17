package main

import (
	"fmt"
)

type question struct {
	question string
	answer   int
	choices  []string
}

const (
	Name      = "Chatty"
	BirthYear = 2022
)

func main() {
	printGreetings(Name, BirthYear)
	getPlayerName()
	guessPlayerAge()
	count()

	q := question{
		question: "Why do we use methods?",
		answer:   2,
		choices: []string{"1. To repeat a statement multiple times.",
			"2. To decompose a program into several small subroutines.",
			"3. To determine the execution time of a program.",
			"4. To interrupt the execution of a program."},
	}
	startQuiz(q)
	printGoodbye()
}

// prints a greeting message into stdout
func printGreetings(name string, birthYear int) {
	fmt.Printf("Hello! My name is %s.\n", name)
	fmt.Printf("I was created in %d.\n", birthYear)
}

// prints a goodbye message into stdout
func printGoodbye() {
	fmt.Println("Congratulations, have a nice day!")
}

// reads the player name from stdin and print it with a message
func getPlayerName() {
	var playerName string
	fmt.Printf("Please, remind me your name.\n")

	_, err := fmt.Scan(&playerName)
	if err != nil {
		return
	}

	fmt.Printf("What a great name you have, %s!\n", playerName)
}

// reads the 3 int values from stdin and prints the calculated age into stdout with a message
func guessPlayerAge() {
	fmt.Printf("Let me guess your age.\n")
	fmt.Printf("Enter remainders of dividing your age by 3, 5 and 7\n")

	var remainder3, remainder5, remainder7 int
	_, err := fmt.Scan(&remainder3, &remainder5, &remainder7)
	if err != nil {
		return
	}

	age := (remainder3*70 + remainder5*21 + remainder7*15) % 105

	fmt.Printf("Your age is %d; that's a good time to start programming!\n", age)
}

// reads an int value from stdin and prints all values from 0 up to the value into stdout
func count() {
	fmt.Printf("Now I will prove to you that I can count to any number you want.\n")

	var requestNumber int
	_, err := fmt.Scan(&requestNumber)
	if err != nil {
		return
	}

	for i := 0; i <= requestNumber; i++ {
		fmt.Printf("%d !\n", i)
	}
}

// takes a question struct and prints it into stdout (looping a message until correct answer is read from stdin)
func startQuiz(q question) {
	fmt.Println("Let's test your programming knowledge.")
	fmt.Println(q.question)

	for _, choice := range q.choices {
		fmt.Println(choice)
	}

	for {
		var n int
		_, err := fmt.Scan(&n)
		if err != nil {
			return
		}

		if n == q.answer {
			break
		}

		fmt.Println("Please, try again.")
	}
}
