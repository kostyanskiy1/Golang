package pattern

import "fmt"

// Transport - продукт
type Transport interface {
	move()
	getName() string
}

// ------------------ CAR ------------------

// CarFromFactory - concreteProduct
type CarFromFactory struct {
	name string
}

func NewCarFromFactory() Transport {
	return &CarFromFactory{"car"}
}

func (c *CarFromFactory) getName() string {
	return c.name
}

func (c *CarFromFactory) move() {
	fmt.Printf("%v is moving...\n", c.getName())
}

// ------------------ BOAT ------------------

// BoatFromFactory - concreteProduct
type BoatFromFactory struct {
	name string
}

func NewBoatFromFactory() Transport {
	return &BoatFromFactory{"boat"}
}

func (b *BoatFromFactory) getName() string {
	return b.name
}

func (b *BoatFromFactory) move() {
	fmt.Printf("%v is moving...\n", b.getName())
}

// -------------------------------------------

type Factory interface {
	factoryMethod()
}

// factoryMethod - сборка лодки или машины
func factoryMethod(str string) Transport {
	switch str {
	case "car":
		return NewCarFromFactory()
	case "boat":
		return NewBoatFromFactory()
	default:
		return nil
	}
}

func FactoryMethodFunc() {
	driver := factoryMethod("car")
	driver.move()

	driver = factoryMethod("boat")
	driver.move()
}
