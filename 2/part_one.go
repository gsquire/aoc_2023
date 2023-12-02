package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isValidGame(game string) bool {
	const (
		TOTAL_RED   = 12
		TOTAL_GREEN = 13
		TOTAL_BLUE  = 14
	)

	for _, pull := range strings.Split(game, ",") {
		pull = strings.TrimSpace(pull)
		metadata := strings.Split(pull, " ")
		count, _ := strconv.Atoi(metadata[0])
		color := metadata[1]
		if color == "red" && count > TOTAL_RED {
			return false
		} else if color == "green" && count > TOTAL_GREEN {
			return false
		} else if color == "blue" && count > TOTAL_BLUE {
			return false
		}
	}

	return true
}

// Return true if the game was possible.
func parseGame(line string) bool {
	gameData := strings.Split(line, ": ")[1]
	games := strings.Split(gameData, ";")

	for _, game := range games {
		if !isValidGame(game) {
			return false
		}
	}

	return true
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sum := 0
	gameID := 1
	lines := bufio.NewScanner(f)
	for lines.Scan() {
		if parseGame(lines.Text()) {
			sum += gameID
		}
		gameID++
	}

	if err = lines.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum)
}
