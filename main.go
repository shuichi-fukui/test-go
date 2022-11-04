package main

import "fmt"

func main() {
	msg := test()
	fmt.Println("msg:", msg)
}

func test() string {
	return "test!"
}
