package pattern

/*
Можно использовать если:
нужно собирать сложные составные объекты
код должен создавать разные представления какого-то объекта
	Например, деревянные и железобетонные дома

+:
 Позволяет создавать продукты пошагово
 Позволяет использовать один и тот же код для создания различных продуктов
 Изолирует сложный код сборки продукта от его основной бизнес-логики

-:
Усложняет код программы из-за введения дополнительных классов
Клиент будет привязан к конкретным классам строителей
	так как в интерфейсе директора может не быть метода получения результата
*/
import "fmt"

type IChief interface {
	setBread()
	setMeat()
	setGreens()
	getBurger() burger
}

func getChief(chief string) IChief {
	if chief == "default" {
		return newDefaultChief()
	}

	if chief == "vegan" {
		return newVeganChief()
	}
	return nil
}

type burger struct {
	bread  string
	meat   string
	greens string
}

type veganChief struct {
	bread  string
	meat   string
	greens string
}

func newVeganChief() *veganChief {
	return &veganChief{}
}

func (v *veganChief) setBread() {
	v.bread = "Веганский хлеб"
}
func (v *veganChief) setMeat() {
	v.meat = "Соевое мясо"
}
func (v *veganChief) setGreens() {
	v.greens = "Много зелени"
}
func (v *veganChief) getBurger() burger {
	return burger{bread: v.bread,
		meat:   v.meat,
		greens: v.greens}
}

type defaultChief struct {
	bread  string
	meat   string
	greens string
}

func newDefaultChief() *defaultChief {
	return &defaultChief{}
}

func (d *defaultChief) setBread() {
	d.bread = "Белый хлеб"
}
func (d *defaultChief) setMeat() {
	d.meat = "Говяжья котлета"
}
func (d *defaultChief) setGreens() {
	d.greens = "Пара огурчиков"
}
func (d *defaultChief) getBurger() burger {
	return burger{bread: d.bread,
		meat:   d.meat,
		greens: d.greens}
}

type Director struct {
	chief IChief
}

func newChief(c IChief) *Director {
	return &Director{
		chief: c,
	}
}

func (d *Director) setChief(c IChief) {
	d.chief = c
}

func (d *Director) makeBurger() burger {
	d.chief.setBread()
	d.chief.setMeat()
	d.chief.setGreens()
	return d.chief.getBurger()
}
func main() {
	burgerShop := newChief(getChief("vegan"))
	veganBurger := burgerShop.makeBurger()
	fmt.Println(veganBurger)
	burgerShop.setChief(getChief("default"))
	fmt.Println(burgerShop.makeBurger())
}
