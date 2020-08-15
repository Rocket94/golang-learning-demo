package main

import "fmt"

func main(){
	//条件
	if 4>3{
		fmt.Println("既然你诚心诚意地发问了")
	}
	if a:=10;a==10{
		fmt.Println("我就大发慈悲地告诉你")
	}

	var num int
	goto End
	fmt.Scanf("%d\n",&num)
	switch num{
	case 1:
		fmt.Println("武藏")
		//break//go语言保留了break关键字,不用写，默认包含
		fallthrough//后面无条件执行
	case 2:
		fmt.Println("小次郎")
		fallthrough
	case 3:
		fmt.Println("喵喵")
	}
	//循环
	for i:=0;i<10;i++{
		fmt.Print(i)
	}
	End:
	fmt.Println()

	//range迭代
	str:="abc"
	for i,data:=range str{
		fmt.Printf("str[%d]=%c\n",i,data)
	}
	for i,_:=range str{
		fmt.Printf("str[%d]=%c\n",i,str[i])
	}

}
