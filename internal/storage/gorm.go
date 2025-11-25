package storage

import (
	"fmt"
	"sync"

	"github.com/glebarez/sqlite"

	"gorm.io/gorm"
)

type GORMStore struct {
	db  *gorm.DB
	mu  sync.Mutex
}

func NewGORMStore(dbPath string) *GORMStore {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Erreur ouverture base SQLite : %v", err))
	}

	// Auto-migration crée la table SI elle n'existe pas
	if err := db.AutoMigrate(&Contact{}); err != nil {
		panic(fmt.Sprintf("Erreur migration SQLite : %v", err))
	}

	return &GORMStore{
		db: db,
	}
}

func (gs *GORMStore) Add(contact *Contact) error {
	gs.mu.Lock()
	defer gs.mu.Unlock()

	return gs.db.Create(contact).Error
}

func (gs *GORMStore) GetAll() ([]*Contact, error) {
	gs.mu.Lock()
	defer gs.mu.Unlock()

	var contacts []*Contact
	err := gs.db.Find(&contacts).Error
	return contacts, err
}

func (gs *GORMStore) GetByID(id int) (*Contact, error) {
	gs.mu.Lock()
	defer gs.mu.Unlock()

	var contact Contact
	res := gs.db.First(&contact, id)

	if res.Error != nil {
		return nil, fmt.Errorf("contact ID %d introuvable", id)
	}

	return &contact, nil
}

func (gs *GORMStore) Update(id int, newName, newEmail string) error {
	gs.mu.Lock()
	defer gs.mu.Unlock()

	var c Contact
	if err := gs.db.First(&c, id).Error; err != nil {
		return fmt.Errorf("contact ID %d introuvable", id)
	}

	if newName != "" {
		c.Name = newName
	}
	if newEmail != "" {
		c.Email = newEmail
	}

	return gs.db.Save(&c).Error
}

func (gs *GORMStore) Delete(id int) error {
	gs.mu.Lock()
	defer gs.mu.Unlock()

	return gs.db.Delete(&Contact{}, id).Error
}
