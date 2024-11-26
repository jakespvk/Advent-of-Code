package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
    fileName := "./short-input.txt"
    file, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    var lines []string
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    for i := 1; i < len(lines) - 1; i++ {
        curr_line := lines[i]
        prev_line := lines[i - 1]
        next_line := lines[i + 1]

        for j, rune_char := range(curr_line) {
            char := string(rune_char)

            if is_digit(char) {
                fmt.Println("here", j)
                if has_adjacent_symbol(j, i, curr_line, prev_line, next_line) {
                    fmt.Println("here2")
                }
            }
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func is_digit(char string) bool {
    switch char {
        case "1", "2", "3", "4", "5", "6", "7", "8", "9" : return true
    }
    return false
}

func is_symbol(char string) bool {
    switch char {
        case "!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "+", "=",
            "-", "_", "/", ":", ";", "'", "<", ">", "?", ",", "`", "|",
            "[", "]", "{", "}", "~": return true
    }
    return false
}

func has_adjacent_symbol(
    char_idx int, 
    line_idx int, 
    curr_line string,
    prev_line string, 
    next_line string,
) bool {
    if is_symbol(string(curr_line[char_idx - 1])) {
        fmt.Println("here3")
        return true
    }
    return false
}

