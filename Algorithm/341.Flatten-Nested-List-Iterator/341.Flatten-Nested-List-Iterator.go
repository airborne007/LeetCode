package flattennestedlistiterator

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

// Solution1: recursion
type NestedIterator1 struct {
	items []int
	idx   int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator1 {
	var dfs func([]*NestedInteger) []int
	dfs = func(ns []*NestedInteger) []int {
		ans := []int{}
		for _, n := range ns {
			if n.IsInteger() {
				ans = append(ans, n.GetInteger())
			} else {
				ans = append(ans, dfs(n.GetList())...)
			}
		}
		return ans
	}
	return &NestedIterator1{
		items: dfs(nestedList),
	}
}

func (n *NestedIterator1) Next() int {
	item := n.items[n.idx]
	n.idx++
	return item
}

func (n *NestedIterator1) HasNext() bool {
	return n.idx < len(n.items)
}

// Solution2: use stack
type NestedIterator struct {
	stack [][]*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{[][]*NestedInteger{nestedList}}
}

func (n *NestedIterator) Next() int {
	queue := n.stack[len(n.stack)-1]
	val := queue[0].GetInteger()
	n.stack[len(n.stack)-1] = queue[1:]
	return val
}

func (n *NestedIterator) HasNext() bool {
	for len(n.stack) > 0 {
		queue := n.stack[len(n.stack)-1]
		if len(queue) == 0 {
			n.stack = n.stack[:len(n.stack)-1]
			continue
		}
		item := queue[0]
		if item.IsInteger() {
			return true
		}
		n.stack[len(n.stack)-1] = queue[1:]
		n.stack = append(n.stack, item.GetList())
	}
	return false
}
