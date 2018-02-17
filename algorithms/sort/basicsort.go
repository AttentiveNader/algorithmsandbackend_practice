package main

import (
	"fmt"
)

var low int = 0 

func SortSel(slice []int) []int {
	var high int = len(slice) 
	
	for i := low ; i < high ; i++{
		iterator := slice[i]
		index := 0
		incase := false
		for  k := low ; k < high  ; k++{
			if k > i{
				if slice[k] < iterator{
					iterator = slice[k]
					index = k
					fmt.Println(iterator,slice[k],slice[i])
					incase = true
				}
			}
		}
		if incase {
			slice[index] = slice[i]
			slice[i] = iterator
		}
		//slice[index] = tem
		fmt.Println("slice",slice,slice[i],slice[index],iterator)
	}
	return slice
}
func BubbleSort(slice []int) []int {
	for {
		swapped := false
		for i := 0 ; i < len(slice) ;i++ {
			if i != len(slice) -1 {
				if slice[i] > slice[i+1] {
					slice[i],slice[i+1] = slice[i+1],slice[i]
					swapped = true
				}
			}
			fmt.Println("1BubbleSort",slice)
		}
		fmt.Println("BubbleSort",slice)
		if !swapped{
			break
		}
	}
	return slice
}

func sorting() {
	var slice []int = []int{4,9,7,6,5,3,1}
	fmt.Println(SortSel(slice))
	fmt.Println(BubbleSort(slice))
}