package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

// 1
type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func SayHello(name string, greet Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", greet.LanguageName(), greet.Greet(name))
}

// 2
type Italian struct {
}

func (obj Italian) LanguageName() string {
	return "Italian"
}

func (obj Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

// 3
type Portuguese struct {
}

func (obj Portuguese) LanguageName() string {
	return "Portuguese"
}

func (obj Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}
