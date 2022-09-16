from math import floor, ceil, log
import argparse


def check_args(a):
    if a.type != "annuity" and a.type != "diff":
        print("Incorrect parameters")
        return False
    if a.type == "diff" and a.payment != "":
        print("Incorrect parameters")
        return False
    if a.interest is None:
        print("Incorrect parameters")
        return False
    if len(vars(a)) < 5:
        print("Incorrect parameters")
        return False
    if a.interest is not None and a.interest < 0:
        print("Incorrect parameters")
        return False
    if a.payment is not None and a.payment < 0:
        print("Incorrect parameters")
        return False
    if a.principal is not None and a.principal < 0:
        print("Incorrect parameters")
        return False
    if a.periods is not None and a.periods < 0:
        print("Incorrect parameters")
        return False

    return True


def diff(principal, periods, interest):
    overpayment = 0
    i = interest / (12 * 100)

    for m in range(periods + 1):
        if m == 0:
            continue

        result = ceil(principal / periods + i * (principal - (principal - (principal * m - 1) / periods)))
        overpayment += result
        print(f"Month {m}: payment is {result}")

    print(f"Overpayment = {overpayment - principal}")


def monthly_annuity(principal, periods, interest):
    i = interest / (12 * 100)
    a = principal * (i * (1 + i) ** periods) / ((1 + i) ** periods - 1)
    print(f"Your monthly payment = {ceil(a)}!")


def loan_principal(annuity, periods, interest):
    i = interest / (12 * 100)
    p = annuity / ((i * (1 + i) ** periods) / ((1 + i) ** periods - 1))
    print(f"Your loan principal = {p}!")


def time(principal, payment, interest):
    months = 12

    i = interest / (months * 100)
    n = ceil(log((payment / (payment - i * principal)), 1 + i))

    m = n % months
    y = floor(n / months)

    if y == 0 and m > 0:
        o = f"It will take {m} month{'s' if m == 1 else ''} to repay this loan!"
    elif m == 0 and y > 0:
        o = f"It will take {y} year{'s' if y == 1 else ''} to repay this loan!"
    else:
        o = f"It will take {y} year{'s' if y == 1 else ''} and {m} month{'s' if m == 1 else ''} to repay this loan!"

    print(o)
    print(f"Overpayment = {principal - n * payment}")


give_opts = """What do you want to calculate?
type "n" for number of monthly payments,
type "a" for annuity monthly payment amount,
type "p" for loan principal:"""

parser = argparse.ArgumentParser()
# filename = args[0]
parser.add_argument("--type", choices=["annuity", "diff"], required=True,
                    help="You need to choose only one from the list.")
parser.add_argument("--payment", type=int, help="Can't be used with type=diff")
parser.add_argument("--interest", type=float, help="Must be set")
parser.add_argument("--principal", type=int)
parser.add_argument("--periods", type=int)

args = parser.parse_args()

check_args(args)

if args.interest is not None:
    if args.type == "annuity":
        if args.periods is None:
            time(args.principal, args.payment, args.interest)
        elif args.payment is None:
            monthly_annuity(args.principal, args.periods, args.interest)
        elif args.principal is None:
            loan_principal(args.payment, args.periods, args.interest)
    elif args.type == "diff":
        diff(args.principal, args.periods, args.interest)
