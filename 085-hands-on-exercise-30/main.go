package main

import (
	"fmt"

)

func main(){
	xi := []int{42, 43, 44, 45, 46, 47}
	for index, value := range xi{
	fmt.Printf("%v %v \n", index, value)
	}

	
	


}