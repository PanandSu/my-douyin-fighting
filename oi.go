package main

import (
	"fmt"
)

func main() {
	var t, n, m int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&n, &m)
		arr := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&arr[i])
		}
		sum := f(arr, m)
		fmt.Printf("%f\n", float64(sum)/float64(m))
	}
}

// 长度为k的连续子数组的最大和
func f(arr []int, k int) int {
	if k <= 0 || k > len(arr) {
		return 0
	}
	sum := 0
	for i := range k {
		sum += arr[i]
	}
	maxi := sum
	for i := k; i < len(arr); i++ {
		tmp := sum + arr[i] - arr[i-k]
		maxi = max(maxi, tmp)
	}
	return maxi
}
