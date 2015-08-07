package main

import "fmt"

type testInt func(int) bool

func isOdd(i int) bool {
	return i%2 == 1
}

func isEven(i int) bool {
	return i%2 == 0
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, val := range slice {
		if f(val) {
			result = append(result, val)
		}
	}

	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 87, 9, 90, 8, 5, 4, 444, 32, 3, 23243, 42342}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd)

	fmt.Println(" odd elements:", odd)

	even := filter(slice, isEven)
	fmt.Println("even Elements:", even)

}
