# Mini CRM en Go (Cobra + Viper + GORM + JSON)

Un mini CRM en ligne de commande, entièrement écrit en Go, avec :

- Une CLI professionnelle utilisant **Cobra**
- Une configuration externe via **Viper**
- Une architecture propre & découplée
- De multiples backends de stockage :
  - **SQLite / GORM** (par défaut)
  - **JSON**
  - **Mémoire** (pour tests)
- Un CRUD complet : **add, list, update, delete**

---

# 📂 Structure du projet

```
efrei-go-project/
├── cmd/                 # Commandes CLI (Cobra)
│   ├── root.go
│   ├── add.go
│   ├── list.go
│   ├── update.go
│   └── delete.go
│
├── internal/
│   └── storage/         # Backends de stockage
│       ├── storage.go   # Interface Storer + struct Contact
│       ├── memory.go
│       ├── json.go
│       └── gorm.go
│
├── config.yaml          # Configuration externe (Viper)
├── go.mod
├── go.sum
└── main.go
```

---

# 🚀 Installation & Démarrage

## 1. Cloner le projet

```bash
git clone <URL_DU_REPO>
cd efrei-go-project
```

---

## 2. Installer les dépendances

```bash
go mod tidy
```

---

## 3. Fichier de configuration (Viper)

Le fichier `config.yaml` doit être à la racine du projet :

```yaml
storage:
  type: gorm         # json | gorm | memory
  db_path: data/contacts.db
  json_path: data/contacts.json
```

---

## 4. Lancer le programme

```bash
go run .
```

Le backend utilisé dépend de `config.yaml`.

---

# 🛠️ Backends disponibles

## 🔷 1. SQLite via GORM (recommandé, et par défaut)

```yaml
storage:
  type: gorm
  db_path: data/contacts.db
```

Lancement :

```bash
go run . list
```

---

## 🟧 2. JSON (format simple et lisible)

```yaml
storage:
  type: json
  json_path: data/contacts.json
```

---

## 🟩 3. Mémoire (aucune persistance)

```yaml
storage:
  type: memory
```

---

# 🔧 Override du backend avec un flag

Tu peux override la config (Viper) avec :

```bash
go run . --store json list
go run . --store gorm list
go run . --store memory list
```

---

# 📚 Commandes CLI

Voici toutes les commandes disponibles via Cobra :

---

## ➕ Ajouter un contact

### Mode interactif :

```bash
go run . add
```

### Avec flags :

```bash
go run . add -n "Alice" -e "alice@mail.com"
```

---

## 📋 Lister les contacts

```bash
go run . list
```

---

## ✏️ Modifier un contact existant

```bash
go run . update <id>
```

Exemple :

```bash
go run . update 2
```

L’interface te demandera les nouvelles valeurs.

---

## 🗑️ Supprimer un contact

```bash
go run . delete <id>
```

Exemple :

```bash
go run . delete 3
```

---

# 🧪 Scénario complet pour tester tout

Voici un scénario simple pour vérifier que tout fonctionne :

---

## 1. Ajouter des contacts

```bash
go run . add -n "Alice" -e "alice@mail.com"
go run . add -n "Loris" -e "loris@mail.com"
```

---

## 2. Lister

```bash
go run . list
```

---

## 3. Modifier un contact

```bash
go run . update 1
```

---

## 4. Supprimer un contact

```bash
go run . delete 2
```

---

## 5. Tester JSON

```bash
go run . --store json add -n "JSONUser" -e "json@test.com"
go run . --store json list
```

---

## 6. Tester MemoryStore (pas de persistance)

```bash
go run . --store memory add -n "Temp" -e "temp@test.com"
go run . --store memory list     # OK
go run . --store memory list     # vide → normal !
```

---

# ♻️ Rebuild du binaire

Pour compiler le CLI :

```bash
go build -o crm .
```

Puis :

```bash
./crm list
```

---

# 🎉 Félicitations

Vous disposez maintenant d’un CRM complet, modulaire et extensible :

- CLI professionnelle  
- Backends interchangeables  
- Config externe  
- Persistance fiable  
- Architecture propre et évolutive  

N’hésitez pas à ajouter : export CSV, recherche, filtres, logs, tests…
