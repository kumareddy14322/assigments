package main

import (
	"fmt"
	"time"
)

func Mergesort(numbers []int) []int {
	l := len(numbers)
	if l <= 1 {
		return numbers
	}

	m := l / 2

	sortedLeft := Mergesort(numbers[0:m])
	sortedRight := Mergesort(numbers[m:l])

	return Merge(sortedLeft, sortedRight)
}

func Merge(left []int, right []int) []int {
	leftLength := len(left)
	rightLength := len(right)

	if leftLength == 0 {
		return right
	}
	if rightLength == 0 {
		return left
	}

	result := make([]int, (leftLength + rightLength))

	li := 0
	ri := 0
	resulti := 0
	var rnum, lnum int

	for li < leftLength || ri < rightLength {
		if li < leftLength && ri < rightLength {
			lnum = left[li]
			rnum = right[ri]

			if lnum <= rnum {
				result[resulti] = lnum
				li++
			} else {
				result[resulti] = rnum
				ri++
			}

		} else if li < leftLength {
			lnum = left[li]
			result[resulti] = lnum
			li++
		} else if ri < rightLength {
			rnum = right[ri]
			result[resulti] = rnum
			ri++
		}

		resulti++
	}

	return result
}

func MergeSortAsync(numbers []int, resultChan chan []int) {
	l := len(numbers)
	if l <= 1 {
		resultChan <- numbers
		return
	}

	m := l / 2

	lchan := make(chan []int, 1)
	rchan := make(chan []int, 1)

	go MergeSortAsync(numbers[0:m], lchan)
	go MergeSortAsync(numbers[m:l], rchan)
	go MergeAsync(<-lchan, <-rchan, resultChan)
}

func MergeAsync(left []int, right []int, resultChannel chan []int) {
	leftLength := len(left)
	rightLength := len(right)

	if leftLength == 0 {
		resultChannel <- right
		return
	}
	if rightLength == 0 {
		resultChannel <- left
		return
	}

	result := make([]int, (leftLength + rightLength))
	li := 0
	ri := 0
	resulti := 0
	var rnum, lnum int

	for li < leftLength || ri < rightLength {
		if li < leftLength && ri < rightLength {
			lnum = left[li]
			rnum = right[ri]

			if lnum <= rnum {
				result[resulti] = lnum
				li++
			} else {
				result[resulti] = rnum
				ri++
			}

		} else if li < leftLength {
			lnum = left[li]
			result[resulti] = lnum
			li++
		} else if ri < rightLength {
			rnum = right[ri]
			result[resulti] = rnum
			ri++
		}

		resulti++
	}

	resultChannel <- result
}

func MergesortDemo() {
	lim := 5
	largeArr := make([]int, lim)

	for i := 0; i < lim; i++ {
		largeArr[i] = lim - i
	}

	fmt.Println(largeArr)

	fmt.Println("Normal Mergesort")
	st1 := time.Now()
	r := Mergesort(largeArr)
	e1 := time.Now()
	fmt.Println(r)
	fmt.Println("Time: ", e1.UnixNano()-st1.UnixNano())

	fmt.Println("Mergesort with Goroutines")
	st2 := time.Now()
	resultChan := make(chan []int, 1)
	MergeSortAsync(largeArr, resultChan)
	k := <-resultChan
	e2 := time.Now()
	fmt.Println(k)
	fmt.Println("Time: ", e2.UnixNano()-st2.UnixNano())
}
