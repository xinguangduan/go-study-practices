package main

import "fmt"

func main() {
	var arr [10]int

	for i, _ := range arr {
		arr[i] = i + 1
		fmt.Println(i, arr[i])
	}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	fmt.Println(len(arr))

	arr1 := [2]string{"wwww", "weee"}
	var arr2 [2]string

	arr2[0] = "shs"
	arr2[1] = "ccc"
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Printf("array type %T", arr2)
}
