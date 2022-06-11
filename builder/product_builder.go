package builder

type ProductBuilder interface {
	SetData(product *Product) ProductBuilder
	Build() string
}
