package pattern

import "fmt"

// Builder интерфейс.
type Builder interface {
	MakeHeader(str string)
	MakeBody(str string)
	MakeFooter(str string)
}

// Класс Director, который будет распоряжаться строителем и отдавать ему команды в нужном порядке, а строитель будет их выполнять;
type Director struct {
	builder Builder
}

// Construct говорит builder что делать и в каком порядке
func (d *Director) Construct() {
	d.builder.MakeHeader("Header")
	d.builder.MakeBody("Body")
	d.builder.MakeFooter("Footer")
}

// ConcreteBuilder который реализует интерфейс строителя и взаимодействует со сложным объектом;
type ConcreteBuilder struct {
	product *Product
}

// сделать заголовок документа
func (b *ConcreteBuilder) MakeHeader(str string) {
	b.product.Content += "<header>" + str + "</header>\n"
}

// сделать тело документа
func (b *ConcreteBuilder) MakeBody(str string) {
	b.product.Content += "<article>" + str + "</article>\n"
}

// сделать конец документа
func (b *ConcreteBuilder) MakeFooter(str string) {
	b.product.Content += "<footer>" + str + "</footer>\n"
}

type Product struct { //сложный объект
	Content string
}

//
func (p *Product) Show() string {
	return p.Content
}
func Build() {
	product := new(Product)

	director := Director{&ConcreteBuilder{product}}
	director.Construct()

	res := product.Show()
	fmt.Println(res)
}
