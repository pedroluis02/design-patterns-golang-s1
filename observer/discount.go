package observer

import "fmt"

func NewProductDiscount(product *TechProduct, discount float64) *ProductDiscount {
	p := &ProductDiscount{
		product: product,
		value:   discount,
		total:   0,
	}

	p.product.registerObserver(p)

	return p
}

type ProductDiscount struct {
	product *TechProduct
	value   float64
	total   float64
}

func (s *ProductDiscount) GetTotal() float64 {
	return s.total
}

func (s *ProductDiscount) Update(subject *Subject) {
	price := s.product.GetPrice()
	s.total = (price * s.value) / 100

	fmt.Println(fmt.Sprintf("price=%.2f discount=%.2f", price, s.total))
}
