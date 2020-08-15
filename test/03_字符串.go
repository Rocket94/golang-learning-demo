package main

import "fmt"

func main(){
	var str1 string
	str1="abc"
	str2:="def"
	fmt.Println(str1,str2)
	fmt.Println(len(str2))
	fmt.Printf("str1=%c,str2=%c\n",str1[0],str2[1])
	//复数
	var t1,t2 complex128
	t1=2.1+3.14i
	t2=3.2+5.64i
	fmt.Println(t1+t2,real(t1),imag(t2))

	//格式化%v自动匹配类型
	a:=21
	b:="abc"
	c:='c'
	d:=3.14
	fmt.Printf("a=%v,b=%v,c=%v,d=%v\n",a,b,c,d)

	//键盘输入
	fmt.Printf("请输入a的值：")
	fmt.Scanf("%d",&a)
	fmt.Println(a)

	//类型别名
	type bigint int64
	var e bigint
	fmt.Printf("a类型为%T\n",e)
}