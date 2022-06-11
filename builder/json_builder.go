package builder

import "encoding/json"

type ProductJsonBuilder struct {
	product *Product
}

func (b *ProductJsonBuilder) SetData(product *Product) ProductBuilder {
	b.product = product
	return b
}

func (b *ProductJsonBuilder) Build() string {
	data, _ := json.Marshal(b.product)

	return string(data)
}
