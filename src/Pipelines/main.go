package main

import (
	"html/template"
	"os"
)
// This App does :
// Create a template to get the product
// and then execute that template

const tax = 6.75 / 100

// Product struct has two fields
// Name and Price
type Product struct {
	Name string
	Price float32
}

// One method attaches to Product struct
// Multiple the price that going to find into the App
func (p Product) PriceWithTax() float32 {
	return p.Price * (1+ tax)
}


// Template
// no reason to have a `string` into that syntax {{ }}
// just for test
// `-` Skip all write space after or before
const templateString  = `
{{- "Item information" }}
Name: {{ .Name }}
Price: {{ printf "$%.2f" .Price }}
Price with Tax: {{ .PriceWithTax | printf "$%.2f" }}
`

// Find a product and his price with
func main()  {
	p := Product{
		Name: "Name of product",
		Price: 2.18,
	}

	// templateString form the content of the template
	t:= template.Must(template.New("").Parse(templateString))
	t.Execute(os.Stdout, p)
}