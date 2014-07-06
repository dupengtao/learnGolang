package main

import (
	"fmt"
)

func main(){
	i :=1
	p :=&i
	var p1 *int = &i
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(*p1)
}
