package main

import "fmt"

func print(i interface{}) {
	fmt.Println(i)
}

func min(data []int) int {
	return min_acc(data, data[0])
}

func min_acc(data []int, acc int) int {
	if len(data) == 0 {
		return acc
	}
	return min_acc(data[1:], compare_min(acc, data[0]))
}

func compare_min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(data []int) int {
	return max_acc(data, 0)
}

func max_acc(data []int, acc int) int {
	if len(data) == 0 {
		return acc
	}
	return max_acc(data[1:], compare_max(acc, data[0]))
}

func compare_max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func sum(data []int, acc int) int {
	if len(data) == 0 {
		return acc
	}
	return sum(data[1:], acc+data[0])
}

type applied_func func(interface{}) bool

func any(data []int, fn applied_func) bool {
	for _, v := range data {
		if fn(v) {
			return true
		}
	}
	return false
}

func all(data []int, fn applied_func) bool {
	for _, v := range data {
		if !(fn(v)) {
			return false
		}
	}
	return true
}
