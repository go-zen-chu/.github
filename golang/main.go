package main

import "fmt"

func hello() error {
	fmt.Println("hello")
	return nil
}

func main() {
	if err := hello(); err != nil {
		fmt.Println("Error:", err)
	}
}
