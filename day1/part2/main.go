package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

func main() {
    digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
                "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

    fileName := "./input.txt"
    file, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    var sum int

    for scanner.Scan() {
        line := scanner.Text()

        var value1 string
        var value2 string

        for i := 0; i < len(line); i++ {
            char := string(line[i])
            if slices.Contains(digits, char) {
                value1 = char
                break
            }
            if i + 2 < len(line) {
                slice := line[i:i+3]
                char := string(slice)
                if slices.Contains(digits, char) {
                    value1 = char
                    break
                }
            }
            if i + 3 < len(line) {
                slice := line[i:i+4]
                char := string(slice)
                if slices.Contains(digits, char) {
                    value1 = char
                    break
                }
            }
            if i + 4 < len(line) {
                slice := line[i:i+5]
                char := string(slice)
                if slices.Contains(digits, char) {
                    value1 = char
                    break
                }
            }
            if i + 5 < len(line) {
                slice := line[i:i+6]
                char := string(slice)
                if slices.Contains(digits, char) {
                    value1 = char
                    break
                }
            }
        }
        for i := len(line) - 1; i >= 0; i-- {
            char := string(line[i])
            if slices.Contains(digits, char) {
                value2 = char
                break
            }
            if i - 2 >= 0 {
                char := string(line[i-2:i+1])
                if slices.Contains(digits, char) {
                    value2 = char
                    break
                }
            }
            if i - 3 >= 0 {
                char := string(line[i-3:i+1])
                if slices.Contains(digits, char) {
                    value2 = char
                    break
                }
            }
            if i - 4 >= 0 {
                char := string(line[i-4:i+1])
                if slices.Contains(digits, char) {
                    value2 = char
                    break
                }
            }
            if i - 5 >= 0 {
                char := string(line[i-5:i+1])
                if slices.Contains(digits, char) {
                    value2 = char
                    break
                }
            }
        }

        val, err := strconv.Atoi(string(value1 + value2))
        if err != nil {
            switch value1 {
                case "one": value1 = "1"
                case "two": value1 = "2"
                case "three": value1 = "3"
                case "four": value1 = "4"
                case "five": value1 = "5"
                case "six": value1 = "6"
                case "seven": value1 = "7"
                case "eight": value1 = "8"
                case "nine": value1 = "9"
            }
            switch value2 {
                case "one": value2 = "1"
                case "two": value2 = "2"
                case "three": value2 = "3"
                case "four": value2 = "4"
                case "five": value2 = "5"
                case "six": value2 = "6"
                case "seven": value2 = "7"
                case "eight": value2 = "8"
                case "nine": value2 = "9"
            }
            val, err := strconv.Atoi(string(value1 + value2))
            if err != nil {
                log.Fatal(err)
            } else {
                sum += val
            }
        } else {
            sum += val
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Println(sum)

}

