package maximizesumofarrayafterknegations

import "sort"

func largestSumAfterKNegations(nums []int, k int) int {
	ans := 0
	sort.Ints(nums)
	for i := range nums {
		if nums[i] < 0 && k > 0 {
			nums[i] *= -1
			k--
		}
		ans += nums[i]
	}

	if k%2 == 1 {
		sort.Ints(nums)
		ans -= 2 * nums[0]
	}

	return ans
}
