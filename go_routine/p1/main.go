/*
Write a program to print even and odd number from 2 goroutine ex, go routine one will going to print even and go routine 2 will going to print odd
*/

package main

import (
	"fmt"
	"sync"
)

func EvenNum(ch chan<- int, num []int) {
	defer close(ch)
	for v := range num {
		if v%2 == 0 {
			ch <- v
		}
	}
}

func OddNum(ch chan<- int, num []int) {
	defer close(ch)
	for v := range num {
		if v%2 != 0 {
			ch <- v
		}
	}
}

func PrintEven(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Println("Even Numbers :", v)
	}

}

func PrintOdd(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Println("Odd Numbers :", v)
	}
}

func main() {
	ch := make(chan int, 10)
	ch2 := make(chan int, 10)
	var wg sync.WaitGroup
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	wg.Add(2)
	go EvenNum(ch, num)
	go OddNum(ch2, num)

	go PrintEven(ch, &wg)
	go PrintOdd(ch2, &wg)

	wg.Wait()

}
