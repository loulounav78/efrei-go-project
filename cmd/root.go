package cmd

import (
	"fmt"
	"os"

	"github.com/loulounav78/efrei-go-project/internal/storage"
	"github.com/spf13/cobra"
)

var (
	storeType string
	store     storage.Storer
)

var rootCmd = &cobra.Command{
	Use:   "crm",
	Short: "Mini CRM CLI",
	Long:  "Mini CRM complet utilisant un système de stockage interchangeable (JSON ou SQLite via GORM).",
}

func init() {
	rootCmd.PersistentFlags().StringVar(&storeType, "store", "gorm", "Type de stockage: json | gorm")
}

func Execute() {
	switch storeType {

	case "json":
		store = storage.NewJSONStore("data/contacts.json")

	case "gorm":
		store = storage.NewGORMStore("data/contacts.db")

	default:
		fmt.Println("Store inconnu. Utiliser: json ou gorm")
		os.Exit(1)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
