package operations

func Addition(nums []int) (result int) {
	for _, num := range nums {
		result += num
	}

	return
}

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
