package main

import "fmt"

//1
func Pswap(a, b *int) {
	*a, *b = *b, *a
}

//2
type User struct {
	Age int
}

func birthday(u *User) {
	u.Age++
}

//3
func isNil(ptr *int) {
	if ptr == nil {
		fmt.Print("null ptr")
	} else {
		fmt.Print(*ptr)
	}
}
