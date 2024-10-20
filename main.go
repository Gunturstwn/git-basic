package main

import "fmt"

func main() {
	fmt.Println("asasa")

	slice := []int{343, 34, 234, 5}

	for i, i2 := range slice {
		fmt.Printf("index : %d, values : %d\n", i, i2)
	}
}
