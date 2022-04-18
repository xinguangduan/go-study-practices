package main

import "fmt"

func main() {
	sum, count, avg := GetScore(98.23, 87.22, 83.5, 90.5)
	fmt.Printf("%.2f\n", sum)
	fmt.Printf("%v\n", count)
	fmt.Printf("%.2f\n", avg)

	input := []float64{98.23, 97.22, 93.5, 99.5}
	name := ""
	sum, count, avg, name = GetScores("张三", input...)
	fmt.Printf("%.2f\n", sum)
	fmt.Printf("%v\n", count)
	fmt.Printf("%.2f\n", avg)
	fmt.Printf("%s\n", name)
}

func GetScore(scores ...float64) (sum float64, count int, avg float64) {
	for _, value := range scores {
		sum += value
		count++
	}
	avg = sum / float64(count)
	return
}

func GetScores(name string, scores ...float64) (sum float64, count int, avg float64, user string) {
	for _, value := range scores {
		sum += value
		count++
	}
	avg = sum / float64(count)
	user = name
	return
}
