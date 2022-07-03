package climbingstairs

func climbStairs(n int) int {
	temp := make([]int, n+2)
	temp[0] = 1
	temp[1] = 2
	for i := 2; i < n; i++ {
		temp[i] = temp[i-1] + temp[i-2]
	}
	return temp[n-1]
}
