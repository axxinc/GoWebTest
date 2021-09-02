package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumber1(ch chan bool) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	ch <- true
}

func printLetters1(ch chan bool) {
	for i := 'A'; i < 'A'+10; i++ {
		fmt.Printf("%c ", i)
	}
	ch <- true
}

func printNumber2(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Microsecond)
		fmt.Printf("%d ", i)
	}
	wg.Done()
}

func printLetters2(wg *sync.WaitGroup) {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(time.Microsecond)
		fmt.Printf("%c ", i)
	}
	wg.Done()
}

func thrower(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("Threw >> ", i)
	}
}

func catcher(ch chan int) {
	for i := 0; i < 5; i++ {
		num := <-ch
		fmt.Println("Caught << ", num)
	}
}

// func print1() {
// 	printNumber1()
// 	printLetters1()
// }

// func goPrint1() {
// 	go printNumber1()
// 	go printLetters1()
// }

// func print2() {
// 	printNumber2()
// 	printLetters2()
// }

// func goPrint2() {
// 	go printNumber2()
// 	go printLetters2()
// }

func callerA(ch chan string) {
	ch <- "Hello World"
	close(ch)
}

func callerB(ch chan string) {
	ch <- "Hello Mundo"
	close(ch)
}

func main() {
	// var wg sync.WaitGroup
	// wg.Add(2)
	// go printNumber2(&wg)
	// go printLetters2(&wg)
	// wg.Wait()

	// ch1, ch2 := make(chan bool), make(chan bool)
	// go printLetters1(ch1)
	// go printNumber1(ch2)
	// <-ch1
	// <-ch2

	// // ch3 := make(chan int)
	// ch3 := make(chan int, 3)
	// go thrower(ch3)
	// go catcher(ch3)
	// time.Sleep(time.Second)
	a, b := make(chan string), make(chan string)
	go callerA(a)
	go callerB(b)

	// for i := 0; i < 5; i++ {
	// 	time.Sleep(time.Microsecond)
	// 	select {
	// 	case msg := <-a:
	// 		fmt.Printf("%s from A\n", msg)
	// 	case msg := <-b:
	// 		fmt.Printf("%s from B\n", msg)
	// 	default:
	// 		fmt.Println("No message")
	// 	}

	// }
	var msg string
	okA, okB := true, true
	for okA || okB {
		select {
		case msg, okA = <-a:
			if okA {
				fmt.Printf("%s form A\n", msg)
			}
		case msg, okB = <-b:
			if okB {
				fmt.Printf("%s form B\n", msg)
			}
		}
	}
}
