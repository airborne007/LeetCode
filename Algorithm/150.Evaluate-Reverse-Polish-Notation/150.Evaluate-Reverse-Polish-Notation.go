package evaluatereversepolishnotation

import "strconv"

func evalRPN(tokens []string) int {
	stk := []int{}
	for _, str := range tokens {
		num, err := strconv.Atoi(str)
		if err == nil {
			stk = append(stk, num)
		} else {
			num := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			switch str {
			case "+":
				stk[len(stk)-1] += num
			case "-":
				stk[len(stk)-1] -= num
			case "*":
				stk[len(stk)-1] *= num
			case "/":
				stk[len(stk)-1] /= num
			}
		}
	}
	return stk[0]
}
