package cmd

import (
	"fmt"
	"os"

	"github.com/loulounav78/efrei-go-project/internal/storage"
	"github.com/spf13/cobra"
)

var (
	store storage.Storer
	rootCmd = &cobra.Command{
		Use:   "crm",
		Short: "Mini CRM CLI pour gérer vos contacts",
		Long:  "Mini CRM complet en Go, utilisant Cobra pour gérer les contacts de manière persistante.",
	}
)

func Execute() {
	// Utilisation automatique du JSONStore
	store = storage.NewJSONStore("data/contacts.json")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
