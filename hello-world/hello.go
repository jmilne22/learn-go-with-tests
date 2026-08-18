package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"
	hebrew  = "Hebrew"

	englisHelloPrefix  = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	hebrewHelloPrefix  = "שלום, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case hebrew:
		prefix = hebrewHelloPrefix
	default:
		prefix = englisHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
