package main

import (
	"sort"
	"fmt"
	"io"
	"os"
)

type array []int

func (a array) Len() int           { return len(a) }
func (a array) Less(i, j int) bool { return a[i] < a[j] }
func (a array) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	a := array{1, 4, 5, 6, 2, 4, 23, 1, 14, 5}
	sort.Sort(a)
	fmt.Println(a)
	os.Stdout.Write()
}
