package main

import (
	"fmt"
)

type Student struct {
	name string
	id int
}
func main(){
	i:=make([]interface{},3)
	i[0]=1
	i[1]="fuck you"
	i[2]=Student{
		name: "shabi",
		id:   1,
	}
	for _,data:=range i{
		if value,ok:=data.(int);ok==true{
			fmt.Println("存储类型为Int，内容为",value)
		}else if value,ok:=data.(string);ok==true{
			fmt.Println("存储类型为String，内容为",value)
		}else if value,ok:=data.(Student);ok==true{
			fmt.Println("存储类型为Student，内容为",value)
		}
	}

}
