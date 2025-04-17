package main

import (
	"fmt"
	"time"
)

// 1
func hello(name string) {
	fmt.Print("Hello", name)
}

// 2
func sum(nums ...int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}

func swap(a, b string) (string, string) {
	return b, a
}

func divide(a, b int) (result int, err error) {
	if b == 0 {
		err = fmt.Errorf("школьная программа по математике не пройдена")
		return
	}
	result = a / b
	return

}

func average(nums ...float64) float64 {
	var sum float64
	for _, i := range nums {
		sum += i
	}
	return sum / float64(len(nums))
}

func greet(name string) (greeting string, message string) {
	greeting = "Hello " + name
	message = "\nCurrent date and time: " + time.Now().Format("02-01-2006 15:04")
	return
}

func product(nums ...int) int {
	result := 1
	for _, i := range nums {
		result *= i
	}
	return result
}

func reverse(a, b int) (rA, rB int) {
	rA = b
	rB = a
	return
}
