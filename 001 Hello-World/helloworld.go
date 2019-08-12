package main

import "fmt"

func main() {
	n, err := fmt.Println("Hello World!", 42, true)
	fmt.Println(n, "bytes")
	fmt.Println(err, "error")
}
