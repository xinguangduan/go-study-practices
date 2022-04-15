package main

import "fmt"

func main() {
	res := AddSum(1, 3, 4, 6)
	fmt.Println("res:", res)
	fmt.Println("res:", AddSum())
	mums := []int{1, 2, 3, 4}
	fmt.Println("array:", AddSum(mums...))

	loop(1, 2, 3, 4, 5, 6)
}

func AddSum(nums ...int) (sum int) {
	fmt.Printf("%T\n", nums)
	for _, value := range nums {
		sum += value
	}
	return
}
func loop(nums ...int) {
	sum := 0
	for _, value := range nums {
		sum += value
	}
	fmt.Println(sum)
}

func Variable(v1 ...int) {

}
