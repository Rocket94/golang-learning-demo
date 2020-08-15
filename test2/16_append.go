package main

import "fmt"

func main(){
	s1:=[]int{}
	s1 = append(s1, 1)
	fmt.Printf("len = %d,cap = %d\n",len(s1),cap (s1))
	s1 = append(s1, 2)
	fmt.Printf("len = %d,cap = %d\n",len(s1),cap (s1))
	s1 = append(s1, 3)

	s2:=[]int{8,8,8,8,8}
	copy(s2,s1)
	fmt.Printf("s2 = %v",s2)
}
