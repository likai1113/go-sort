package quicksort

import "awesomeProject/utils"

func Sort(list []int, left, right int)  {
	if right < left {
		return
	}
	flag := list[left]
	start := left
	end := right
	for {
		if start == end {
			break
		}
		for list[end] >= flag && end > start {
			end--
		}
		for list[start] <= flag && end > start {
			start++
		}
		if end > start {
			utils.Swap(list, start, end)
		}
	}
	utils.Swap(list, left, start)
	Sort(list, left, start - 1)
	Sort(list, start + 1, right)
}
