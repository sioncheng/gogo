package algr

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		result := make([]int, n+1)
		result[0] = 0
		result[1] = 1
		for i := 2; i < len(result); i++ {
			result[i] = result[i-1] + result[i-2]
		}
		return result[n]
	}
}
