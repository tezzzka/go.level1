package main

import (
	"fmt"
	"os"
	"strconv"
)

var box = map[int]int{0: 0, 1: 1, 2: 1}

func main() {
	defer fmt.Println(box)

	fmt.Print(fiRecursive(20))
}

func fiRecursive(N int) int {
	if N < 0 {
		return 0
	}
	if N < 2 {
		return N
	}
	if box[N-1] > 0 && box[N-2] > 0 {
		box[N] = box[N-1] + box[N-2]
		return box[N]
	} else {
		return fiRecursive(N-1) + fiRecursive(N-2)
	}
}

func fi(N uint64) uint64 {
	if N == 0 {
		return 0
	} else if N == 1 {
		return 1
	} else {
		return fi(N-1) + fi(N-2)
	}
}
func fiWrapper() {
	var N string
	ExitCode := []string{"q", "Q", "Quit", "QUIT"}
	mmap := map[uint64]uint64{}
	for {
		fmt.Fscan(os.Stdin, &N)
		for _, code := range ExitCode {
			if N == code {
				return
			} else if N == "show" {
				fmt.Println(mmap)
				return
			} else {
				continue
			}
		}
		n, e := strconv.ParseUint(string(N), 10, 64)
		if e != nil {
			panic("I cant convert the types. Sorry.")
		}
		mmap[n] = fi(n)
		fmt.Println(fi(n))
	}
}
