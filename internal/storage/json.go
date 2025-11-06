package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type JSONStore struct {
	filePath string
	contacts map[int]*Contact
	nextID   int
	mu       sync.Mutex
}

func NewJSONStore(filePath string) *JSONStore {
	// Crée le dossier si nécessaire
	dir := filepath.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}

	js := &JSONStore{
		filePath: filePath,
		contacts: make(map[int]*Contact),
		nextID:   1,
	}
	js.load()
	return js
}

func (js *JSONStore) load() {
	file, err := os.Open(js.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		fmt.Printf("Erreur ouverture fichier JSON : %v\n", err)
		return
	}
	defer file.Close()

	var data []*Contact
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		fmt.Printf("Erreur décodage JSON : %v\n", err)
		return
	}

	for _, c := range data {
		js.contacts[c.ID] = c
		if c.ID >= js.nextID {
			js.nextID = c.ID + 1
		}
	}
}

func (js *JSONStore) save() error {
	file, err := os.Create(js.filePath)
	if err != nil {
		return fmt.Errorf("erreur création fichier JSON : %w", err)
	}
	defer file.Close()

	var data []*Contact
	for _, c := range js.contacts {
		data = append(data, c)
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

func (js *JSONStore) Add(contact *Contact) error {
	js.mu.Lock()
	defer js.mu.Unlock()

	contact.ID = js.nextID
	js.contacts[contact.ID] = contact
	js.nextID++

	return js.save()
}

func (js *JSONStore) GetAll() ([]*Contact, error) {
	js.mu.Lock()
	defer js.mu.Unlock()

	var all []*Contact
	for _, c := range js.contacts {
		all = append(all, c)
	}
	return all, nil
}

func (js *JSONStore) GetByID(id int) (*Contact, error) {
	js.mu.Lock()
	defer js.mu.Unlock()

	c, ok := js.contacts[id]
	if !ok {
		return nil, fmt.Errorf("contact avec ID %d non trouvé", id)
	}
	return c, nil
}

func (js *JSONStore) Update(id int, newName, newEmail string) error {
	js.mu.Lock()
	defer js.mu.Unlock()

	c, ok := js.contacts[id]
	if !ok {
		return fmt.Errorf("contact avec ID %d non trouvé", id)
	}

	if newName != "" {
		c.Name = newName
	}
	if newEmail != "" {
		c.Email = newEmail
	}

	return js.save()
}

func (js *JSONStore) Delete(id int) error {
	js.mu.Lock()
	defer js.mu.Unlock()

	if _, ok := js.contacts[id]; !ok {
		return fmt.Errorf("contact avec ID %d non trouvé", id)
	}
	delete(js.contacts, id)
	return js.save()
}
