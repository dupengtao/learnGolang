package main

import (
	"fmt"
)

func main() {
	//if
	if i := 10; i > 1 {
		fmt.Println("success")
	}else {
		fmt.Println("error")
	}// i 的作用域为if语句块

	i := 2
	if i > 1 {
		fmt.Println("ok")
	}

	if a, b, c := 1, 2, 3; a+b+c > 3 {
		fmt.Println("ok")
	}
	//for
	//第一种
	j := 0
	for {//无线循环
		j++
		if j > 10 {
			break
		}
		fmt.Println(j)
	}
	//第二种
	k := 5
	for k < 10 {//有条件的循环
		k++
	}
	fmt.Println(k)
	//第三种
	s := "dupengtao"
	l := len(s)
	for i := 0; i < l; i++ {
		fmt.Println(i)
	}
	//switch
	//第一种
	a := 0
	switch a{
	case 0:
		fmt.Println(a)
	case 1:
		fmt.Println(a)
	default:
		fmt.Println("default")
	}
	//第二种
	a=1
	switch {
	case a>=0:
		fmt.Println("a>=0")
		fallthrough//继续下面的case判断
	case a>=1:
		fmt.Println("a>=1")
	default :
		fmt.Println("default")
	}
	//第三种
	switch a:=1;{
	case a>=0:
		fmt.Println("a>=0")
		fallthrough
	case a>=1:
		fmt.Println("a>=1")
	default:
		fmt.Println("default")
	}

LABEL:
	for i:=0;i<10;i++{
		for{
			fmt.Println(i)
			continue LABEL
			//break
			//goto
		}
	}
	//break 中断执行
	//continue 继续执行，i的数量一直增长
	//goto 重新执行
}
