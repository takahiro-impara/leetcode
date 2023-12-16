package array

import "strconv"

func findNumbers(nums []int) int {
	num := 0
	for _, v := range nums {
		if len(strconv.Itoa(v))%2 == 0 {
			num += 1
		}
	}
	return num
}
