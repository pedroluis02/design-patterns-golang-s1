package observer

import "fmt"

func NewProductTax(product *TechProduct) *ProductTax {
	p := &ProductTax{
		product: product,
		value:   18,
		total:   0,
	}

	p.product.RegisterObserver(p)

	return p
}

type ProductTax struct {
	product *TechProduct
	value   float64
	total   float64
}

func (s *ProductTax) GetTotal() float64 {
	return s.total
}

func (s *ProductTax) Update(subject *Subject) {
	price := s.product.GetPrice()
	s.total = (price * s.value) / 100

	fmt.Println(fmt.Sprintf("price=%.2f tax=%.2f", price, s.total))
}
