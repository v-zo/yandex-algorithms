package main

func merge(arr []int, lf int, mid int, rg int) (result []int) {
	i := lf
	j := mid
	res := append([]int(nil), arr[:lf]...)

	for {
		if arr[i] < arr[j] {
			res = append(res, arr[i])
			i++
		} else {
			res = append(res, arr[j])
			j++
		}

		if i == mid {
			result = append(res, arr[j:]...)
			return
		}

		if j > rg {
			result = append(res, arr[i:mid]...)
			if j <= len(arr)-1 {
				result = append(result, arr[j:]...)
			}

			return
		}
	}
}

func merge_sort(arr []int, lf int, rg int) {
	if rg == lf {
		return
	}

	if rg-lf == 1 {
		if arr[lf] > arr[rg] {
			arr[lf], arr[rg] = arr[rg], arr[lf]
		}
		return
	}

	mid := (rg + lf) / 2

	merge_sort(arr, lf, mid-1)
	merge_sort(arr, mid, rg)

	merged := merge(arr, lf, mid, rg)
	for i, m := range merged {
		arr[i] = m
	}
}

func main() {
	//arr := &[]int{3,2,1}
	//merge_sort(arr,0,2)
	//fmt.Println(*arr)

	//arr1 := &[]int{6,5,7}
	//merge_sort(arr1,0,2)
	//fmt.Println(*arr1)

	//arr2 := &[]int{2,5,6,1,3,4}
	//merge_sort(arr2,0,5)
	//fmt.Println(*arr2)
	//
	//arr3 := &[]int{39,28,44,4,10,83,11}
	//merge_sort(arr3,0,6)
	//fmt.Println(*arr3)
	//
	//arr4 := &[]int{4, 5, 3, 0, 1, 2, 3}
	//merge_sort(arr4,1,4)
	//fmt.Println(*arr4)
}
