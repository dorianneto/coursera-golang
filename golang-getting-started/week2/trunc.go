package main

import "fmt"

func main() {
	var input float32

	fmt.Println("Type a float number:")
	fmt.Scan(&input)

	fmt.Println(int(input))
}
