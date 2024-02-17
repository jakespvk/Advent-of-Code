f = open("input.txt", "r")
input = f.readlines()
output_from_line = ""
sum = 0

#def strings_to_numbers(argument):
#    switcher = { "one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six":
                 #    6, "seven": 7, "eight": 8, "nine": 9, }
#
#    return switcher.get(argument, "nothing")

for line in input:
    nums_in_line = {}
    for char in line:
        if (char.isnumeric()):
            nums_in_line[line.find(char)] = int(char)
    if (line.__contains__("one")):
        nums_in_line[line.index("one")] = 1
    elif (line.__contains__("two")):
        nums_in_line[line.index("two")] = 2
    elif (line.__contains__("three")):
        nums_in_line[line.index("three")] = 3
    elif (line.__contains__("four")):
        nums_in_line[line.index("four")] = 4
    elif (line.__contains__("five")):
        nums_in_line[line.index("five")] = 5
    elif (line.__contains__("six")):
        nums_in_line[line.index("six")] = 6
    elif (line.__contains__("seven")):
        nums_in_line[line.index("seven")] = 7
    elif (line.__contains__("eight")):
        nums_in_line[line.index("eight")] = 8
    elif (line.__contains__("nine")):
        nums_in_line[line.index("nine")] = 9
    
    min = 9999
    max = 0
    for key in nums_in_line:
        if key < min:
            min = key
        if key > max:
            max = key

    output_from_line = str(nums_in_line[min]) + str(nums_in_line[max])
    #output_from_line = output_from_line[0:1] + output_from_line[len(output_from_line) - 1 : len(output_from_line)]
    sum += int(output_from_line)
    output_from_line = ""

print(sum)
