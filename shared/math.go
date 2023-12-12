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

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func MinMax(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

func Sum(nums []int) int {
	res := 0
	for _, n := range nums {
		res += n
	}
	return res
}
