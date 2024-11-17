package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const fileName = "tasks.json"

func main() {
	tasks, err := LoadTasks(fileName)
	if err != nil {
		fmt.Println("Erreur lors du chargement des tâches :", err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n=== Gestionnaire de Tâches ===")
		fmt.Println("1. Ajouter une tâche")
		fmt.Println("2. Lister les tâches")
		fmt.Println("3. Marquer une tâche comme terminée")
		fmt.Println("4. Supprimer une tâche")
		fmt.Println("5. Quitter")
		fmt.Print("Choisissez une option : ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Titre de la tâche : ")
			scanner.Scan()
			title := scanner.Text()
			task := NewTask(len(tasks)+1, title)
			tasks = append(tasks, task)
			fmt.Println("Tâche ajoutée :", task.Title)

		case "2":
			fmt.Println("\nListe des tâches :")
			for _, task := range tasks {
				status := "À faire"
				if task.Completed {
					status = "Terminé"
				}
				fmt.Printf("%d. [%s] %s\n", task.ID, status, task.Title)
			}

		case "3":
			fmt.Print("ID de la tâche à marquer comme terminée : ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())
			for i := range tasks {
				if tasks[i].ID == id {
					tasks[i].Completed = true
					fmt.Println("Tâche marquée comme terminée :", tasks[i].Title)
				}
			}

		case "4":
			fmt.Print("ID de la tâche à supprimer : ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())
			for i, task := range tasks {
				if task.ID == id {
					tasks = append(tasks[:i], tasks[i+1:]...)
					fmt.Println("Tâche supprimée :", task.Title)
					break
				}
			}

		case "5":
			if err := SaveTasks(fileName, tasks); err != nil {
				fmt.Println("Erreur lors de la sauvegarde des tâches :", err)
			} else {
				fmt.Println("Tâches sauvegardées. À bientôt !")
			}
			return

		default:
			fmt.Println("Option invalide. Veuillez réessayer.")
		}
	}
}
