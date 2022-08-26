package scoreofparentheses

// Solution1: recursion, time O(n^2), space O(n)
func scoreOfParentheses1(s string) int {
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		var ans, bal int
		for k := i; k < j; k++ {
			if s[k] == '(' {
				bal += 1
			} else {
				bal -= 1
			}
			if bal == 0 {
				// only have "()"
				if k-i == 1 {
					ans += 1
				} else {
					ans += 2 * dfs(i+1, k)
				}
				i = k + 1
			}
		}
		return ans
	}
	return dfs(0, len(s))
}

// Solution2: use stack, time O(n), space O(n)
func scoreOfParentheses2(s string) int {
	stk := []int{0}
	for _, ch := range s {
		if ch == '(' {
			stk = append(stk, 0)
		} else {
			v := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			stk[len(stk)-1] += max(2*v, 1)
		}
	}
	return stk[0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Solution3: mathematics, time O(n), space O(1)
func scoreOfParentheses(s string) int {
	var ans, bal int
	for i, ch := range s {
		if ch == '(' {
			bal += 1
		} else {
			bal -= 1
			if s[i-1] == '(' {
				ans += 1 << bal
			}
		}
	}
	return ans
}
