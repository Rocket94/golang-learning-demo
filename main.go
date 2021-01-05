package main

import (
	"fmt"
	"reflect"
)

func main() {
	a:=[]int{1,3,2}
	b:=[]int{1,3,2}
	fmt.Println(reflect.DeepEqual(a,b))
}
