package helloworld

import "fmt"

// you can group constants like this 
const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func main() {
	fmt.Println(Hello("", "something"))
}

// public functions start with upper case
func Hello(name string, language string) string {

	if name == "" {
		return englishHelloPrefix + "World"
	}

	return greetingPrefix(language) + name

}

// the prefix is the return name / type of the function
// this creates a variable prefix thats usable in the function
// by default it is assigned the "Zero" value for that type
// you can return whatever it is set to by just saying `return`

// private functions start with lower case
func greetingPrefix(language string) (prefix string) {

	switch language {
	case "French":
		prefix = frenchHelloPrefix

	case "Spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	// why can there be an empty return here ?
	// because prefix is a named return value, so when you type empty return
	// the named return value, prefix, will be returned 
	return

}
