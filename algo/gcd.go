package algo

func gcd[T int | int8 | int16 | int32 | int64 | uint | uint16 | uint32 | uint64](a, b T) T {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func GCD[T int | int8 | int16 | int32 | int64 | uint | uint16 | uint32 | uint64](nums []T) T {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result = gcd(result, nums[i])
	}
	return result
}
