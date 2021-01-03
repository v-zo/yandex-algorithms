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
	result string
	knobs  []string
}

func (c *Comb) combine(currentKnobIndex int) {
	if currentKnobIndex == 0 {
		c.result = strings.Join(
			strings.Split(
				c.knobs[currentKnobIndex],
				""),
			" ")
	} else {
		spr := strings.Split(c.result, " ")
		var newRes []string
		for _, s := range spr {
			currentKnob := c.knobs[currentKnobIndex]
			for i := 0; i < len(currentKnob); i++ {
				newRes = append(newRes, s+string(currentKnob[i]))
			}
		}
		c.result = strings.Join(newRes, " ")
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

	c := &Comb{"", knobs}
	c.combine(0)

	writer.WriteString(c.result)
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
