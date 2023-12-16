package array

import (
	"testing"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{1, 1, 0, 1, 1, 1},
			want: 3,
		},
		{
			name: "Example 2",
			nums: []int{1, 0, 1, 1, 0, 1},
			want: 2,
		},
		// Add more test cases here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindMaxConsecutiveOnes(tt.nums)
			if got != tt.want {
				t.Errorf("FindMaxConsecutiveOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}
