package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/loulounav78/efrei-go-project/internal/storage"
	"github.com/spf13/cobra"
)

var (
	nameFlag  string
	emailFlag string
)

func init() {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Ajouter un contact",
		Run:   runAddContact,
	}

	addCmd.Flags().StringVarP(&nameFlag, "name", "n", "", "Nom du contact")
	addCmd.Flags().StringVarP(&emailFlag, "email", "e", "", "Email du contact")

	rootCmd.AddCommand(addCmd)
}

func runAddContact(cmd *cobra.Command, args []string) {
	reader := bufio.NewReader(os.Stdin)

	name := nameFlag
	email := emailFlag

	if name == "" {
		fmt.Print("Entrez le nom du contact : ")
		name, _ = reader.ReadString('\n')
		name = strings.TrimSpace(name)
	}

	if email == "" {
		fmt.Print("Entrez l'email du contact : ")
		email, _ = reader.ReadString('\n')
		email = strings.TrimSpace(email)
	}

	if name == "" || email == "" {
		fmt.Println("Erreur : le nom et l'email sont obligatoires.")
		return
	}

	contact := &storage.Contact{Name: name, Email: email}
	if err := store.Add(contact); err != nil {
		fmt.Printf("Erreur lors de l'ajout : %v\n", err)
		return
	}

	fmt.Printf("OK Contact '%s' ajouté avec l'ID %d.\n", contact.Name, contact.ID)
}
