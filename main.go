package main

import "fmt"

func main() {
	var age int = 21
	var pi float64 = 3.14
	const country = "India"

	name := "Vansh"
	isLearning := true

	sum := add(3, 5)

	q, err := divide(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Quotient:", q)
	}

	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}

	numbers := []int{1, 2}
	numbers = append(numbers, 3)
	fmt.Println(numbers)

	fmt.Println(age, pi, country, name, isLearning, sum)

	a := [3]int{1, 2, 3}
	changeSlice(a[:])
	fmt.Println(a)

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result)
	MapDemo()
	structDemo()
	goroutineDemo()
	channelDemo()
	bufferedChannelDemo()
	WaitGroupDemo()
}
