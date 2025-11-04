package storage

import (
	"fmt"
	"sort"
)

// Implémentation mémoire
type MemoryStore struct {
	contacts map[int]*Contact
	nextID   int
}

// NewMemoryStore crée un nouveau store vide
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		contacts: make(map[int]*Contact),
		nextID:   1,
	}
}

// Add ajoute un contact avec un ID automatique
func (m *MemoryStore) Add(c *Contact) error {
	c.ID = m.nextID
	m.contacts[m.nextID] = c
	m.nextID++
	return nil
}

// Update met à jour un contact existant
func (m *MemoryStore) Update(id int, nom, email string) error {
	contact, ok := m.contacts[id]
	if !ok {
		return fmt.Errorf("aucun contact avec l'ID %d", id)
	}
	return contact.Update(nom, email)
}

// Delete supprime un contact par ID
func (m *MemoryStore) Delete(id int) error {
	if _, ok := m.contacts[id]; !ok {
		return fmt.Errorf("aucun contact avec l'ID %d", id)
	}
	delete(m.contacts, id)
	return nil
}

// List renvoie la liste triée des contacts
func (m *MemoryStore) List() []*Contact {
	var list []*Contact
	var ids []int

	for id := range m.contacts {
		ids = append(ids, id)
	}
	sort.Ints(ids)

	for _, id := range ids {
		list = append(list, m.contacts[id])
	}
	return list
}

// Get renvoie un contact par son ID
func (m *MemoryStore) Get(id int) (*Contact, error) {
	if c, ok := m.contacts[id]; ok {
		return c, nil
	}
	return nil, fmt.Errorf("aucun contact avec l'ID %d", id)
}
