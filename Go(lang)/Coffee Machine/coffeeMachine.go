package main

import "fmt"

type coffeeMachine struct {
	water int
	milk  int
	beans int
	sugar int
	sCups int
	mCups int
	lCups int
	money int
}

func main() {
	machine := &coffeeMachine{
		water: 400,
		milk:  540,
		sugar: 100,
		beans: 120,
		sCups: 9,
		mCups: 9,
		lCups: 9,
		money: 550,
	}

	startAction(machine)
}

func startAction(machine *coffeeMachine) {
	fmt.Println("Write action (buy, fill, take, remaining, exit):")

	for {
		var choice string
		_, err := fmt.Scanf("%s", &choice)
		if err != nil {
			fmt.Println(err)
		}

		switch choice {
		case "buy":
			chooseCoffee(machine)
		case "fill":
			fillMachine(machine)
		case "take":
			takeMoney(machine)
		case "remaining":
			printCurrentIngredients(machine)
		case "exit":
			return
		}
	}
}

func printCurrentIngredients(machine *coffeeMachine) {
	fmt.Println("The coffee machine has:")
	fmt.Printf("%d ml of water\n", machine.water)
	fmt.Printf("%d ml of milk\n", machine.milk)
	fmt.Printf("%d g of coffee beans\n", machine.beans)
	fmt.Printf("%d g of sugar\n", machine.sugar)
	fmt.Printf("%d small disposable cups\n", machine.sCups)
	fmt.Printf("%d middle disposable cups\n", machine.mCups)
	fmt.Printf("%d large disposable cups\n", machine.lCups)
	fmt.Printf("$%d of money\n", machine.money)
}

func makeCoffee(water, milk, beans, sugar, sCups, mCups, lCups, money int, machine *coffeeMachine) {
	isMissing, ingredient := missingIngredient(water, milk, beans, sugar, sCups, mCups, lCups, machine)
	if isMissing {
		fmt.Printf("Sorry, not enough %s!\n", ingredient)
	} else {
		fmt.Println("I have enough resources, making you a coffee!")

		machine.water -= water
		machine.milk -= milk
		machine.beans -= beans
		machine.sugar -= sugar
		machine.sCups -= sCups
		machine.mCups -= mCups
		machine.lCups -= lCups
		machine.money += money
	}
}

func missingIngredient(water, milk, beans, sugar, sCups, mCups, lCups int, machine *coffeeMachine) (bool, string) {
	if machine.water < water {
		return true, "water"
	}
	if machine.milk < milk {
		return true, "milk"
	}
	if machine.beans < beans {
		return true, "beans"
	}
	if machine.sugar < sugar {
		return true, "sugar"
	}
	if machine.sCups < sCups {
		return true, "small cups"
	}
	if machine.mCups < mCups {
		return true, "medium cups"
	}
	if machine.lCups < lCups {
		return true, "large cups"
	}

	return false, ""
}

func fillMachine(machine *coffeeMachine) {
	var tmpVal int

	fmt.Println("Write how many ml of water you want to add:")
	machine.water += getValueForFill(tmpVal)
	fmt.Println("Write how many ml of milk you want to add:")
	machine.milk += getValueForFill(tmpVal)
	fmt.Println("Write how many grams of coffee beans you want to add:")
	machine.beans += getValueForFill(tmpVal)
	fmt.Println("Write how many grams of sugar you want to add:")
	machine.sugar += getValueForFill(tmpVal)
	fmt.Println("Write how many small disposable cups you want to add:")
	machine.sCups += getValueForFill(tmpVal)
	fmt.Println("Write how many medium disposable cups you want to add:")
	machine.mCups += getValueForFill(tmpVal)
	fmt.Println("Write how many large disposable cups you want to add:")
	machine.lCups += getValueForFill(tmpVal)
}

func getValueForFill(fillType int) int {
	_, err := fmt.Scanf("%d", &fillType)
	if err != nil {
		fmt.Println(err)
	}

	return fillType
}

func takeMoney(machine *coffeeMachine) {
	fmt.Printf("I gave you $%d\n", machine.money)
	machine.money = 0
}

func chooseSize() string {
	fmt.Println("Choose your size? 1 - small, 2 - medium, 3 - large, back - to main menu:")

	var size string
	_, err := fmt.Scanf("%s", &size)
	if err != nil {
		fmt.Println(err)
	}

	return size
}

func chooseSugar() string {
	fmt.Println("Do you want sugar? yes, no, back - to main menu:")

	var sugarChoice string
	_, err := fmt.Scanf("%s", &sugarChoice)
	if err != nil {
		fmt.Println(err)
	}

	return sugarChoice
}

func chooseCoffee(machine *coffeeMachine) {
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:")

	var coffeeTyp string
	_, err := fmt.Scanf("%s", &coffeeTyp)
	if err != nil {
		fmt.Println(err)
	}

	var sugar int
	switch chooseSugar() {
	case "yes":
		sugar = 0
	case "no":
		sugar = 3
	case "back":
		return
	}

	var sc, mc, lc int
	switch chooseSize() {
	case "1":
		sc, mc, lc = 1, 0, 0
	case "2":
		sc, mc, lc = 0, 1, 0
	case "3":
		sc, mc, lc = 0, 0, 1
	case "back":
		return
	}

	switch coffeeTyp {
	case "1":
		makeCoffee(250, 0, 16, sugar, sc, mc, lc, 4, machine)
	case "2":
		makeCoffee(350, 75, 20, sugar, sc, mc, lc, 7, machine)
	case "3":
		makeCoffee(200, 100, 12, sugar, sc, mc, lc, 6, machine)
	case "back":
		return
	}
}
