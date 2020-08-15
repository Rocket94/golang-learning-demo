package main

import "fmt"

func main(){
	var a int
	//每个变量有两个含义，内存和地址
	fmt.Printf("a=%d\n",a)//内存
	fmt.Printf("&a=%v\n",&a)//地址，也叫指针

	//&是取变量地址(指针)，*是同过指针访问目标对象(被指对象)
	var p *int
	fmt.Println(p)
	p=&a
	fmt.Printf("p=%v,&a=%v\n",p,&a)
	*p=666
	*&a=1
	fmt.Printf("p=%v,&a=%v\n，a=%d\n",p,&a,a)

	fmt.Println("------------")
	var q *int
	q=new(int)
	*q=123
	fmt.Printf("*q= %v\n",*q)
}
