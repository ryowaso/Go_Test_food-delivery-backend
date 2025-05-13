package main

func lengthAfterTransformations(s string, t int) int {
	const mod = 1e9 + 7
	dp := make([][]int, t+1)
	for i := range dp {
		dp[i] = make([]int, 26)
	}
	for j := 0; j < 26; j++ {
		dp[0][j] = 1
	}

	for i := 1; i <= t; i++ {
		for j := 0; j < 26; j++ {
			if j == 25 {
				dp[i][j] = (dp[i-1][0] + dp[i-1][1]) % mod
			} else {
				dp[i][j] = dp[i-1][j+1]
			}
		}
	}
	res := 0
	for _, v := range s {
		res = (res + dp[t][v-'a']) % mod
	}
	return res
}
