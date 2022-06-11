package builder

type Product struct {
	Id    string  `json:"id" xml:"id,attr"`
	Name  string  `json:"name" xml:"name,attr"`
	Price float32 `json:"price" xml:"price,attr"`
	Stock int     `json:"stock" xml:"stock,attr"`
}
