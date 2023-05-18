package main

func Sum(numbers []int) (sum int) {
	for _, n := range numbers {
		sum += n
	}
	return
}

func SumAll(numbersList ...[]int) (sumList []int) {
	for _, n := range numbersList {
		sumList = append(sumList, Sum(n))
	}
	return sumList
}

func SumAllTails(numbersList ...[]int) (sumList []int) {
	for _, n := range numbersList {
		if len(n) == 0 {
			sumList = append(sumList, 0)
		} else {
			sumList = append(sumList, Sum(n[1:]))
		}
	}
	return
}
