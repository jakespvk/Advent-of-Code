f = open("input.txt", "r")
input = f.readlines()
f.close()

available_cubes = { "red" : 12, "green" : 13, "blue" : 14 }
input_cubes = {}
accepted_games = []

for line in input:
    # find games where there are less than the available cubes
    input_words = line.split(' ')
    for i in range(2, len(input_words)):
        print(available_cubes[input_words[i + 1].strip(',')])
        if available_cubes[input_words[i + 1].strip(',')] <= int(input_words[i]):
            break
        accepted_games.append(input_words[1])

sum = 0
for each in accepted_games:
    sum += int(each)

print(sum)
