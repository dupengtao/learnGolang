package main

import (
	"fmt"
)

func main() {
	var array1 [3]int
	fmt.Println(array1)

	array2:= [...]int{1,2,3,4}
	fmt.Println(array2)

	array3:=[20]int{17:9,19:10}
	fmt.Println(array3)

	var array4 [4]*int
	fmt.Println(array4)

	p := new([20]int)
	fmt.Println(p)
	p[2]=3
	fmt.Println(p[2])

	array5 :=[...]int{5,23,0,-54,22}
	n:=len(array5)
	for i:=0;i<n;i++{
		for j:=i+1;j<n;j++{
			if array5[i]>array5[j]{
				var temp = array5[i]
				array5[i]=array5[j]
				array5[j]=temp
			}
		}
	}
	fmt.Println(array5)

	pArray:=&array5

	for i:=0;i<n;i++{
		for j:=i+1;j<n;j++{
			if(pArray[i]<pArray[j]){
				temp:=pArray[i]
				pArray[i]=pArray[j]
				pArray[j]=temp
			}
		}
	}
	fmt.Println(*pArray)
}
