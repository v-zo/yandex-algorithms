package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file := openFile("input.txt")
	knobs, N := readData(file)
	file.Close()
	score := getScore(knobs, N)

	fmt.Println(score)
}

func getScore(knobs []string, k int) int {
	knobCountsMap := make(map[string]int)
	notNumericKnob := "."

	for _, knob := range knobs {
		if knob != notNumericKnob {
			knobCountsMap[knob]++
		}
	}

	DoubleK := 2 * k
	score := 0

	for _, knobCount := range knobCountsMap {
		hasEnoughFingers := knobCount <= DoubleK

		if hasEnoughFingers {
			score++
		}
	}

	return score
}

func readData(inputFile *os.File) ([]string, int) {
	dimension := 4
	scanner := bufio.NewScanner(inputFile)
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())
	var knobs []string

	for i := 0; i < dimension; i++ {
		scanner.Scan()
		line := scanner.Text()
		lineKnobs := strings.Split(line, "")
		knobs = append(knobs, lineKnobs...)
	}

	return knobs, k
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
