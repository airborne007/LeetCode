package miniparser

import (
	"strconv"
	"unicode"
)

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */

type NestedInteger struct{}

func (n *NestedInteger) SetInteger(int) int { return 0 }
func (n *NestedInteger) Add(NestedInteger)  {}

// Solution1: DFS
func deserialize(s string) *NestedInteger {
	idx := 0
	var dfs func() *NestedInteger
	dfs = func() *NestedInteger {
		ni := &NestedInteger{}
		if s[idx] == '[' {
			idx++
			for s[idx] != ']' {
				ni.Add(*dfs())
				if s[idx] == ',' {
					idx++
				}
			}
			idx++
			return ni
		}

		isNegative := s[idx] == '-'
		if isNegative {
			idx++
		}
		num := 0
		for ; idx < len(s) && unicode.IsDigit(rune(s[idx])); idx++ {
			num = num*10 + int(s[idx]-'0')
		}
		if isNegative {
			num = -num
		}
		ni.SetInteger(num)
		return ni
	}
	return dfs()
}

// Solution2: use stack
func deserialize1(s string) *NestedInteger {
	if s[0] != '[' {
		num, _ := strconv.Atoi(s)
		ni := &NestedInteger{}
		ni.SetInteger(num)
		return ni
	}

	stk := []*NestedInteger{}
	num, isNegative := 0, false
	for i, ch := range s {
		if ch == '[' {
			stk = append(stk, &NestedInteger{})
		} else if ch == ']' || ch == ',' {
			if unicode.IsDigit(rune(s[i-1])) {
				if isNegative {
					num = -num
				}
				ni := NestedInteger{}
				ni.SetInteger(num)
				stk[len(stk)-1].Add(ni)
			}
			num, isNegative = 0, false
			if ch == ']' && len(stk) > 1 {
				stk[len(stk)-2].Add(*stk[len(stk)-1])
				stk = stk[:len(stk)-1]
			}
		} else if ch == '-' {
			isNegative = true
		} else if unicode.IsDigit(ch) {
			num = num*10 + int(ch-'0')
		}
	}
	return stk[0]
}
