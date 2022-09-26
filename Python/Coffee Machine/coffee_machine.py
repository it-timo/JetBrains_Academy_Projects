class CoffeeMachine:
    def __init__(self):
        self.water = 400
        self.milk = 540
        self.coffee_beans = 120
        self.cups = 9
        self.money = 550
        self.status = 'default'

    def handle_states(self, text):
        match self.status:
            case'default':
                self.default_menu(text)
            case 'buy':
                self.buy_menu(text)
                self.status = 'default'
            case 'fill_water':
                self.water += int(text)
                self.status = 'fill_milk'
            case 'fill_milk':
                self.milk += int(text)
                self.status = 'fill_coffee_beans'
            case 'fill_coffee_beans':
                self.coffee_beans += int(text)
                self.status = 'fill_cups'
            case 'fill_cups':
                self.cups += int(text)
                print('')
                self.status = 'default'
            case _:
                print('Wrong option!')

    def default_menu(self, text):
        match text:
            case 'buy':
                self.status = 'buy'
            case 'fill':
                self.status = 'fill_water'
            case 'take':
                self.take_money()
                self.status = 'default'
            case 'remaining':
                self.show_remaining()
                self.status = 'default'
            case 'exit':
                self.status = 'exit'
            case _:
                print('Wrong option!')
        print('')

    def buy_menu(self, text):
        match text:
            case '1':
                self.make_coffee(250, 0, 16, 1, 4)  # espresso
            case '2':
                self.make_coffee(350, 75, 20, 1, 7)  # latte
            case '3':
                self.make_coffee(200, 100, 12, 1, 6)  # cappuccino
            case _:
                print('Wrong option!')
        print('')

    def make_coffee(self, water, milk, coffee_beans, cups, money):
        if self.water < water:
            print('Sorry, not enough water!')
        elif self.milk < milk:
            print('Sorry, not enough milk!')
        elif self.coffee_beans < coffee_beans:
            print('Sorry, not enough coffee beans!')
        elif self.cups < cups:
            print('Sorry, not enough disposable cups!')
        else:
            print('I have enough resources, making you a coffee!')
            self.water -= water
            self.milk -= milk
            self.coffee_beans -= coffee_beans
            self.cups -= cups
            self.money += money

    def fill_machine(self, water, milk, coffee_beans, cups):
        self.water += water
        self.milk += milk
        self.coffee_beans += coffee_beans
        self.cups += cups

    def take_money(self):
        print(f'I gave you ${self.money}')
        self.money = 0
        self.status = 'menu'

    def show_remaining(self):
        print('The coffee machine has:')
        print(f'{self.water} ml of water')
        print(f'{self.milk} ml of milk')
        print(f'{self.coffee_beans} g of coffee beans')
        print(f'{self.cups} disposable cups')
        print(f'${self.money} of money')


def main():
    coffee_machine = CoffeeMachine()

    while coffee_machine.status != 'exit':
        match coffee_machine.status:
            case 'default':
                coffee_machine.handle_states(input('Write action (buy, fill, take, remaining, exit):\n'))
            case 'buy':
                coffee_machine.handle_states(input(
                    'What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu: \n'))
            case 'fill_water':
                coffee_machine.handle_states(input('Write how many ml of water you want to add:\n'))
            case 'fill_milk':
                coffee_machine.handle_states(input('Write how many ml of milk you want to add:\n'))
            case 'fill_coffee_beans':
                coffee_machine.handle_states(input('Write how many grams of coffee beans you want to add:\n'))
            case 'fill_cups':
                coffee_machine.handle_states(input('Write how many disposable cups you want to add:\n'))


if __name__ == "__main__":
    main()
