package storage

// Storer décrit les opérations de stockage possibles
type Storer interface {
	Add(c *Contact) error
	Update(id int, nom, email string) error
	Delete(id int) error
	List() []*Contact
	Get(id int) (*Contact, error)
}

