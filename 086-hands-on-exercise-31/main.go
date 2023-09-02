package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	age := m["James"]
	fmt.Println(age)
	
	if _, ok := m["Q"]; ok {
		fmt.Println("Q is in the m")
		} else {
		fmt.Println("Q is not in m")
	}

}
