# 📇 Mini CRM en ligne de commande (Go)

Un petit projet d'apprentissage en **Golang**, simulant un mini système de gestion de contacts (CRM) en ligne de commande.

---

## 🚀 Fonctionnalités

- **Afficher un menu principal** en boucle
- **Ajouter** un contact (ID, Nom, Email)
- **Lister** tous les contacts
- **Supprimer** un contact par son ID
- **Mettre à jour** un contact existant
- **Quitter** proprement l’application

---

## 🧠 Concepts Go utilisés

- `map[int]Contact` → stockage en mémoire (clé = ID)
- `for {}` → boucle infinie du menu
- `switch` → sélection du choix utilisateur
- `if err != nil` → gestion des erreurs
- `comma ok idiom` → vérification de l’existence d’un contact
- `strconv.Atoi` → conversion string → int
- `os.Stdin` + `bufio` → lecture propre du clavier
- `strings.TrimSpace` → nettoyage des entrées
- `sort.Ints` → tri des contacts par ID

---

## :gear: Installation & Exécution

### :one: Cloner & lancer le projet

- git clone https://github.com/loulounav78/efrei-go-project.git
- cd efrei-go-project

- go run cmd/crm/main.go
