package numberofstudentdoinghomeworkatagiventime

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	ans := 0
	for idx := range startTime {
		if startTime[idx] <= queryTime && endTime[idx] >= queryTime {
			ans++
		}
	}
	return ans
}
