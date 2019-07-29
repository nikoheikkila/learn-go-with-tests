package main

import "fmt"

const englishPrefix = "Hello"
const frenchPrefix = "Bonjour"
const spanishPrefix = "Hola"

func getPrefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix = frenchPrefix
	case "Spanish":
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}

	return
}

func hello(name string, language string) string {
	prefix := getPrefix(language)

	if name == "" {
		name = "world"
	}

	return fmt.Sprintf("%s, %s!", prefix, name)
}

func main() {
	message := hello("world", "Spanish")

	fmt.Println(message)
}
