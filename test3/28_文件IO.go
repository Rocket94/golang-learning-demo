package main

import (
	"fmt"
	"io"
	"os"
)

func WriteFile(path string){
	f,err:=os.Create(path)
	if err !=nil{
		fmt.Println("err =",err)
		return
	}
	defer f.Close()
	var buf string
	for i:=0;i<10;i++{

		buf=fmt.Sprintf("i = %d\n",i)
		n,err:=f.WriteString(buf)
		if err!=nil{
			fmt.Println("err =",err)
			return
		}
		fmt.Println("n =", n)
	}
}
func ReadFile(path string){
	f,err1:=os.Open(path)
	defer f.Close()
	if err1!=nil{
		fmt.Println("err1 =",err1)
	}
	buf:=make([]byte,1024)
	s,err2:=f.Read(buf)
	if err2!=nil&&err2!=io.EOF{
		fmt.Println("err2 =",err2)
		return
	}
	fmt.Println(string(buf[:s]))
}
func main(){
	//os.Stdout.Close()//关闭无法输出，默认是打开的
	fmt.Println("are you ok?")
	var a string
	//os.Stdin.Close()
	//fmt.Scan(&a)
	if a=="ok"{
		fmt.Println("i'm very ok!")
		return
	}else {
		fmt.Println("why?")
	}
	fmt.Println("--------------")
	//WriteFile("./test.txt")
	ReadFile("./test.txt")
}