# Gestionnaire de Tâches avec Go et Gin

Ce projet est une application simple de gestion de tâches développée en **Go** avec le framework web **Gin**.
L'application permet de gérer des tâches via une interface utilisateur web interactive, comprenant des fonctionnalités pour :

- Ajouter une nouvelle tâche.
- Voir toutes les tâches.
- Marquer une tâche comme terminée.
- Supprimer une tâche.

---

## Fonctionnalités

- **CRUD des tâches :**
  - Créer, lire, mettre à jour et supprimer des tâches via une API REST.
- **Interface utilisateur :**
  - Interface web simple pour gérer les tâches.
- **Stockage JSON :**
  - Sauvegarde des tâches dans un fichier JSON pour une persistance simple.

---

## Installation

### Prérequis

- **Go** (version 1.19 ou supérieure)
- Un terminal pour exécuter les commandes.

### Étapes

1. Clonez le projet :

   ```bash
   git clone https://github.com/has1elb/taskmanager.git
   cd taskmanager
   ```
2. Installez les dépendances :

   ```bash
   go mod tidy
   ```
3. Démarrez l'application :

   ```bash
   go run .
   ```
4. Accédez à l'application via votre navigateur :

   ```
   http://localhost:8080
   ```

---

## Structure du Projet

```plaintext
taskmanager/
├── main.go           // Point d'entrée du programme
├── task.go           // Définition de la structure Task et fonctions associées
├── storage.go        // Gestion de la persistance des tâches dans un fichier JSON
├── webserver.go      // Serveur Gin avec API REST et interface utilisateur
├── templates/
│   └── index.html    // Interface utilisateur HTML
├── static/
│   ├── style.css     // Fichier de style pour l'interface utilisateur
│   └── script.js     // Script JavaScript pour les interactions dynamiques
├── tasks.json        // Fichier de sauvegarde des tâches (généré automatiquement)
└── go.mod            // Fichier des dépendances Go
```

---

## Utilisation

### 1. API REST

#### Endpoints disponibles :

- **Lister toutes les tâches** : `GET /tasks`
- **Ajouter une tâche** : `POST /tasks`
  - Corps JSON : `{"title": "Nom de la tâche"}`
- **Marquer une tâche comme terminée** : `PUT /tasks/:id`
- **Supprimer une tâche** : `DELETE /tasks/:id`

### 2. Interface Utilisateur

L'interface web est accessible à l'adresse [http://localhost:8080](http://localhost:8080). Elle permet de :

- Ajouter une tâche via un formulaire.
- Voir toutes les tâches dans une liste.
- Marquer une tâche comme terminée ou la supprimer via des boutons interactifs.

---

## Captures d'écran

### Interface utilisateur principale :

![Interface utilisateur principale](https://via.placeholder.com/800x400?text=Capture+d%27%C3%A9cran)

---

## Contribution

Les contributions sont les bienvenues ! Suivez les étapes ci-dessous pour contribuer :

1. Forkez le projet.
2. Créez une branche pour votre fonctionnalité (`git checkout -b ma-nouvelle-fonctionnalite`).
3. Effectuez vos modifications et validez-les (`git commit -m "Ajout d'une nouvelle fonctionnalité"`).
4. Poussez vos modifications (`git push origin ma-nouvelle-fonctionnalite`).
5. Créez une Pull Request.

---

## Licence

Ce projet est sous licence **MIT**. Consultez le fichier [LICENSE](LICENSE) pour plus d'informations.

---

## Auteurs

- **[E](https://github.com/<votre_nom_d_utilisateur>)L BAHRAOUI HASSAN**

🎉 Merci d'utiliser le gestionnaire de tâches ! Bon codage ! 🚀
