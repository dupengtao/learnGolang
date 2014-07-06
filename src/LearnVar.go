package main

import (
	"fmt"
)

func main(){
	//第一种
	var f float32
	var s string
	//第二种
	var (
		i int
		b bool
	)
	fmt.Println(i)
	fmt.Println(f)
	fmt.Println(s)
	fmt.Println(b)
	//第三种
	i2 :=2
	fmt.Println(i2)
	//第四种
	var i3 =6
	fmt.Println(i3)
	var w,_,y,z=1,2,3,4
	//var w,_,y,z int =1,2,3,4;
	fmt.Println(w)
	fmt.Println(y)
	fmt.Println(z)
	//类型转换
	var f2 float32=2.22
	i2=int(f2)
	fmt.Println(i2)
	i4:=int(f2)
	fmt.Println(i4)

	s2:=string(65)
	fmt.Println(s2)//A

	// i5:=int("A")
	// fmt.Println(i5)
}
