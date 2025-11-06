package main

import (
	"github.com/loulounav78/efrei-go-project/internal/storage"
	"github.com/loulounav78/efrei-go-project/internal/app"
)

// La fonction main orchestre l'application. Elle dépend de l'interface Storer,
// pas de MemoryStore directement. C'est ça, l'injection de dépendances !
func main() {
	var store = storage.NewMemoryStore()
	// var store = storage.NewJSONStore()
	app.Run(store)
}
