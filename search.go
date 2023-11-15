package main

import (
	"fmt"
)

func Search(list []int, element int) int {
	start, end := 0, len(list)-1

	for start <= end {
		mid := (start + end) / 2

		if list[mid] == element {
			return mid
		} else if element < list[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}

func SearchAppend(list []int, element int) int {
	if list[len(list)-1] < element {
		return len(list)
	}

	start, end, index := 0, len(list)-1, -1

	for start <= end {
		mid := (start + end) / 2
		index = mid

		if list[mid] == element {
			return -1
		} else if element < list[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return index
}
