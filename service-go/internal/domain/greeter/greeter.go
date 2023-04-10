package greeter

import (
    "fmt"
)

const (
    greetingFormatter = "Hello, %s"
)

func Greet(name string) string {
	if name == "" {
		name = "stranger"
	}

    greeting := fmt.Sprintf(greetingFormatter, name)
	return greeting
}
