## GOAL: take the first digit (even if its written as string. example: "one") and last digit number (even if its written as string. example: "one")
#        on a single line and form a two digit number only as an output then take its sum. 


file = open("input.txt", "r")
# read the lines and store it in a list
lines = file.readlines()

valid_num_strings = {"one":"1", "two":"2", "three":"3", "four":"4", "five":"5",
                     "six":"6", "seven":"7","eight":"8", "nine":"9"}


new_lines = []

# Replace the alphabet number with a string number
for line in lines:
    for key in valid_num_strings:
        if key in line:
            line = line.replace(key,valid_num_strings[key])
    new_lines.append(line)
    
    
# Function to get the first number 
def get_first_number(line: str):
    for i in range(len(line)):
        #check if the i-th element in line is a number
        if line[i].isnumeric():
            # return that number which is still a STRING ! 
            return line[i]


# Function to get the last number
def get_last_number(line: str):
    for i in reversed(range(len(line))):
        #check if the i-th element in line is a number
        if line[i].isnumeric():
            # return that number which is still a STRING ! 
            return line[i]

# Function that lets you add two digits of type str and return a single two-digits number as an int
def add_two_str_numbers(a: str,  b: str) -> int:
    return int(a+b)


# create a list of all numbers 
number_list = []

for line in new_lines:
    first_num = get_first_number(line)
    last_num = get_last_number(line)
    two_digits_num = add_two_str_numbers(first_num,last_num)
    number_list.append(two_digits_num)

total = sum(number_list)

print("The sum of your puzzle is:", total)
