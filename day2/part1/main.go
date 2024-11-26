package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
    fileName := "./input.txt"
    file, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    MAX_RED_CUBES := 12
    MAX_GREEN_CUBES := 13
    MAX_BLUE_CUBES := 14

    game_id := 0
    var successful_hands []int
    for scanner.Scan() {
        game_id += 1
        line := scanner.Text()

        splitLine := strings.Split(line, ":")
        matches := strings.Split(splitLine[1], ";")

        FAIL := false

        for _, match := range(matches) {
            match = strings.TrimSpace(match)
            hands := strings.Split(match, ",")

            for _, hand := range(hands) {
                hand = strings.TrimSpace(hand)

                var num_cubes string
                var color_cubes string
                var SPACE_FOUND = false
                for i := 0; i < len(hand); i++ {
                    char := string(hand[i])

                    if char == " " {
                        SPACE_FOUND = true
                    } else if !SPACE_FOUND {
                        num_cubes = num_cubes + char
                    } else if SPACE_FOUND {
                        color_cubes = color_cubes + char
                    }
                }

                int_num_cubes, err := strconv.Atoi(num_cubes)
                if err != nil {
                    log.Fatal(err)
                }

                var red_cubes int
                var green_cubes int
                var blue_cubes int
                switch color_cubes {
                    case "red" : red_cubes = int_num_cubes
                    case "green" : green_cubes = int_num_cubes
                    case "blue" : blue_cubes = int_num_cubes
                }

                if red_cubes > MAX_RED_CUBES {
                    FAIL = true
                    break
                } else if green_cubes > MAX_GREEN_CUBES {
                    FAIL = true
                    break
                } else if blue_cubes > MAX_BLUE_CUBES {
                    FAIL = true
                    break
                } 
            }
        }

        if !FAIL {
            successful_hands = append(successful_hands, game_id)
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    var sum_of_ids int
    for _, each := range successful_hands {
        sum_of_ids += each
    }
    fmt.Println(sum_of_ids)
}
