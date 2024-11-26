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
    digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

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
        }
        for i := len(line) - 1; i >= 0; i-- {
            char := string(line[i])
            if slices.Contains(digits, char) {
                value2 = char
                break
            }
        }

        val, err := strconv.Atoi(string(value1 + value2))
        if err != nil {
            log.Fatal(err)
        } else {
            sum += val
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Println(sum)

}

