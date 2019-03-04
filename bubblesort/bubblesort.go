package bubblesort

import (
	"awesomeProject/utils"
)

func Sort(list []int, left, right int)  {
	if right == 0 {
		return
	}
	for index,num := range list {
		if index < right && num > list[index + 1] {
			utils.Swap(list, index, index + 1)
		}
	}
	Sort(list, left, right - 1)
}
