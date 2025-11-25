package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "Lister tous les contacts",
		Run:   runListContacts,
	}
	rootCmd.AddCommand(listCmd)
}

func runListContacts(cmd *cobra.Command, args []string) {
	contacts, err := store.GetAll()
	if err != nil {
		fmt.Printf("Erreur : %v\n", err)
		return
	}

	if len(contacts) == 0 {
		fmt.Println("Aucun contact trouvé.")
		return
	}

	fmt.Println("\n--- Liste des contacts ---")
	for _, c := range contacts {
		fmt.Printf("ID: %d | Nom: %s | Email: %s\n", c.ID, c.Name, c.Email)
	}
}
