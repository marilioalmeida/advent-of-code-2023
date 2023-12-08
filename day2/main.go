package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	Name   string
	Rounds []Round
}

type Round struct {
	Colors map[string]int
}

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

func parseGame(str string) (Game, error) {
	// Split the string into game name and rounds
	parts := strings.SplitN(str, ":", 2)
	if len(parts) != 2 {
		return Game{}, fmt.Errorf("invalid format")
	}

	// Parse game name
	gameName := strings.TrimSpace(parts[0])

	// Split rounds string into individual rounds
	roundsStr := strings.SplitN(parts[1], ";", -1)

	// Initialize game and rounds slice
	game := Game{
		Name:   gameName,
		Rounds: make([]Round, 0, len(roundsStr)),
	}

	// Parse each round
	for _, roundStr := range roundsStr {
		round, err := parseRound(roundStr)
		if err != nil {
			return Game{}, fmt.Errorf("error parsing round: %w", err)
		}

		game.Rounds = append(game.Rounds, round)
	}

	return game, nil
}

func parseRound(str string) (Round, error) {
	round := Round{
		Colors: make(map[string]int),
	}

	// Split the round string into color pairs
	pairs := strings.Split(str, ",")

	for _, pair := range pairs {
		// Trim spaces and split into color and count
		parts := strings.SplitN(strings.TrimSpace(pair), " ", 2)
		if len(parts) != 2 {
			return Round{}, fmt.Errorf("invalid color pair format")
		}

		// Convert count to integer
		count, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			return Round{}, fmt.Errorf("error parsing color count: %w", err)
		}

		round.Colors[strings.TrimSpace(parts[1])] = count
	}

	return round, nil
}

func problem1() {
	lines, _ := getLines("input1.txt")

	var Games []Game

	for _, line := range lines {
		game, err := parseGame(line)
		Games = append(Games, game)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	availableCubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	possibleGames := []int{}
	for i, game := range Games {
		isPossible := true
		for _, round := range game.Rounds {
			for color, count := range round.Colors {
				if count > availableCubes[color] {
					isPossible = false
					break
				}
			}
			if !isPossible {
				break
			}
		}

		if isPossible {
			possibleGames = append(possibleGames, i+1)
		}
	}

	sumOfIDs := 0
	for _, id := range possibleGames {
		sumOfIDs += id
	}

	fmt.Println(sumOfIDs)

}

func problem2() {
	lines, _ := getLines("input2.txt")
	fmt.Println(lines)

}

func main() {
	problem1()
}
