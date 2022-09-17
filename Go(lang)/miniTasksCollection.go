package main

import "fmt"

func addingNumbers() {
	var a, b int

	_, err := fmt.Scan(&a, &b)
	if err != nil {
		return
	}

	fmt.Println(a + b)
}

func gradingTests() {
	var score int

	_, err := fmt.Scanf("%d", &score)
	if err != nil {
		return
	}

	if score >= 71 {
		fmt.Println("Passed!")
	} else {
		fmt.Println("Failed!")
	}
}

func movieRecommendations() {
	var age int

	_, err := fmt.Scanf("%d", &age)
	if err != nil {
		return
	}

	if age <= 14 {
		fmt.Println("Toy Story 4")
	} else if 15 <= age && age <= 18 {
		fmt.Println("The Matrix")
	} else if 19 <= age && age <= 25 {
		fmt.Println("John Wick")
	} else if 26 <= age && age <= 35 {
		fmt.Println("Constantine")
	} else if age > 35 {
		fmt.Println("Speed")
	}
}

func positiveNegativeNumbers() {
	var number int

	_, err := fmt.Scanf("%d", &number)
	if err != nil {
		return
	}

	if number > 0 {
		fmt.Println("Positive!")
	} else if number < 0 {
		fmt.Println("Negative!")
	} else {
		fmt.Println("Zero!")
	}
}

func printingNumbersInRange() {
	for i := 2; i <= 1023; i++ {
		fmt.Println(i)
	}
}

func reversingNumber() {
	var number int

	_, err := fmt.Scanf("%d", &number)
	if err != nil {
		return
	}

	newInt := 0
	for number > 0 {
		remainder := number % 10
		newInt *= 10
		newInt += remainder
		number /= 10
	}

	fmt.Println(newInt)
}

func findingFactorial() {
	var number int

	_, err := fmt.Scanf("%d", &number)
	if err != nil {
		return
	}
	// put your code here
	factVal := uint64(1)
	for i := 1; i <= number; i++ {
		factVal *= uint64(i)
	}

	fmt.Println(factVal)
}

func printingOddNumbersInRange() {
	for i := 2; i <= 1023; i++ {
		if i%2 == 1 {
			fmt.Println(i)
		}
	}
}
