package factory

type ProductType int

const (
	ProductAType ProductType = 1 << iota
	ProductBType
)

func NewProduct(t ProductType) ProductInterface {
	switch t {
	case ProductAType:
		return NewProductA()
	case ProductBType:
		return NewProductB()
	default:
		return NewProductA()
	}
}
