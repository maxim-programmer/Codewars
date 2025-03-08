package kata

func Maps(x []int) []int {
	result := make([]int, len(x))
	for i, z := range x {
		result[i] = 2 * z
	}
	return result
}
