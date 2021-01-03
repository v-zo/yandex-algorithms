package main

import (
	"bufio"
	"os"
	"strings"
)

var knobsMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func main() {
	file := openFile("input.txt")
	defer file.Close()

	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(os.Stdout)

	Solve(reader, writer)
}

type Comb struct {
	result []string
	knobs  []string
}

func (c *Comb) combine(currentKnobIndex int) {
	currentKnob := c.knobs[currentKnobIndex]
	currentKnobArr := strings.Split(currentKnob, "")

	if currentKnobIndex == 0 {
		c.result = currentKnobArr
	} else {
		var newRes []string
		for _, s := range c.result {
			for i := 0; i < len(currentKnob); i++ {
				newRes = append(newRes, s+currentKnobArr[i])
			}
		}
		c.result = newRes
	}

	if currentKnobIndex == len(c.knobs)-1 {
		return
	}

	c.combine(currentKnobIndex + 1)
}

func Solve(reader *bufio.Reader, writer *bufio.Writer) {
	ns := readData(reader)

	var knobs []string
	for i := 0; i < len(ns); i++ {
		knobsValue := knobsMap[string(ns[i])]
		knobs = append(knobs, knobsValue)
	}

	c := &Comb{nil, knobs}
	c.combine(0)

	writer.WriteString(strings.Join(c.result, " "))
	writer.WriteString("\n")

	writer.Flush()
}

func readData(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')

	return line
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
