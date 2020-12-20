package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	parenthesesPairs = map[string]string{
		"[": "]",
		"(": ")",
		"{": "}",
	}
)

func main() {
	file := openFile("input.txt")
	defer file.Close()

	reader := bufio.NewReader(file)

	fmt.Println(Solve(reader))
}

func Solve(r *bufio.Reader) string {
	scanner := bufio.NewScanner(r)
	scanner.Split(ScanParentheses)

	stack := &Stack{}

	for scanner.Scan() {
		ch := scanner.Text()

		if strings.Contains("[{(", ch) {
			stack.push(ch)
		} else {
			top, errTop := stack.top()

			if errTop != nil {
				return "False"
			}

			cl, _ := getClosing(top)
			if ch != cl {
				return "False"
			}

			err := stack.pop()

			if err != nil {
				return "False"
			}
		}
	}

	if len(stack.data) > 0 {
		return "False"
	}

	return "True"
}

func getClosing(ch string) (s string, err error) {
	s = parenthesesPairs[ch]

	if s == "" {
		err = errors.New(ch + " is bad symbol")
	}

	return s, err
}

func ScanParentheses(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if i := bytes.IndexAny(data, "[{()}]"); i >= 0 {
		iOne := i + 1
		return iOne, data[i:iOne], nil
	}

	return 0, nil, nil
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}

type Stack struct {
	data []string
}

func (stack *Stack) top() (string, error) {
	if len(stack.data) == 0 {
		return "", errors.New("there is no top of empty stack")
	}

	return stack.data[len(stack.data)-1], nil
}

func (stack *Stack) push(val string) {
	stack.data = append(stack.data, val)
}

func (stack *Stack) pop() error {
	if len(stack.data) == 0 {
		return errors.New("error")
	}

	stack.data = stack.data[:len(stack.data)-1]
	return nil
}
