package main

import (
	"sort"
	"slices"
)

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func Sort(numbers []int) []int {
	sort.Ints(numbers)
	return numbers
}

func SortAll(numbersToSort ...[]int) []int {
	var sorted []int
	sorted = slices.Concat(numbersToSort...)
	return Sort(sorted)
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums[]int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
		}
	}
	return sums
}
