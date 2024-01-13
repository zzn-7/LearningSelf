package main

import (
	"GoStudy/user"
	"fmt"
)

var (
	name  string = "zzn"
	id    int    = 111
	clear int    = 112
)

func main() {
	s := user.Hello()
	fmt.Printf("s: %v\n", s)
	fmt.Printf("id: %v\n", id)
	fmt.Printf("clear: %v\n", clear)
	fmt.Printf("name: %v\n", name)
}
