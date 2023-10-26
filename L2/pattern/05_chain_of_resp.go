package pattern

import "fmt"

// Handler интерфейс.
type Handler interface {
	SendRequest(message int) string
}

// обработчик A;
type ConcreteHandlerA struct {
	next Handler
}

// отправка запроса.
func (h *ConcreteHandlerA) SendRequest(message int) (result string) {
	if message == 1 {
		result = "Im handler 1"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

// обработчик Б;
type ConcreteHandlerB struct {
	next Handler
}

// отправка запроса.
func (h *ConcreteHandlerB) SendRequest(message int) (result string) {
	if message == 2 {
		result = "Im handler 2"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

// обработчик С;
type ConcreteHandlerC struct {
	next Handler
}

// отправка запроса.
func (h *ConcreteHandlerC) SendRequest(message int) (result string) {
	if message == 3 {
		result = "Im handler 3"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

func ChainOfResponsabilityFunc() {
	handlers := &ConcreteHandlerA{
		next: &ConcreteHandlerB{
			next: &ConcreteHandlerC{},
		},
	}

	result := handlers.SendRequest(2)
	fmt.Println(result)
}
