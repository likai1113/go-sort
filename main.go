package main

import (
	"awesomeProject/heapsort"
	"fmt"
)

func main()  {
	list := []int{8, 3, 6, 9, 11, 2, 7, 23, 65, 13, 9}
	fmt.Println("start: ", list)
	//quicksort.Sort(list, 0, len(list) - 1)
	//bubblesort.Sort(list, 0, len(list) - 1)
	//selectsort.Sort(list, 0, len(list) - 1)
	//insertsort.Sort(list, 0, len(list) - 1)
	//shellsort.Sort(list, 0, len(list) - 1)
	//result := mergesort.Sort(list, 0, len(list) - 1)
	//result := bucketsort.Sort(list)
	//result := countsort.Sort(list)
	//radixsort.Sort(list)
	heapsort.Sort(list)
	fmt.Println("end: ", list)
}
