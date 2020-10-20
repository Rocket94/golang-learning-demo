package main

import (
	"fmt"
	"net"
)

func main(){
	//监听
	listener,err:=net.Listen("tcp","127.0.0.1:8000")
	if err!=nil{
		fmt.Println("err =",err)
		return
	}
	defer listener.Close()

	for{
		conn,err:=listener.Accept()
		defer conn.Close()
		if err!=nil{
			fmt.Println("err =",err)
			continue
		}
		buf:=make([]byte,1024)
		n,err1:=conn.Read(buf)
		if err1!=nil{
			fmt.Println("err1 =",err1)
			continue
		}
		fmt.Println(string(buf[:n]))
		fmt.Println("----------------------------------")

	}
}
