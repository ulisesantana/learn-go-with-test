package hello_world

import "fmt"

const (
	english = "English"
	french  = "French"

	canarianHelloPattern = "Oh! Qué pasó %s?!"
	englishHelloPattern  = "Hello, %s!"
	frenchHelloPattern   = "Bonjour, %s!"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "mi niño"
	}
	return fmt.Sprintf(greetingPrefix(language), name)
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPattern
	case english:
		prefix = englishHelloPattern
	default:
		prefix = canarianHelloPattern
	}
	return
}

func main() {
	fmt.Println(Hello("", ""))
}
