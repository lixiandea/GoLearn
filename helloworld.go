package main

import "fmt"

func sum(values []int, resultChan chan int) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	resultChan <- sum

}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	resultChan := make(chan int, 2)

	go sum(values[:len(values)/2], resultChan)
	go sum(values[len(values)/2:], resultChan)
	go sum(values[len(values)/4:len(values)/2], resultChan)
	sum1, sum2, sum3 := <-resultChan, <-resultChan, <-resultChan

	fmt.Println("result :", sum1, sum2, sum3)
}
