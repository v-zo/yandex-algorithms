package main

type Array []int

func GetPrimes(n int) []int {
	var numbers Array = []int{0, 0}
	for i := 2; i <= n; i++ {
		numbers = append(numbers, i)
	}

	for num := 2; num <= n; num++ {
		if numbers[num] > 0 {
			for j := 2 * num; j <= n; j += num {
				numbers[j] = 0
			}
		}
	}

	return numbers.Filter(testNonZero)
}

func testNonZero(item int) bool {
	if item > 0 {
		return true
	}

	return false
}

func (arr Array) Filter(test func(int) bool) []int {
	var filteredArr []int

	for _, item := range arr {
		if test(item) {
			filteredArr = append(filteredArr, item)
		}
	}

	return filteredArr
}
