package main

func Sum(numbers []int) (sum int) {
	for _, n := range numbers {
		sum += n
	}
	return
}
