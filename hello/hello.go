package main

import "fmt"

const englishHelloPrefix = "Hello,"

func hello(name string) string {
	if name == "" {
		name = "world"
	}

	return fmt.Sprintf("%s %s!", englishHelloPrefix, name)
}

func main() {
	message := hello("world")

	fmt.Println(message)
}
