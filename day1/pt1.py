f = open("input.txt", "r")
input = f.readlines()
f.close()
output_from_line = ""
sum = 0

for line in input:
    for char in line:
        if (char.isnumeric()):
            output_from_line = output_from_line + char
    output_from_line = output_from_line[0:1] + output_from_line[len(output_from_line) - 1 : len(output_from_line)]
    sum += int(output_from_line)
    output_from_line = ""

print(sum)
