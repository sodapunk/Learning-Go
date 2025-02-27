package main

import (
	"fmt"
)

func main() {
	var w1, w2, w3, w4, w5, w6 string
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	n, _ := fmt.Scanln(&w1, &w2, &w3, &w4, &w5, &w6)
	fmt.Println("Hello, World.")
	if n != 0 {
		fmt.Printf("%s %s %s %s %s %s", w1, w2, w3, w4, w5, w6)
	}

}
