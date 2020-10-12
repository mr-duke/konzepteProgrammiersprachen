package main

import "fmt"

func swap(x, y int) (int, int) {
	return y, x
}

func swap2 (x *int, y *int){
	*x, *y = *y, *x
}


func main() {
	a,b := 1, 2
	c,d := 3, 4
	
	a, b = swap(a, b)
	swap2(&c,&d)
	fmt.Printf("Swap 1: a=%d, b=%d\n", a, b)
	fmt.Printf("Swap 2: c=%d, d=%d\n", c, d)
}

