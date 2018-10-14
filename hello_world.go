package main

import "fmt"

func main() {
	fmt.Println("What is your name?")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Hello, world. This is golang! Welcome %v! \n", name)
}
