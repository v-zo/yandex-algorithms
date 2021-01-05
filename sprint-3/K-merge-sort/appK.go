package main

func merge(arr []int, lf int, mid int, rg int) (result []int) {
	i := lf
	j := mid
	var res []int

	for {
		if arr[i] < arr[j] {
			res = append(res, arr[i])
			i++
		} else {
			res = append(res, arr[j])
			j++
		}

		if i == mid {
			result = append(res, arr[j:rg]...)
			return
		}

		if j >= rg {
			result = append(res, arr[i:mid]...)
			return
		}
	}
}

func merge_sort(arr []int, lf int, rg int) {
	if rg == lf || rg-lf == 1 || len(arr) == 1 {
		return
	}

	mid := (rg + lf) / 2

	merge_sort(arr, lf, mid)
	merge_sort(arr, mid, rg)

	merged := merge(arr, lf, mid, rg)
	for i, m := range merged {
		arr[i+lf] = m
	}
}
