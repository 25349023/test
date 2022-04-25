package main

import "fmt"

func main() {
	fmt.Println(3 >= 2)
	if a := 2; a == 3 {
		fmt.Println("wow")
	} else {
		fmt.Println("hello")
	}

	for i := range []int{1, 2, 3} {
		fmt.Println(i)
	}
}
