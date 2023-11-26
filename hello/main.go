package main

import "fmt"

const prefixEnglish = "Hello, "
const prefixSpanish = "Hola, "
const prefixFrench = "Bonjour, "

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "spanish":
		prefix = prefixSpanish
	case "french":
		prefix = prefixFrench
	default:
		prefix = prefixEnglish
	}

	return
}

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
