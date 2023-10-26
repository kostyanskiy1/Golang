package pattern

import "fmt"

// Visitor интерфейс.
type Visitor interface {
	VisitSushiBar(p *SushiBar) string
	VisitPizzeria(p *Pizzeria) string
	VisitBurgerBar(p *BurgerBar) string
}

// предоставляет интерфейс для места, которое посетитель должен посетить.
type Place interface {
	Accept(v Visitor) string
}

// People удовлетворяет Visitor интерфейс.
type People struct {
}

// SushiBar.
func (v *People) VisitSushiBar(p *SushiBar) string {
	return p.BuySushi()
}

// Pizzeria.
func (v *People) VisitPizzeria(p *Pizzeria) string {
	return p.BuyPizza()
}

// BurgerBar.
func (v *People) VisitBurgerBar(p *BurgerBar) string {
	return p.BuyBurger()
}

// City реализует коллекцию мест для посещения.
type City struct {
	places []Place
}

// Add добавляет место к коллекции.
func (c *City) Add(p Place) {
	c.places = append(c.places, p)
}

// Accept реализует посещение всех мест в городе.
func (c *City) Accept(v Visitor) string {
	var result string
	for _, p := range c.places {
		result += p.Accept(v)
	}
	return result
}

// SushiBar реализует Place интерфейс.
type SushiBar struct {
}

// Accept реализация.
func (s *SushiBar) Accept(v Visitor) string {
	return v.VisitSushiBar(s)
}

// BuySushi реализация.
func (s *SushiBar) BuySushi() string {
	return "Buy sushi..."
}

// Pizzeria реализует Place интерфейс.
type Pizzeria struct {
}

// Accept реализация.
func (p *Pizzeria) Accept(v Visitor) string {
	return v.VisitPizzeria(p)
}

// BuyPizza реализация.
func (p *Pizzeria) BuyPizza() string {
	return "Buy pizza..."
}

// BurgerBar реализует Place интерфейс.
type BurgerBar struct {
}

// Accept реализация.
func (b *BurgerBar) Accept(v Visitor) string {
	return v.VisitBurgerBar(b)
}

// BuyBurger реализация.
func (b *BurgerBar) BuyBurger() string {
	return "Buy burger..."
}

func VisitorFunc() {

	city := new(City)

	city.Add(&SushiBar{}) //добавляем места для посещения
	city.Add(&Pizzeria{})
	city.Add(&BurgerBar{})

	result := city.Accept(&People{}) //посещение всех мест в городе.
	fmt.Println(result)
}
