package main

import (
	"fmt"
	"github.com/arielcr/factory-method-go/factory"
)

func main() {

	productA := factory.NewProduct(factory.ProductAType)
	fmt.Println(productA.MakeSomething("Awesome"))

	productB := factory.NewProduct(factory.ProductBType)
	fmt.Println(productB.MakeSomething("Nice"))

}
