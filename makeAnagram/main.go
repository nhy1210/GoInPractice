package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
func makeAnagram(a string, b string) int {
	var arr1, arr2 [26]int
	ret := 0
	for index1 := 0; index1 < len(a); index1++ {
		arr1[a[index1]-'a']++
	}
	for index2 := 0; index2 < len(b); index2++ {
		arr2[b[index2]-'a']++
	}
	for j := 0; j < 26; j++ {
		if arr1[j] != arr2[j] {
			ret += abs(arr1[j] - arr2[j])
		}
	}
	return ret
}
func main() {

	a := "cde"
	b := "abc"
	ret := makeAnagram(a, b)
	fmt.Println("Make anagram need to remove : ", ret, "elements")

}
