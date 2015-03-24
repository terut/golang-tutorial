package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // pointer
	fmt.Println(*p) // value
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
