package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name, language string) string {
	retMsg := getPrefix(language)

	if name == "" {
		name = "world"
	}

	retMsg += name
	return retMsg
}

func getPrefix(language string) (prefix string) {
	switch language {
	case "spanish":
		prefix += spanishHelloPrefix
	case "french":
		prefix += frenchHelloPrefix
	default:
		prefix += englishHelloPrefix
	}
	return
}

func main() {
	name := "Joe"
	fmt.Println(Hello(name, "english"))
}
