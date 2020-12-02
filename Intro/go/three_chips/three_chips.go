package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	s := readData()
	N, _ := strconv.Atoi(s[0])
	stringArray := strings.Fields(s[1])

	values := strArrToIntArr(stringArray)

	result := findSumMin(values, N)

	fmt.Print(result)
}

func findSumMin(origValues []int, N int) int {
	values := append([]int(nil), origValues...)
	sort.Ints(values)
	length := len(values)
	var Snow int
	So := values[0] + values[1] + values[length-1]

	for i := 0; i < length; i++ {
		k := i + 1
		j := length - 1
		for k < j {
			Snow = values[i] + values[j] + values[k]

			if Snow < N {
				k++
			} else if Snow > N {
				j--
			} else {
				return Snow
			}
		}

		if intAbs(Snow-N) < intAbs(So-N) {
			So = Snow
		}
	}

	return So
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readData() []string {
	data, err := ioutil.ReadFile("input.txt")
	check(err)
	dataString := string(data)

	return strings.Split(dataString, "\n")
}

func strArrToIntArr(stringArray []string) []int {
	intArray := make([]int, 0, len(stringArray))

	for _, strElement := range stringArray {
		numericEl, _ := strconv.Atoi(strElement)
		intArray = append(intArray, numericEl)
	}

	return intArray
}

func intAbs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
