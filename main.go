package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fiWrapper()
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
