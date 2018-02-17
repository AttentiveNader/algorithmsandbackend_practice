package main

import (
	"fmt"
)
var slice []int

func MergeSlices(zsli,bsli []int) []int {
	n := (len(zsli) + len(bsli))
	slice = []int{}
	index1 := 0
	index2 := 0
	breac1 := true
	breac2 := true
	for i := 0 ; i < n ; i++  {
		if	bsli[index2] <= zsli[index1] && breac2{
			slice = append(slice,bsli[index2])
			index2++
			if len(bsli) <= index2 {
				breac2 = false
			}
		}else if zsli[index1] < bsli[index2] && breac1 {
			slice = append(slice,zsli[index1])
			index1++
			if len(zsli) <= index1  {
				breac1 = false 
				index1 = (len(zsli)-1)
			}
		}else if !breac1 {
			slice = append(slice,bsli[index2])
			index2++
		}else if !breac2{
			slice = append(slice,zsli[index1])
			index1++
		}
	}
	return slice
}
func SliceSplitter(slice []int) [][]int{
	n := len(slice)
	slices := [][]int{}
	slice1 := []int{}
	slice2 := []int{}
	if n % 2 == 0 {
		slice1 = slice[:(n / 2)]
		slice2 = slice[(n / 2):]
	}else{
		slice1 = slice[:((n+1) / 2)]
		slice2 = slice[((n+1) / 2):]
	}
	if len(slice1) <= 2  {
	  slices = append(slices,slice1)
	}else{
		SliceSplitter(slice1)
		subslices := SliceSplitter(slice2)
		for i := 0; i < len(subslices); i++ {
			slices = append(slices,subslices[i])
		}
	}
	if len(slice2) <= 2  {
	  slices = append(slices,slice2)
	}else{
		subslices := SliceSplitter(slice2)
		for i := 0; i < len(subslices); i++ {
			slices = append(slices,subslices[i])
		}
	}
	return slices
	//SliceSplitter()
}
func MergeSort(slice []int) []int {
	slices := SliceSplitter(slice)
	for i := 0; i < len(slices); i++ {
		if len(slices[i]) > 1 {
			if slices[i][1] < slices[i][0]{
				 slices[i][1],slices[i][0] = slices[i][0],slices[i][1]
				 fmt.Println(slice[i])
			}
		}
	}
 	return slice
}

func main() {
	 //zsli := []int{4,3,6,7,5}
	// slice := []int{11,1,2,10,8,4,3,6,7,5,14}
	slice:=[]int{8,7,2,4,5,1}
	fmt.Println(MergeSort(slice))
}