package main

import (
	"fmt"
)

type Person struct {
	name string
	sex byte
	age int
}
//面向对象，先给类型绑定一个函数
type para int
//i是接收者，接收者本质是传递的一个参数
func (i para)Add(j para)para{
	return i+j
}
func (tmp Person)PrintP() {
	fmt.Println(tmp)
}
//接收类型本身不能是指针类型
func (p *Person)initPerson(name string,sex byte,age int){
	p.name=name
	p.age=age
	p.sex=sex
}
func main(){
	//定义变量
	var a para=2
	//调用方法：变量名.函数（所需函数）
	result:=a.Add(3)
	fmt.Println(result)

	p:=Person{
		name: "小袜子",
		sex:  'f',
		age:  1,
	}
	p.PrintP()
	var b Person
	(&b).initPerson("王慧咏",'f',24)
	b.PrintP()
}
