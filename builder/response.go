package builder

type ProductResponse struct {
	builder ProductBuilder
}

func (r *ProductResponse) SetBuilder(builder ProductBuilder) *ProductResponse {
	r.builder = builder
	return r
}

func (r *ProductResponse) BuildResponse(product *Product) string {
	return r.builder.SetData(product).Build()
}
