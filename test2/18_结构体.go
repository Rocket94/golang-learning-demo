package main

import "fmt"

//定义,值传递
type Student struct {
	id int
	name string
	sex byte
	age int
	addr string
}
func main(){
	//初始化
	var s1 Student =Student{1,"peng da mo",'m',26,"Tianjin"}
	fmt.Println("s1 =",s1)
	//指针
	var p1 *Student=&s1
	fmt.Println("*p1 =",*p1)

	var s2 Student
	var p2=&s2//or p2=new(Student)
	p2.id=2
	p2.name="wangzi"
	p2.sex='f'
	p2.addr="Taian"
	p2.age=24
	fmt.Println("*p2 =",*p2)

	test01(s2)
	fmt.Println("值传递结果：s2 =",s2)
	test02(p2)
	fmt.Println("引用传递结果：s2 =",s2)
}
func test01(s Student) {
	s.name="socks"
}
func test02(p *Student){
	p.name="wawazi"
}

