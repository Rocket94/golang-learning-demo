package main

import "fmt"

type Person struct {
	name string
	age int
	sex byte
}
type Student struct {
	Person
	id int
	major string
}
func(p *Person)Printp(){
	if p.sex=='m'{
		fmt.Printf("这位帅哥姓名是%s,今年%d岁\n",p.name,p.age)
	}else{
		fmt.Printf("这位美女芳名是%s,今年%d岁\n",p.name,p.age)
	}
}
//重写
func(p *Student)Printp(){
	if p.sex=='m'{
		fmt.Printf("这位帅哥姓名是%s,今年%d岁,还是%s专业的学生呢\n",p.name,p.age,p.major)
	}else{
		fmt.Printf("这位美女芳名是%s,今年%d岁,还是%s专业的学生呢\n",p.name,p.age,p.major)
	}
}
func main(){
	a:=Person{
		name: "彭荣坤",
		age:  26,
		sex:  'm',
	}
	a.Printp()
	b:=Student{
		Person: Person{
			name: "王慧咏",
			age:  24,
			sex:  'f',
		},
		id:     1024,
		major:  "测控技术与仪器",
	}
	//先调用自己作用域的，找不到再去继承域上找
	b.Printp()
	b.Person.Printp()
	//方法能赋值给变量
	ttFunc:=b.Printp
	ttFunc()
	//方法表达式
	ppFunc:=(*Person).Printp
	ppFunc(&a)
}
