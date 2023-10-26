package pattern

import "fmt"

//предоставляет общий интерфейс для различных состояний.
type State interface {
	publish()
}

type Document struct {
	state State
}

//устанавливет документу другое состояние
func (p *Document) setState(st State) {
	p.state = st
}

type Draft struct {
	doc *Document
}

// отправляет на модерацию
func (d *Draft) publish() {
	fmt.Printf("Отправлено на модерацию.\n")
	d.doc.setState(&Moderation{d.doc})
}

type Moderation struct {
	doc *Document
}

// реализует модерацию
func (m *Moderation) publish() {
	fmt.Printf("Произошла модерация.\n")
	m.doc.setState(&Published{m.doc})
}

type Published struct {
	doc *Document
}

// публикует
func (p *Published) publish() {
	fmt.Printf("Опубликовано.\n")
}

func StateFunc() {
	document := Document{}
	document.setState(&Draft{&document})

	document.state.publish()
	document.state.publish()
	document.state.publish()
}
