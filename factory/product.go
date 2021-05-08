package factory

type Product struct {
	name string
}

func (p *Product) MakeSomething(feature string) string {
	return feature + " " + p.name
}