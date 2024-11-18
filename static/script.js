document.addEventListener("DOMContentLoaded", () => {
    const taskList = document.getElementById("task-list");
    const form = document.getElementById("add-task-form");

    // Charger les tâches depuis l'API
    const loadTasks = async () => {
        const response = await fetch("/tasks");
        const tasks = await response.json();
        taskList.innerHTML = ""; // Effacer la liste existante
        tasks.forEach(task => {
            const li = document.createElement("li");
            li.textContent = `${task.title} ${task.completed ? "(Terminé)" : ""}`;

            // Ajouter un bouton pour marquer comme terminé
            if (!task.completed) {
                const completeButton = document.createElement("button");
                completeButton.textContent = "Terminer";
                completeButton.addEventListener("click", async () => {
                    await fetch(`/tasks/${task.id}`, {
                        method: "PUT"
                    });
                    loadTasks();
                });
                li.appendChild(completeButton);
            }

            // Ajouter un bouton pour supprimer la tâche
            const deleteButton = document.createElement("button");
            deleteButton.textContent = "Supprimer";
            deleteButton.addEventListener("click", async () => {
                await fetch(`/tasks/${task.id}`, {
                    method: "DELETE"
                });
                loadTasks();
            });
            li.appendChild(deleteButton);

            taskList.appendChild(li);
        });
    };

    // Ajouter une nouvelle tâche
    form.addEventListener("submit", async (event) => {
        event.preventDefault();
        const title = document.getElementById("task-title").value;
        if (!title) return;

        await fetch("/tasks", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ title })
        });
        document.getElementById("task-title").value = "";
        loadTasks();
    });

    loadTasks(); // Charger les tâches au chargement de la page
});
