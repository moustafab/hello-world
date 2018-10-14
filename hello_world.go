package main

import "fmt"

func main() {
	fmt.Println("What is your name?")
	var name string
	fmt.Scanln(&name)
	output := fmt.Sprintf("Hello, world. This is golang! Welcome %v!", name)
	fmt.Println(output)
}
