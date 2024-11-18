package main

func main() {
	// Charger les tâches depuis le fichier JSON
	var err error
	tasks, err = LoadTasks("tasks.json")
	if err != nil {
		panic("Erreur lors du chargement des tâches : " + err.Error())
	}

	// Démarrer le serveur web
	router := setupRouter()
	router.Run(":8080") // Le serveur écoutera sur le port 8080
}
