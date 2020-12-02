package main

import (
	"bufio"
	"os"
	"strconv"
)

type ProblemSolver struct {
	scanner     *bufio.Scanner
	writer      *bufio.Writer
	homeCounter int
}

func main() {
	file := openFile("input.txt")
	defer file.Close()

	problemSolver := createProblemSolver(file, os.Stdout)

	problemSolver.
		ProcessHead().
		ProcessBody().
		ProcessTail().
		Finish()
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}

func createProblemSolver(input *os.File, output *os.File) ProblemSolver {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)
	scanner.Scan() // skip useless n

	writer := bufio.NewWriter(output)

	return ProblemSolver{scanner, writer, 0}
}

func (ps ProblemSolver) WriteStringWithSpace(str string) {
	ps.writer.WriteString(str)
	ps.writer.WriteString(" ")
}

func (ps ProblemSolver) ProcessHead() ProblemSolver {
	writeHead := func(homeCounter int) {
		for i := homeCounter; i > 0; i-- {
			ps.WriteStringWithSpace(strconv.Itoa(i))
		}
	}

	var home string
	counter := -1
	for home != "0" {
		ps.scanner.Scan()
		home = ps.scanner.Text()
		counter++
	}

	if counter > 0 {
		writeHead(counter)
	}

	ps.WriteStringWithSpace(home)

	return ps
}

func (ps ProblemSolver) ProcessBody() ProblemSolver {
	writeBody := func(homeCounter int) {
		halfOfSegment := homeCounter / 2
		halfOfSegmentPlusOne := halfOfSegment + homeCounter%2
		for i := 1; i <= halfOfSegmentPlusOne; i++ {
			ps.WriteStringWithSpace(strconv.Itoa(i))
		}
		for i := halfOfSegment; i > 0; i-- {
			ps.WriteStringWithSpace(strconv.Itoa(i))
		}
	}

	for ps.scanner.Scan() {
		home := ps.scanner.Text()

		if home == "0" {
			if ps.homeCounter > 0 {
				writeBody(ps.homeCounter)
				ps.homeCounter = 0
			}

			ps.WriteStringWithSpace(home)
		} else {
			ps.homeCounter++
		}
	}

	return ps
}

func (ps ProblemSolver) ProcessTail() ProblemSolver {
	if ps.homeCounter == 0 {
		return ps
	}

	writeTail := func(homeCounter int) {
		for i := 1; i <= homeCounter; i++ {
			ps.writer.WriteString(strconv.Itoa(i))
			if i < homeCounter {
				ps.writer.WriteString(" ")
			}
		}
	}

	writeTail(ps.homeCounter)

	return ps
}

func (ps ProblemSolver) Finish() ProblemSolver {
	ps.writer.WriteString("\n")
	ps.writer.Flush()

	return ps
}
