package old

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file := openFile("input.txt")
	sites, n := readData(file)
	file.Close()

	zeroesPositions := GetZeroesPositions(sites)

	result := SolveProblem(zeroesPositions, n)

	fmt.Println(intArrToStr(result))
}

func SolveProblem(zeroesPositions []int, n int) []int {
	totalZeroPositions := len(zeroesPositions)

	var headSegment []int
	if zeroesPositions[0] != 0 {
		headSegment = MakeRangeReversed(zeroesPositions[0], 1)
	}

	getSegment := func(i int) []int {
		diff := zeroesPositions[i] - 1 - zeroesPositions[i-1]
		radius := diff / 2

		leftPart := MakeRange(0, radius)
		rightPart := MakeRangeReversed(radius+diff%2, 1)

		return append(leftPart, rightPart...)
	}

	var bodySegments []int
	for i := 1; i < totalZeroPositions; i++ {
		segment := getSegment(i)
		bodySegments = append(bodySegments, segment...)
	}

	var tailSegment []int
	lastZeroPosition := zeroesPositions[totalZeroPositions-1]
	tailSegment = MakeRange(0, n-lastZeroPosition-1)

	result := append(headSegment, bodySegments...)

	return append(result, tailSegment...)
}

func MakeRange(start, stop int) []int {
	dim := stop - start
	a := make([]int, dim+1)

	for i := range a {
		a[i] = start + i
	}

	return a
}

func MakeRangeReversed(stop, start int) []int {
	dim := stop - start
	a := make([]int, dim+1)

	for i := range a {
		a[dim-i] = start + i
	}

	return a
}

func GetZeroesPositions(sites []string) []int {
	var zeroesPositions []int
	for position, homeNumber := range sites {
		if homeNumber == "0" {
			zeroesPositions = append(zeroesPositions, position)
		}
	}

	return zeroesPositions
}

func readData(inputFile *os.File) ([]string, int) {
	scanner := bufio.NewScanner(inputFile)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	line := scanner.Text()
	sites := strings.Split(line, " ")

	return sites, n
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}

func intArrToStr(result []int) string {
	return strings.Trim(
		strings.Replace(
			fmt.Sprint(result),
			" ", " ", -1),
		"[]")
}
