package main

import (
	"fmt"
)

// Message represents the greet message
type Message string

// Greeter must greet the others with a specific message
type Greeter struct {
	Message Message
}

// Event is the hook for domain, must execute the logic
type Event struct {
	Greeter Greeter
}

// NewMessage returns the message "Hello, world!"
func NewMessage() Message {
	return Message("Hello, world!")
}

// NewGreeter receives a Message that are used to instance a new Gretter, returns the Greeter
func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

// Greet is a method of Greeter, must return the Message
func (g Greeter) Greet() Message {
	return g.Message
}

// NewEvent receives a Greeter that are used to instance a new Event, returns the Event
func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

// Start is a method of Event, must execute the bussiness logic
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

// main is the entrypoint of application, will instance the application
func main() {
	message := NewMessage()
	greeter := NewGreeter(message)
	event := NewEvent(greeter)

	event.Start()
}
