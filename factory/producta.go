package factory

type ProductA struct {
	Product
}

func NewProductA() ProductInterface {
	return &ProductA{
		Product: Product{
			name: "Product A",
		},
	}
}