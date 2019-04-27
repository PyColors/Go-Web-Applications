package main

import (
	"fmt"
	"os"
	"text/template" // also "html/template" both are quite similar but not for security issues
)

func main() {
	// String going to form the basic template
	templateString := `template name`

	// Create the template
	t, err := template.New("title").Parse(templateString)

	// Handle the error if the template fails
	if err != nil {
		fmt.Print(err)
	}

	// Execute the template
	// He is actually another operation can be fail (the error value)
	// First parameter: `Stdout` allows us to see the print out in the console
	// Second parameter: `nil` - by the name `data interface {}`
	err = t.Execute(os.Stdout, nil)

	// if errors come in
	if err != nil {
		fmt.Print(err)
	}
}
