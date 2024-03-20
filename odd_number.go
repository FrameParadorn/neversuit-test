package main

func FindOddNumber(text []int) int {
	// TODO : start your code here
	for number, count := range countNumberToMap(text) {
		if count%2 != 0 {
			return number
		}
	}

	return 0

}

func countNumberToMap(numbers []int) map[int]int {
	results := map[int]int{}

	for _, n := range numbers {
		results[n]++
	}

	return results

}
