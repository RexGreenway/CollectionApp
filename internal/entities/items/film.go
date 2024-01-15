package items

import (
	"fmt"
)

// Film is an example of implementing the Item interface in order to create
// pre-defined 'collectables'.

const FILM ItemType = "Film"

type Film struct {
	ID   string
	Name string
	Year int
}

func (f Film) GetID() string {
	return f.ID
}

func (f Film) GetName() string {
	return f.Name
}

func (f Film) GetType() ItemType {
	return FILM
}

func (f Film) String() string {
	return fmt.Sprintf("Film: %s", f.Name)
}
