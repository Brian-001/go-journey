package main

import "fmt"

func main() {
	slice := make([]string, 3, 5)
	fmt.Println("Length:", len(slice))
	fmt.Println("Capacity:", cap(slice))
	fmt.Println(slice)
}