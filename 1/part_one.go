package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

// Return the two digit number formed by the first and last digit found in the string.
func findDigits(line string) int {
	digits := make([]string, 0)

	for _, c := range line {
		if unicode.IsDigit(c) {
			digits = append(digits, string(c))
		}
	}

	i, err := strconv.Atoi(digits[0] + digits[len(digits)-1])
	if err != nil {
		panic(err)
	}

	return i
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sum := 0
	lines := bufio.NewScanner(f)
	for lines.Scan() {
		sum += findDigits(lines.Text())
	}

	if err = lines.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum)
}
