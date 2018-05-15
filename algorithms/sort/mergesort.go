package sort

import (
	"fmt"
)

var slice []int

func MergeTwoSlices(slice1, slice2 []int) []int {
	n := (len(slice1) + len(slice2))
	slice = []int{}
	index1 := 0
	index2 := 0
	breac1 := true
	breac2 := true
	for i := 0; i < n; i++ {
		if slice2[index2] <= slice1[index1] && breac2 {
			slice = append(slice, slice2[index2])
			index2++
			if len(slice2) <= index2 {
				breac2 = false
				index2 = (len(slice2) - 1)
			}
		} else if slice1[index1] < slice2[index2] && breac1 {
			slice = append(slice, slice1[index1])
			index1++
			if len(slice1) <= index1 {
				breac1 = false
				index1 = (len(slice1) - 1)
			}
		} else if !breac1 {
			slice = append(slice, slice2[index2:]...)
			break
		} else if !breac2 {
			slice = append(slice, slice1[index1:]...)
			break
		}
	}
	return slice
}
func SliceSplitter(slice []int) [][]int {
	n := len(slice)
	slices := [][]int{}
	slice1 := []int{}
	slice2 := []int{}
	if n%2 == 0 {
		slice1 = slice[:(n / 2)]
		slice2 = slice[(n / 2):]
	} else {
		slice1 = slice[:((n + 1) / 2)]
		slice2 = slice[((n + 1) / 2):]
	}
	if len(slice1) <= 2 {
		if len(slice1) == 2 {
			if slice1[0] > slice1[1] {
				slice1[0], slice1[1] = slice1[1], slice1[0]
			}
		}
		slices = append(slices, slice1)
	} else {
		subslices := SliceSplitter(slice1)
		slices = append(slices, subslices...)
	}
	if len(slice2) <= 2 {
		if len(slice2) == 2 {
			if slice2[0] > slice2[1] {
				slice2[0], slice2[1] = slice2[1], slice2[0]
			}
		}
		slices = append(slices, slice2)
	} else {
		subslices := SliceSplitter(slice2)
		slices = append(slices, subslices...)
	}
	return slices
}
func SliceMerger(slices [][]int) []int {
	slice := []int{}
	slices1 := [][]int{}
	slices2 := [][]int{}
	slice1 := []int{}
	slice2 := []int{}
	n := len(slices)
	if n > 2 {
		if n%2 == 0 {
			slices1 = slices[:(n / 2)]
			slices2 = slices[(n / 2):]
		} else {
			slices1 = slices[:((n + 1) / 2)]
			slices2 = slices[((n + 1) / 2):]
		}
		slice1 = SliceMerger(slices1)
		slice2 = SliceMerger(slices2)
		slice = MergeTwoSlices(slice1, slice2)
	} else {
		if n == 2 {
			slice = MergeTwoSlices(slices[0], slices[1])
		} else {
			slice = slices[0]
		}
	}
	return slice
}
func MergeSort(slice []int) []int {
	slices := SliceSplitter(slice)
	slice = SliceMerger(slices)
	return slice
}
func main() {
	slice := []int{11, 1, 2, 10, 8, 4, 3, 6, 7, 5, 14, 9}
	fmt.Println(MergeSort(slice))
}
