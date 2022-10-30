package main

import (
	"fmt"
)

func main() {
	a := [3]int{1, 2, 3}
	b := a[1:3]
	b[0] = 333
	tmp := append(b, 4)
	tmp[0] = 444
	fmt.Println(tmp)
	fmt.Println(a)
}
