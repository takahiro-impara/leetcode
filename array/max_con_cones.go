package array

func FindMaxConsecutiveOnes(nums []int) int {
	maxOnes := 0
	temp := 0
	for _, v := range nums {
		if v == 1 {
			temp += 1
			if temp > maxOnes {
				maxOnes = temp
			}
		} else {
			temp = 0
		}
	}
	return maxOnes
}
