package main

import "fmt"

func main() {
	var arr [10]int
	fp := fmt.Println

	for i, _ := range arr {
		arr[i] = i + 1
		fp(i, arr[i])
	}
	for i := 0; i < len(arr); i++ {
		fp(arr[i])
	}
	fp(len(arr))

	arr1 := [2]string{"wwww", "weee"}
	var arr2 [2]string

	arr2[0] = "shs"
	arr2[1] = "ccc"
	fp(arr1)
	fp(arr2)
	fp("array type %T", arr2)
}
