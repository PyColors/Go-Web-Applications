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

// Template
// no reason to have a `string` into that syntax {{ }}
// just for test
// `-` Skip all write space after or before
const templateString  = `
{{- "Item information" }}
Name: {{ .Name }}
Price: {{ printf "$%.2f" .Price }}
Price with Tax: {{ calctax .Price | printf "$%.2f" }}
`

// Find a product and his price with
func main()  {
	p := Product{
		Name: "Name of product",
		Price: 2.18,
	}

	// New function map
	fm := template.FuncMap{}
	// new key to invoke the function into the template
	fm["calctax"] = func(price float32) float32 {
		return price * (1 + tax)
	}
	// templateString form the content of the template
	t:= template.Must(template.New("").Funcs(fm).Parse(templateString))
	t.Execute(os.Stdout, p)
}