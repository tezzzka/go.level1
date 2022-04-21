package main

import "fmt"

func main() {
	arr := []int{5, 9, 0, 1, 2}
	Quicksort(arr)
	fmt.Println(arr)
}

func Quicksort(ar []int) {
	if len(ar) <= 1 {
		return
	}

	split := partition(ar)

	Quicksort(ar[:split])
	Quicksort(ar[split:])
}

func partition(ar []int) int {
	pivot := ar[len(ar)/2]

	left := 0
	right := len(ar) - 1

	for {
		for ; ar[left] < pivot; left++ {
		}

		for ; ar[right] > pivot; right-- {
		}

		if left >= right {
			return right
		}

		swap(ar, left, right)
	}
}
func swap(ar []int, i, j int) {
	tmp := ar[i]
	ar[i] = ar[j]
	ar[j] = tmp
}
