package contact

import (
	"fmt"
	"sort"
)

type Contact struct {
	ID    int
	Nom   string
	Email string
}

// Ajouter un contact à la map
func AddContact(contacts map[int]Contact, c Contact) error {
	if _, exists := contacts[c.ID]; exists {
		return fmt.Errorf("un contact avec l'ID %d existe déjà", c.ID)
	}
	contacts[c.ID] = c
	return nil
}

// Lister tous les contacts
func ListContacts(contacts map[int]Contact) {
	if len(contacts) == 0 {
		fmt.Println("Aucun contact enregistré.")
		return
	}

	fmt.Println("Liste des contacts :")

	// 1. Extraire toutes les clés
	var ids []int
	for id := range contacts {
		ids = append(ids, id)
	}

	// 2. Trier les IDs
	sort.Ints(ids)

	// 3. Afficher dans l’ordre
	for _, id := range ids {
		c := contacts[id]
		fmt.Printf("ID: %d | Nom: %s | Email: %s\n", c.ID, c.Nom, c.Email)
	}
}

// Supprimer un contact
func DeleteContact(contacts map[int]Contact, id int) error {
	if _, ok := contacts[id]; !ok { // <-- comma ok idiom
		return fmt.Errorf("aucun contact avec l'ID %d", id)
	}
	delete(contacts, id)
	return nil
}

// Mettre à jour un contact
func UpdateContact(contacts map[int]Contact, id int, nom, email string) error {
	c, ok := contacts[id]
	if !ok {
		return fmt.Errorf("aucun contact avec l'ID %d", id)
	}
	if nom != "" {
		c.Nom = nom
	}
	if email != "" {
		c.Email = email
	}
	contacts[id] = c
	return nil
}