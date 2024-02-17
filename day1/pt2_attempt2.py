f = open("input.txt", "r")
input = f.readlines()
f.close()
OUTPUT_FROM_LINE = ""
SUM = 0

# converts alpha numbers to their numerical counterparts
def strings_to_numbers(argument: str):
    switcher = { "one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six":
                    6, "seven": 7, "eight": 8, "nine": 9, }

    return switcher.get(argument, "nothing")

# checks if the digit from the input needs to be converted to a numerical
def check_digit_needs_converting(digit):
    if (digit.isalpha()):
        digit = strings_to_numbers(digit)
    return digit

possible_digits = [ "1", "2", "3", "4", "5", "6", "7", "8", "9", "one", "two",
                   "three", "four", "five", "six", "seven", "eight", "nine" ]

for line in input:

    # strings to store the characters as the loop counts in from either side
    beginning_str = ""
    ending_str = ""
    # flags to check if the first and last digits have been found
    beginning_flag = True
    ending_flag = True

    for char in line:
        beginning_str = beginning_str + char
        if (beginning_flag):
            for digit in possible_digits:
                if (beginning_str.find(digit) != -1):
                    digit = check_digit_needs_converting(digit)
                    OUTPUT_FROM_LINE = str(digit) + OUTPUT_FROM_LINE
                    beginning_flag = False

    for char in line[::-1]:
        ending_str = char + ending_str
        if (ending_flag):
            for digit in possible_digits:
                if (ending_str.find(digit) != -1):
                    digit = check_digit_needs_converting(digit)
                    OUTPUT_FROM_LINE = OUTPUT_FROM_LINE + str(digit)
                    ending_flag = False


    SUM += int(OUTPUT_FROM_LINE)
    OUTPUT_FROM_LINE = ""

print(SUM)
