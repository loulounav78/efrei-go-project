package contact

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// Contact représente un contact client
type Contact struct {
	ID    int
	Nom   string
	Email string
}

// NewContact est un constructeur qui crée et valide un contact
func NewContact(id int, nom, email string) (*Contact, error) {
	nom = strings.TrimSpace(nom)
	email = strings.TrimSpace(email)

	// Vérification du nom
	if nom == "" {
		return nil, errors.New("le nom ne peut pas être vide")
	}
	if !isValidName(nom) {
		return nil, errors.New("le nom contient des caractères non valides (lettres et espaces uniquement)")
	}

	// Vérification de l'email
	if !isValidEmail(email) {
		return nil, errors.New("l'email n'est pas valide")
	}

	return &Contact{
		ID:    id,
		Nom:   nom,
		Email: email,
	}, nil
}

// isValidEmail vérifie que l'email est au bon format
func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// isValidName vérifie que le nom contient uniquement lettres, espaces, tirets
func isValidName(name string) bool {
	re := regexp.MustCompile(`^[a-zA-ZÀ-ÿ\s\-]+$`)
	return re.MatchString(name)
}

//
// --- Méthodes associées ---
//

// Ajouter un contact à la collection
func (c *Contact) Add(contacts map[int]*Contact) error {
	if _, exists := contacts[c.ID]; exists {
		return fmt.Errorf("un contact avec l'ID %d existe déjà", c.ID)
	}
	contacts[c.ID] = c
	return nil
}

// Met à jour les informations du contact
func (c *Contact) Update(nom, email string) error {
	if nom != "" {
		if !isValidName(nom) {
			return errors.New("le nouveau nom contient des caractères non valides")
		}
		c.Nom = nom
	}
	if email != "" {
		if !isValidEmail(email) {
			return errors.New("le nouvel email n'est pas valide")
		}
		c.Email = email
	}
	return nil
}

//
// --- Fonctions globales (utilisent la map) ---
//

// Supprimer un contact par ID
func DeleteContact(contacts map[int]*Contact, id int) error {
	if _, ok := contacts[id]; !ok {
		return fmt.Errorf("aucun contact avec l'ID %d", id)
	}
	delete(contacts, id)
	return nil
}

// Lister tous les contacts triés par ID
func ListContacts(contacts map[int]*Contact) {
	if len(contacts) == 0 {
		fmt.Println("Aucun contact enregistré.")
		return
	}

	var ids []int
	for id := range contacts {
		ids = append(ids, id)
	}
	sort.Ints(ids)

	fmt.Println("Liste des contacts :")
	for _, id := range ids {
		c := contacts[id]
		fmt.Printf("ID: %d | Nom: %s | Email: %s\n", c.ID, c.Nom, c.Email)
	}
}
