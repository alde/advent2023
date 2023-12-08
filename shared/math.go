package shared

func GreatestCommonDivider(a, b int) int {
	for b > 0 {
		temp := b
		b = a % b // % is remainder
		a = temp
	}
	return a
}
func LeastCommonMultiple(nums []int) int {
	innerLCM := func(a, b int) int {
		return a * (b / GreatestCommonDivider(a, b))
	}

	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result = innerLCM(result, nums[i])
	}
	return result
}
