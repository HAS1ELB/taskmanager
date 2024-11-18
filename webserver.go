package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var tasks []Task // Variable globale pour stocker les tâches

// Configuration des routes du serveur
func setupRouter() *gin.Engine {
	router := gin.Default()

	// Servir des fichiers statiques (CSS, JS, etc.)
	router.Static("/static", "./static")

	// Route pour l'interface graphique
	router.GET("/", func(c *gin.Context) {
		c.File("templates/index.html")
	})

	// Route pour afficher toutes les tâches en JSON
	router.GET("/tasks", func(c *gin.Context) {
		c.JSON(http.StatusOK, tasks)
	})

	// Route pour ajouter une tâche
	router.POST("/tasks", func(c *gin.Context) {
		var newTask struct {
			Title string `json:"title"`
		}
		if err := c.ShouldBindJSON(&newTask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		task := NewTask(len(tasks)+1, newTask.Title)
		tasks = append(tasks, task)
		c.JSON(http.StatusCreated, task)
	})

	// Route pour marquer une tâche comme terminée
	router.PUT("/tasks/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
			return
		}

		for i := range tasks {
			if tasks[i].ID == id {
				tasks[i].Completed = true
				c.JSON(http.StatusOK, tasks[i])
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
	})

	// Route pour supprimer une tâche
	router.DELETE("/tasks/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
			return
		}

		for i, task := range tasks {
			if task.ID == id {
				tasks = append(tasks[:i], tasks[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
	})

	return router
}
