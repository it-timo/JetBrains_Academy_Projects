import requests

js = requests.get(f'http://www.floatrates.com/daily/{input().strip().lower()}.json').json()

rates = {"usd": js["usd"]["rate"], "eur": js["eur"]["rate"]}

while True:
    currency = input().strip().lower()
    if currency == "":
        break
    amount = float(input().strip())

    print("Checking the cache...")
    if currency in rates:
        print("Oh! It is in the cache!")
    else:
        print("Sorry, but it is not in the cache!")
        rates[currency] = js[currency]["rate"]

    print(f"You received {round(amount * rates[currency], 2)} {currency.upper()}.")
