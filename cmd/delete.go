package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	deleteCmd := &cobra.Command{
		Use:   "delete [id]",
		Short: "Supprimer un contact",
		Args:  cobra.ExactArgs(1),
		Run:   runDeleteContact,
	}
	rootCmd.AddCommand(deleteCmd)
}

func runDeleteContact(cmd *cobra.Command, args []string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("L'ID doit être un nombre.")
		return
	}

	if err := store.Delete(id); err != nil {
		fmt.Printf("Erreur : %v\n", err)
		return
	}

	fmt.Printf("OK  Contact avec ID %d supprimé.\n", id)
}
