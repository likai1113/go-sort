package utils

func Swap(list []int, i, j int)  {
	tmp := list[i]
	list[i] = list[j]
	list[j] = tmp
}
