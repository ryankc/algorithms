package main

import (
	"fmt"
	"github.com/ryankc/algorithms/bagsstacksqueues"
)

func main(){
	stack := bagsstacksqueues.NewStack(2)
	fmt.Println(stack.Size())
	fmt.Println(stack.IsEmpty())
}










/*
func main(){
	list := [][]int{[]int{1,2},[]int{2,3},[]int{4,5},[]int{6,7}}
	fmt.Println(Flatten(list))
}

//Reverses a slice
func Reverse(list []int) []int{
	list2 := make([]int,len(list))
	for num,val := range list {
		list2[len(list)-1-num] = val
	}
	return list2
}

//Finds Kth element of slice
func PrintK(list []int, k int) int{
	return list[k-1]
}

// Checks whether function is palindromic
func Palindrome(list []int) bool{
	//Removes middle value from palindrome for odd slice length
	if len(list) % 2 != 0{
		list = append(list[:(len(list)/2)],list[len(list)/2+1:]...)
	}
	//Evaluates both sides of slice
	a := Reverse(list[(len(list)/2):])
	b := list[:(len(list)/2)]
	for i := 0; i < len(list)/2; i++{
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

//Recursive? or use variadic function as solution? is it even possible, or are both required

func Flatten(list [][]int)[]int{
	var list2 []int

	for _,val := range list{
		switch val.(type){
		case int:
			list2 = append(list2, val)
		case []int:

		default:
			list2 = append(list2 ,Flatten(val)...)

		}
	}
	return list2

}

*/