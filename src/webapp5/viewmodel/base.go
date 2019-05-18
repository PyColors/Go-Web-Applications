package viewmodel

// The basic model for all views
type Base struct {
	Tile string
}

// Initialisation
func NewBase() Base {
	// Constructor function
	return Base {
		Tile: "Default title",
	}

}
