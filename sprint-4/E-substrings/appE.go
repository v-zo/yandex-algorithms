package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//var fnvHash hash2.Hash32 = fnv.New32a()

func main() {
	file := openFile("input.txt")
	defer file.Close()

	reader := bufio.NewReader(file)

	Solve(reader)
}

func uniqueChars(s string) (uw []byte) {
	set := make(map[int32]struct{})
	for _, ch := range s {
		set[ch] = struct{}{}
	}

	for ch := range set {
		uw = append(uw, byte(ch))
	}

	return
}

//func hash(s []byte) uint32 {
//	fnvHash.Write(s)
//	return fnvHash.Sum32()
//}

func uniqueCharsSum(s string) (uw int32) {
	set := make(map[int32]struct{})
	for _, ch := range s {
		set[ch] = struct{}{}
	}

	for ch := range set {
		uw += ch
	}

	return
}

func Solve(reader *bufio.Reader) {
	yaReader := &YaReader{reader}
	s := readData(yaReader)

	res := sol(s)

	fmt.Println(res)
}

//func allUnique {
//
//}

func sol(s string) int {
	lenS := len(s)
	uc := uniqueChars(s)
	maxL := len(uc)
	var _ = maxL

	mp := make(map[uint8]bool)
	for i := 0; i < maxL; i++ {
		if mp[s[i]] {
			maxL -= 1
			break
		} else {
			mp[s[i]] = true
		}
	}

	for i := 0; i+maxL <= lenS; i++ {
		//if getHash(h, i, i + m - 1) == ht * pow[i]) {
		//	// обнаружено совпадение на позиции i
		//}
		//x := s[i]
	}

	return len(uc)
}

func readData(reader *YaReader) string {
	return reader.readString()
}

type YaReader struct {
	*bufio.Reader
}

func (reader *YaReader) readString() string {
	line, _ := reader.ReadString('\n')
	return strings.TrimRight(line, "\n")
}

func (reader *YaReader) readInt() int {
	line, _ := reader.ReadString('\n')
	res, _ := strconv.Atoi(strings.TrimRight(line, "\n"))
	return res
}

func openFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return file
}
