package insertsort

import "awesomeProject/utils"

func Sort(list []int, left, right int)  {
	for index := left;index <= right;index++ {
		if index > 0 {
			for i:=index;i>0;i-- {
				current := list[i]
				pre := list[i-1]
				if current <= pre {
					utils.Swap(list, i, i-1)
				} else {
					break
				}
			}
		}
	}
}
