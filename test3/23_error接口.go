package main

import (
	"errors"
	"fmt"
)

func MyDiv(a,b int) (result int,err error) {
	if(b==0){
		err=errors.New("分母不能为0")
	}else {
		result=a/b
	}
	return
}

func main(){
	err1:=fmt.Errorf("%s","this is normol err1")
	fmt.Println("err1 =",err1)

	err2:=errors.New("this is normol err2")
	fmt.Println("err2 =",err2)

	result,err:=MyDiv(10,0)
	if err==nil {
		fmt.Println("结果为：",result)
	}else {
		fmt.Println(err)
	}
}

