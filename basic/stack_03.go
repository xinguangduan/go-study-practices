package main

import "fmt"

func testLarge() {
	nums1 := make([]int, 8192)
	nums2 := make([]int, 8192)
	fmt.Println(len(nums1), len(nums2))
}
func main() {
	testLarge()
}
