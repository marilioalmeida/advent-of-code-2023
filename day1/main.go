package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func getLines(filePath string) ([]string, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)
	// Create a slice to hold the lines
	var lines []string
	// Read each line
	for scanner.Scan() {
		lines = append(lines, scanner.Text()) // Append the line to the slice
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func problem1() {
	lines, _ := getLines("input.txt")
	fmt.Println(lines)

	sum := 0

	for _, line := range lines {
		var start rune
		var end rune
		for _, char := range line {
			if unicode.IsDigit(char) {
				if start == 0 {
					start = char
				}
				end = char
			}
		}

		combineString := string(start) + string(end)
		combineInt, _ := strconv.Atoi(combineString)
		sum += combineInt
	}

	fmt.Println(sum)
}

func problem2() {
	lines, _ := getLines("input2.txt")
	fmt.Println(lines)

	letters := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	sum := 0

	for _, line := range lines {
		var start string
		var end string
		for i, char := range line {
			if unicode.IsDigit(char) {
				if len(start) == 0 {
					start = string(char)
				}
				end = string(char)
			} else {
				word := string(char)

				for j := i + 1; j < len(line); j++ {
					if unicode.IsDigit(rune(line[j])) {
						break
					}
					word += string(line[j])

					if v, ok := letters[word]; ok {
						if len(start) == 0 {
							start = v
						}
						end = v
						break
					}
				}
			}
		}
		combineString := string(start) + string(end)
		combineInt, _ := strconv.Atoi(combineString)
		sum += combineInt
	}

	fmt.Println(sum)
}

func main() {
	problem2()
}
