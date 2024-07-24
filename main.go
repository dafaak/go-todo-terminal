package main

import (
	"encoding/json"
	"fmt"
	models "github.com/dafaak/go-cli-todo/tasks"
	"io"
	"os"
)

func main() {

	file, err := os.OpenFile("tasks.json", os.O_RDWR, 0666)

	if err != nil {
		fmt.Println("ERROR: %V", err)
	}

	defer file.Close()

	var tasks []models.Task

	info, err := file.Stat()

	if err != nil {
		fmt.Println("ERROR reading file info: %V", err)
	}

	if info.Size() != 0 { // si el archivo no está vacío

		bytesm, err := io.ReadAll(file)

		if err != nil {
			fmt.Println("ERROR reading data: %V", err)
		}

		err = json.Unmarshal(bytesm, &tasks)
		if err != nil {
			fmt.Println("ERROR parsing data: %V", err)
		}
	} else {
		tasks = make([]models.Task, 0)
		//tasks = []models.Task{}
	}

	if len(os.Args) < 2 {
		printUsage()
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "list":
			ListTasks(tasks)
		default:
			fmt.Println("ERROR: unknown command")
		}

	}

}

func ListTasks(tasks []models.Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}
	fmt.Println("Tasks:")
	for _, task := range tasks {

		status := "🤗"

		if !task.COMPLETE {
			status = "🥲"
		}

		fmt.Printf("[%s] [%d] %s \n", status, task.ID, task.DESC)
	}
}

func printUsage() {
	fmt.Println("task-manager:[list|add|delete|complete]")
}
