package builder

import "encoding/xml"

type ProductXmlBuilder struct {
	product *Product
}

func (b *ProductXmlBuilder) SetData(product *Product) ProductBuilder {
	b.product = product
	return b
}

func (b *ProductXmlBuilder) Build() string {
	data, _ := xml.Marshal(b.product)

	return string(data)
}
