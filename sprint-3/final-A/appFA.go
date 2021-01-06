/*

посылка 46162039

-- ПРИНЦИП РАБОТЫ --
	Первая идея: сначало найти индекс j наименьшего элемента бинарным поиском.
Создать новый массив у которого все индексы сдвинуты на j позиций. Таким образом
данный массив будет отсортирован и в нем можно снова запустить процедуру бинарного
поиска но на это раз искомого элемента k.
Данный подход удовлетворяет условию задачи т.к. 2 бинарных поиска это O(log n) + O(log n)
в итоге всеравно O(log n).
	Однако алгоритм можно ускорить, если приступить к поиску k сразу. Для этого при выборе
центральной точки mid нужно выполнить нескольнолько дополнительных проверок, нежели просто
сравнение с k. Существует 6 вариантов взаимного расположения mid и k. Все они схематично изображены
в комеентариях кода функции getSide. В зависимости от варианта вызывается бинарный поиск в "левой"
или "правой" половинке, а далеее алгоритм повторяет обычный бинарный поиск.
Хотя этот способ быстрее первого подхода сложность всеравно та же - O(log n)

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
	По аналогии с обычным бинарным поиском в остортированном массиве, мы имеем рекурсивное
деление "отрезка" пополам, в следствие чего базовый случай когда длина отрезка 1 или 0 неминуем.
Если же искомый элемент k в массиве присутствует то рекурсия завершится раньше - это тоже базовый
случай.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
	Сложность та же, что и для обычного бинарного поиска.
На каждой рекурсии мы делим текущий отрезок пополам за O(1).
Т.е. на первом шаге длина отрезка будет равна n/2, на втором - n/4, затем n/8, ...
В худшем случае на каком то i-ом шаге мы сойдемся к единственному элементу: 1 == n/2^i, или 2^i==n,
откуда получим что нам нужно совершить i = log2(n) рекурсий, сложностью O(1) каждая.
Итоговая временная сложность - O(log n).

*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	file := openFile("input.txt")
	defer file.Close()

	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(os.Stdout)

	Solve(reader, writer)
}

func Solve(reader *bufio.Reader, writer *bufio.Writer) {
	n, k, elements := readData(reader)

	result := binarySearch(elements, 0, n-1, k)

	writer.WriteString(strconv.Itoa(result))
	writer.WriteString("\n")

	writer.Flush()
}

type IntList struct {
	str []string
}

func (list IntList) getEl(i int) (value int) {
	value, _ = strconv.Atoi(list.str[i])
	return
}

type Side int

const (
	undef Side = iota
	left
	right
)

func binarySearch(elements IntList, st int, end int, k int) int {
	stEl := elements.getEl(st)
	endEl := elements.getEl(end)

	if endEl == k {
		return end
	}

	if stEl == k {
		return st
	}

	if end-st <= 1 {
		return -1
	}

	mid := (st + end) / 2
	midEl := elements.getEl(mid)

	if midEl == k {
		return mid
	}

	switch getSide(stEl, midEl, endEl, k) {
	case left:
		return binarySearch(elements, st, mid-1, k)
	case right:
		return binarySearch(elements, mid+1, end, k)
	default:
		return -1
	}
}

func getSide(stEl int, midEl int, endEl int, k int) (side Side) {
	if midEl < endEl {
		if k < midEl {
			side = left //   * * * * * ∙ 🏁 ∙ 👆 ∙ ∙ ∙ ∙
		} else {
			if k < endEl {
				side = right //   * * * * * ∙ ∙ ∙ 👆  ∙ ∙ 🏁 ∙
			} else {
				side = left //   * * 🏁 * * ∙ ∙ ∙ 👆 ∙ ∙ ∙ ∙
			}
		}
	} else {
		if k > midEl {
			side = right //   * * * * 👆 * 🏁 * ∙ ∙ ∙ ∙ ∙
		} else {
			if k > stEl {
				side = left //   * 🏁 * * 👆 * * * ∙ ∙ ∙ ∙ ∙
			} else {
				side = right //   * * * * 👆 * * * ∙ ∙ 🏁 ∙ ∙
			}
		}
	}

	if side == undef {
		panic("side is undefined")
	}

	return
}

func readData(reader *bufio.Reader) (n int, k int, elements IntList) {
	line1, _ := reader.ReadString('\n')
	line2, _ := reader.ReadString('\n')
	line3, _ := reader.ReadString('\n')

	n, _ = strconv.Atoi(strings.TrimRight(line1, "\n"))
	k, _ = strconv.Atoi(strings.TrimRight(line2, "\n"))
	elements = IntList{strings.Fields(line3)}

	return
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}
