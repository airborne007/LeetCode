package sumofsubarrayminimums

import (
	"math"
)

// Solution1: Monotonic stack, multiple loop with space O(n)
func sumSubarrayMins1(arr []int) int {
	MOD := 1000000007
	n := len(arr)
	if n == 0 {
		return 0
	}

	left := make([]int, n)
	right := make([]int, n)

	stk := []int{}
	for i := 0; i < n; i++ {
		for len(stk) > 0 && arr[stk[len(stk)-1]] > arr[i] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 {
			left[i] = -1
		} else {
			left[i] = stk[len(stk)-1]
		}
		stk = append(stk, i)
	}

	stk = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(stk) > 0 && arr[stk[len(stk)-1]] >= arr[i] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 {
			right[i] = n
		} else {
			right[i] = stk[len(stk)-1]
		}
		stk = append(stk, i)
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans += (i - left[i]) * (right[i] - i) * arr[i]
	}
	return ans % MOD
}

// Solution2: Monotonic stack, onetime loop with space O(1)
func sumSubarrayMins(arr []int) int {
	MOD := 1000000007
	n := len(arr)
	if n == 0 {
		return 0
	}
	ans := 0
	stk := []int{}
	for i := -1; i <= n; i++ {
		for len(stk) > 0 && getItem(arr, stk[len(stk)-1], n) > getItem(arr, i, n) {
			cur := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			ans += arr[cur] * (cur - stk[len(stk)-1]) * (i - cur)
		}
		stk = append(stk, i)
	}
	return ans % MOD
}

func getItem(arr []int, idx, length int) int {
	if idx == -1 || idx == length {
		return math.MinInt
	}
	return arr[idx]
}
