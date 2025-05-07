package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

func calculate(value1 string, value2 string) float64 {
	f1, err := strconv.ParseFloat(strings.TrimSpace(value1), 64)
	if err != nil {
		fmt.Println("Error converting value1")
		fmt.Println(err)
	}

	f2, err := strconv.ParseFloat(strings.TrimSpace(value2), 64)
	if err != nil {
		fmt.Println("Error converting value2")
		fmt.Println(err)
	}

	result := f1 + f2
	return math.Round(result*100) / 100
}

func main() {
	result := calculate("12.54", "3.5")
	fmt.Println(result)
}
