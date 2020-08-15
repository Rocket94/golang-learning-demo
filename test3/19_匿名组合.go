package main

import "fmt"

type mystr string
type Person struct {
	name string
	sex  byte
	age int
}
type Student struct {
	Person//只有类型，没有名字，匿名字段，继承了Person成员 
	id int
	addr string
}
type Student2 struct {
	*Person//只有类型，没有名字，匿名字段，继承了Person成员
	id int
	name string//同名字段，会操作该作用域的字段，而不是父字段
	int//基础类型匿名字段
	mystr
}
func main(){
	//init
	var s1 Student=Student{
		Person: Person{"彭荣坤",'m',26},
		id:     0,
		addr:   "天津市",
	}
	fmt.Println("s1 =",s1)
	fmt.Printf("s1 =%+v\n",s1)

	s2:=Student{Person:Person{name:"王慧咏"},addr:"泰安市"}
	fmt.Println("s2 =",s2)
	fmt.Println("-----------")
	s2.age=26
	s2.sex='f'
	s2.id=1
	fmt.Println("s2 =",s2)
	
	s3:=Student2{
		Person: new(Person),
		id:     3,
		name:   "豆豆",
		int:    666,
		mystr:  "是个小胖子",
	}
	fmt.Println("s3 =",s3)
}