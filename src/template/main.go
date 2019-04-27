package main

import (
	"fmt"
	"html/template" // also "text/template" both are quite similar
	"os"
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
	err = t.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Print(err)
	}
}
