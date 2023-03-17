package utils

// Addition takes a slice of integers nums as input and returns the sum of all the
// integers in the slice
func Addition(nums []int) (result int) {
	for _, num := range nums {
		result += num
	}

	return
}

// Subtraction takes a slice of integers nums as input and returns the result of
// subtracting all the integers in the slice from the first integer in the slice.
// If the slice is empty, the function returns 0
func Subtraction(nums []int) (result int) {
	if len(nums) == 0 {
		return 0
	}

	result = nums[0]

	for i := 1; i < len(nums); i++ {
		result -= nums[i]
	}

	return
}

// This function takes a slice of integers nums as input and returns the result
// of multiplying all the integers in the slice together. If the slice is empty,
// the function returns 0
func Multiplication(nums []int) (result int) {
	if len(nums) == 0 {
		return 0
	}

	result = 1

	for _, num := range nums {
		result *= num
	}

	return
}
