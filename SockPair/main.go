package main

import (
	"fmt"
	"sort"
)

func findpair(arr []int, n int) int {
	//We have builtin func //sort.Ints  //sort.Float64s   //sort.Strings
	sort.Ints(arr)
	ret := 0
	for i:=0; i< n-1; i++{
		if arr[i] == arr[i+1] {
			i++;
			ret++
		}
	}
	return ret

}
/*
	In the case array is declared by int32 type, we don't have built in function, so we need to create it ourselves
	func findpair(arr []int32, len int) int{

		
	}



*/
func main(){

	arr := []int{2,3,4,2,3,5,7,8,5}
	//find pair of sock
	n:= len(arr)
	ret := findpair(arr, n)
	fmt.Println(ret)

}