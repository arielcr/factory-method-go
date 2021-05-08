package factory

type ProductB struct {
	Product
}

func NewProductB() ProductInterface {
	return &ProductB{
		Product: Product{
			name: "Product B",
		},
	}
}