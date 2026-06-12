package main

import "fmt"

func main() {
	x := 10
	p := &x

	fmt.Println(x)
	fmt.Println(*p)

	*p = 10
	fmt.Println(x)

}
