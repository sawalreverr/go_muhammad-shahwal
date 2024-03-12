package main

import (
	"fmt"
	"slices"
	"strings"
	"sync"
	"time"
)

func main() {
	fmt.Println("Task 01")
	Task01("kasur")

	fmt.Println("\nTask 02")
	fmt.Println(Task02())

	fmt.Println("\nTask 03")
	fmt.Println(Task03())
}

func Task01(word string) {
	slc := strings.Split(word, "")
	slices.Reverse(slc)

	go func() {
		for i := 0; i < len(slc); i++ {
			for j := 0; j <= i; j++ {
				fmt.Printf(slc[j])
			}
			fmt.Println()
		}
	}()

	time.Sleep(3 * time.Second)
}

func Task02() []int {
	var result []int
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 2; i <= 100; i++ {
			isPrime := true
			for j := 2; j*j <= i; j++ {
				if i%j == 0 {
					isPrime = false
					break
				}
			}

			if isPrime {
				ch <- i
			}
		}
		close(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range ch {
			result = append(result, num*num)
		}
	}()

	wg.Wait()

	return result
}

func Task03() []string {
	var result []string
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 2; i <= 100; i++ {
			isCom := false
			for j := 2; j*j <= i; j++ {
				if i%j == 0 {
					isCom = true
					break
				}
			}
			if isCom {
				ch <- i
			}
		}

		close(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range ch {
			if num*num%2 == 0 {
				result = append(result, "PING")
			} else {
				result = append(result, "PONG")
			}
		}
	}()

	wg.Wait()
	return result
}
