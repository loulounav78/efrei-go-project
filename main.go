package main

import (
	"fmt"
	"os"

	"github.com/loulounav78/efrei-go-project/contact"
	"github.com/loulounav78/efrei-go-project/utils"
)

func main() {
	contacts := make(map[int]*contact.Contact)
	nextID := 1 // ID auto-incrémenté

	for {
		fmt.Println("\n--- MINI CRM ---")
		fmt.Println("1. Ajouter un contact")
		fmt.Println("2. Lister les contacts")
		fmt.Println("3. Supprimer un contact")
		fmt.Println("4. Mettre à jour un contact")
		fmt.Println("5. Quitter")

		choice := utils.ReadString("Choisissez une option : ")

		switch choice {
		case "1":
			nom := utils.ReadString("Nom : ")
			email := utils.ReadString("Email : ")

			newContact, err := contact.NewContact(nextID, nom, email)
			if err != nil {
				fmt.Println("Erreur :", err)
				continue
			}

			if err := newContact.Add(contacts); err != nil {
				fmt.Println("Erreur :", err)
			} else {
				fmt.Printf("Contact ajouté avec succès (ID: %d)\n", nextID)
				nextID++
			}

		case "2":
			fmt.Println("-------v-v-v-------")
			contact.ListContacts(contacts)

		case "3":
			id, err := utils.ReadInt("Entrez l'ID du contact à supprimer : ")
			if err != nil {
				fmt.Println("Erreur : ID invalide.")
				continue
			}
			fmt.Println("-------v-v-v-------")
			if err := contact.DeleteContact(contacts, id); err != nil {
				fmt.Println("Erreur :", err)
			} else {
				fmt.Println("Contact supprimé.")
			}

		case "4":
			id, err := utils.ReadInt("Entrez l'ID du contact à mettre à jour : ")
			if err != nil {
				fmt.Println("Erreur : ID invalide.")
				continue
			}

			c, ok := contacts[id]
			if !ok {
				fmt.Println("Aucun contact trouvé avec cet ID.")
				continue
			}

			nom := utils.ReadString("Nouveau nom (laisser vide pour ne pas changer) : ")
			email := utils.ReadString("Nouvel email (laisser vide pour ne pas changer) : ")
			fmt.Println("-------v-v-v-------")
			if err := c.Update(nom, email); err != nil {
				fmt.Println("Erreur :", err)
			} else {
				fmt.Println("Contact mis à jour.")
			}

		case "5":
			fmt.Println("-------v-v-v-------")
			fmt.Println("Merci de votre participation !")
			os.Exit(0)

		default:
			fmt.Println("-------v-v-v-------")
			fmt.Println("Option invalide, veuillez réessayer.")
		}
		fmt.Println("------------------------^^^------------------------")
	}
}
