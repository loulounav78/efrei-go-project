package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	updateCmd := &cobra.Command{
		Use:   "update [id]",
		Short: "Mettre à jour un contact existant",
		Args:  cobra.ExactArgs(1),
		Run:   runUpdateContact,
	}
	rootCmd.AddCommand(updateCmd)
}

func runUpdateContact(cmd *cobra.Command, args []string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("L'ID doit être un nombre.")
		return
	}

	contact, err := store.GetByID(id)
	if err != nil {
		fmt.Printf("Erreur : %v\n", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Nom actuel (%s) : ", contact.Name)
	newName, _ := reader.ReadString('\n')
	newName = strings.TrimSpace(newName)

	fmt.Printf("Email actuel (%s) : ", contact.Email)
	newEmail, _ := reader.ReadString('\n')
	newEmail = strings.TrimSpace(newEmail)

	if err := store.Update(id, newName, newEmail); err != nil {
		fmt.Printf("Erreur : %v\n", err)
		return
	}

	fmt.Println("OK Contact mis à jour avec succès.")
}
