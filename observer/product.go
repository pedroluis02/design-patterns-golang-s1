package observer

func NewProduct(name string) *TechProduct {
	p := &TechProduct{
		name:  name,
		price: 0,
	}

	p.Init()

	return p
}

type TechProduct struct {
	Subject
	name  string
	price float64
}

func (p *TechProduct) SetName(name string) {
	p.name = name
}

func (p *TechProduct) GetName() string {
	return p.name
}

func (p *TechProduct) SetPrice(price float64) {
	p.price = price

	p.Notify()
}

func (p *TechProduct) GetPrice() float64 {
	return p.price
}
