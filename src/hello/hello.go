package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	fmt.Println(breakString("Hello World!"))
}

func breakString(inString string) string {
	for i := 0; i < len(inString); i++ {
		fmt.Println(string(inString[i]))
	}
	return "Successfully printed characters"
}
