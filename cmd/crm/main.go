package main

import (
	"github.com/loulounav78/efrei-go-project/internal/storage"
	"github.com/loulounav78/efrei-go-project/internal/app"
)

func main() {
	store := storage.NewMemoryStore()
	app.Run(store)
}