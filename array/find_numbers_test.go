package array

import "testing"

func TestFindNumbers(t *testing.T) {
    // Test case 1: Empty array
    nums := []int{}
    expected := 0
    if output := findNumbers(nums); output != expected {
        t.Errorf("findNumbers(%v) = %d; want %d", nums, output, expected)
    }

    // Test case 2: Array with no even-digit numbers
    nums = []int{1, 3, 5, 7, 9}
    expected = 0
    if output := findNumbers(nums); output != expected {
        t.Errorf("findNumbers(%v) = %d; want %d", nums, output, expected)
    }

    // Test case 3: Array with some even-digit numbers
    nums = []int{12, 345, 2, 6, 7896}
    expected = 2
    if output := findNumbers(nums); output != expected {
        t.Errorf("findNumbers(%v) = %d; want %d", nums, output, expected)
    }

    // Test case 4: Array with all even-digit numbers
    nums = []int{10, 100, 1000, 10000}
    expected = 2
    if output := findNumbers(nums); output != expected {
        t.Errorf("findNumbers(%v) = %d; want %d", nums, output, expected)
    }
}
