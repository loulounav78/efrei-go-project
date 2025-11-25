package cmd

import (
	"fmt"
	"os"

	"github.com/loulounav78/efrei-go-project/internal/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	store storage.Storer
)

var rootCmd = &cobra.Command{
	Use:   "crm",
	Short: "Mini CRM CLI",
	Long:  "Mini CRM utilisant un système de stockage configurable via Viper.",
}

func init() {
	cobra.OnInitialize(initConfig)

	// Permettre aussi l’override par flag (optionnel)
	rootCmd.PersistentFlags().String("store", "", "Type de stockage (json | gorm)")
	viper.BindPFlag("storage.type", rootCmd.PersistentFlags().Lookup("store"))
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Impossible de lire config.yaml :", err)
	}

	viper.SetDefault("storage.type", "gorm")

	storageType := viper.GetString("storage.type")

	switch storageType {

	case "json":
		path := viper.GetString("storage.json_path")
		if path == "" {
			path = "data/contacts.json"
		}
		fmt.Println("Utilisation du JSONStore :", path)
		store = storage.NewJSONStore(path)

	case "gorm":
		path := viper.GetString("storage.db_path")
		if path == "" {
			path = "data/contacts.db"
		}
		fmt.Println("Utilisation du GORMStore (SQLite) :", path)
		store = storage.NewGORMStore(path)

	default:
		fmt.Println("Type de stockage inconnu dans config.yaml :", storageType)
		os.Exit(1)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
