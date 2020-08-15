package main

import "fmt"

type Humaner interface {
	sayhi()
}
//多态
func whoSayHi(i Humaner){
	i.sayhi()
}
type Student struct {
	name string
	id int
}
func (tmp Student)sayhi(){
	fmt.Printf("Student[%s,%d] say hi\n",tmp.name,tmp.id)
}

type Teacher struct {
	 	addr string
	 	group string
}
func (tmp *Teacher)sayhi(){
	fmt.Printf("Teacher[%s,%s] say hi\n",tmp.addr,tmp.group)
}

type MyStr string
func (tmp *MyStr)sayhi(){
	fmt.Printf("MyStr[%s] say hi\n",*tmp)
}
func main(){
	var i Humaner
	i=Student{
		name: "彭荣坤",
		id:   0521,
	}
	i.sayhi()
	whoSayHi(i)
	fmt.Println("---------------")
	i=&Teacher{
		addr:  "东北",
		group: "二人转组",
	}
	i.sayhi()
	whoSayHi(i)
	fmt.Println("---------------")
	var a MyStr= "1314"
	i= &a
	i.sayhi()
	whoSayHi(i)
	fmt.Println("---------------")


}
