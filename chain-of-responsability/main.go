package main

import "fmt"

func main() {

	handlerOne := &HandlerOne{}

	handlerTwo := &HandlerTwo{}

	handlerThree := &HandlerThree{}

	handlerOne.setNext(handlerTwo)
	handlerTwo.setNext(handlerThree)

	input := &Input{description: "test COR"}

	handlerOne.execute(input)
}

// COR

// tipo do parametro em comum
type Input struct {
	description string
}

// COR base
type Handler interface {
	execute(*Input)
	setNext(Handler)
}

// handler 1 COR

type HandlerOne struct {
	next Handler
}

func (handler *HandlerOne) execute(input *Input) {

	fmt.Println("Executando Handler 1")
	handler.next.execute(input)
}

func (handler *HandlerOne) setNext(next Handler) {
	handler.next = next
}

// handler 2 COR

type HandlerTwo struct {
	next Handler
}

func (handler *HandlerTwo) execute(input *Input) {

	fmt.Println("Executando Handler 2")
	handler.next.execute(input)
}

func (handler *HandlerTwo) setNext(next Handler) {
	handler.next = next
}

// handler 3 COR

type HandlerThree struct {
	next Handler
}

func (handler *HandlerThree) execute(input *Input) {

	fmt.Println("Executando Handler 3")
}

func (handler *HandlerThree) setNext(next Handler) {
	handler.next = next
}
