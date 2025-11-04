package app

import (
	"fmt"
	"os"

	"github.com/loulounav78/efrei-go-project/internal/storage"
	"github.com/loulounav78/efrei-go-project/utils"
)

func Run(store storage.Storer) {
	

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
			handleAddContact(store)
		case "2":
			handleListContacts(store)
		case "3":
			handleDeleteContact(store)
		case "4":
			handleUpdateContact(store)
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

func handleAddContact(store storage.Storer) {
	nom := utils.ReadString("Nom : ")
	email := utils.ReadString("Email : ")
	
	newContact, err := storage.NewContact(0, nom, email)
	if err != nil {
		fmt.Println("-------v-v-v-------")
		fmt.Println("Erreur :", err)
		return
	}

	if err := store.Add(newContact); err != nil {
		fmt.Println("-------v-v-v-------")
		fmt.Println("Erreur :", err)
		return
	}
	fmt.Println("-------v-v-v-------")
	fmt.Printf("Contact ajouté avec succès (ID: %d)\n", newContact.ID)
}

func handleListContacts(store storage.Storer) {
	contacts := store.List()
	fmt.Println("-------v-v-v-------")
	if len(contacts) == 0 {
		fmt.Println("Aucun contact enregistré.")
		return
	}
	fmt.Println("Liste des contacts :")
	for _, c := range contacts {
		fmt.Printf("ID: %d | Nom: %s | Email: %s\n", c.ID, c.Nom, c.Email)
	}
}

func handleDeleteContact(store storage.Storer) {
	id, err := utils.ReadInt("Entrez l'ID du contact à supprimer : ")
	if err != nil {
		fmt.Println("-------v-v-v-------")
		fmt.Println("Erreur : ID invalide.")
		return
	}

	if err := store.Delete(id); err != nil {
		fmt.Println("-------v-v-v-------")
		fmt.Println("Erreur :", err)
		return
	}
	fmt.Println("-------v-v-v-------")
	fmt.Println("Contact supprimé avec succès.")
}

func handleUpdateContact(store storage.Storer) {
	id, err := utils.ReadInt("Entrez l'ID du contact à mettre à jour : ")
	if err != nil {
		fmt.Println("-------v-v-v-------")
		fmt.Println("Erreur : ID invalide.")
		return
	}

	c, err := store.Get(id)
	if err != nil {
		fmt.Println("-------v-v-v-------")
		fmt.Println("Erreur :", err)
		return
	}

	nom := utils.ReadString("Nouveau nom (laisser vide pour ne pas changer) : ")
	email := utils.ReadString("Nouvel email (laisser vide pour ne pas changer) : ")

	if err := c.Update(nom, email); err != nil {
		fmt.Println("-------v-v-v-------")
		fmt.Println("Erreur :", err)
		return
	}
	fmt.Println("-------v-v-v-------")
	fmt.Println("Contact mis à jour avec succès.")
}

