package main

import "fmt"

const (
	spanishLang         = "Spanish"
	frenchLang          = "French"
	japaneseLang        = "Japanese"
	defaultName         = "World"
	helloPrefix         = "Hello, "
	spanishHelloPrefix  = "Hola, "
	frenchHelloPrefix   = "Bonjour, "
	japaneseHelloPrefix = "Konnichiwa, "
)

func Hello(name, language string) string {

	if name == "" {

		name = defaultName
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {

	case spanishLang:
		prefix = spanishHelloPrefix
	case frenchLang:
		prefix = frenchHelloPrefix
	case japaneseLang:
		prefix = japaneseHelloPrefix
	default:
		prefix = helloPrefix
	}

	return // Al asignar la variable de prefix desde la declaracion, no es necesario especificarlo en el return
}

func main() {
	fmt.Println(Hello("", ""))
}
