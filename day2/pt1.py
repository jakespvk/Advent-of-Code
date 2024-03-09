f = open("input.txt", "r")
input = f.readlines()
f.close()

available_cubes = { "red" : 12, "green" : 13, "blue" : 14 }

for line in input:
    # find games where there are less than the available cubes
