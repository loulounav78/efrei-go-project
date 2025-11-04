package storage

import (
	"errors"
	"fmt"
	"regexp"
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
