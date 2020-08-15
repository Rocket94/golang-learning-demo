package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println("Contains测试：",strings.Contains("wangwazi","waz"))
	fmt.Println("Join测试：",strings.Join([]string{"peng","rong","kun"},"dashuaige"))
	fmt.Println("Index测试：",strings.Index("pengrongkun","k"))
	fmt.Println("Repeat测试：",strings.Repeat("cao",5))
	fmt.Println("Split测试：",strings.Split("ni:ge:s:b",":"))
}
