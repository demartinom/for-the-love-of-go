// Package for doing stuff
package hello

// List of languages and their greetings
const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

// Says hello
func Hello(name string, language string) string {
	if name == "" { //In case no name is provded
		name = "World"
	}

	return greetingPrefix(language) + name
}

// Helper function to decide greeting language
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
