package main

import (
	"encoding/json"
	"errors"
	"os"
)

func SaveTasks(fileName string, tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

func LoadTasks(fileName string) ([]Task, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Task{}, nil
		}
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}
