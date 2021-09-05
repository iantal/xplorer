package domain

// Library represents a third party library from the dependency tree
type Library struct {
	Name  string `json:"name,omitempty"`
	Level int    `json:"level,omitempty"`
}

// NewLibrary creates a Library
func NewLibrary(name string, level int) *Library {
	return &Library{
		Name:  name,
		Level: level,
	}
}
