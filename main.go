package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	models "github.com/dafaak/go-cli-todo/tasks"
	"io"
	"os"
	"strconv"
	"strings"
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
		case "add":

			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Whats your task?")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			tasks = AddTask(name, tasks)
			UpdateJson(file, tasks)
		case "delete":
			if len(os.Args) < 3 {
				fmt.Println("Task ID is required to delete")
				return
			}

			id, errConv := strconv.Atoi(os.Args[2])
			if errConv != nil {
				fmt.Println("ERROR ID must be an int: %V", errConv)
				return
			}
			tasks = DeleteTask(id, tasks)
			UpdateJson(file, tasks)
		default:
			fmt.Println("ERROR: unknown command")
		}

	}

}

func AddTask(task string, tasks []models.Task) []models.Task {

	newTask := models.Task{
		ID:       GenNextId(tasks),
		DESC:     task,
		COMPLETE: false,
	}
	return append(tasks, newTask)
}

func DeleteTask(taskId int, tasks []models.Task) []models.Task {
	for i, task := range tasks {
		if task.ID == taskId {
			return append(tasks[:i], tasks[i+1:]...)
		}
	}
	return tasks
}

func GenNextId(tasks []models.Task) int {

	lastId := len(tasks)

	return lastId + 1

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

func UpdateJson(file *os.File, tasks []models.Task) {

	bytes, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println("ERROR marshalling data: %V", err)
	}

	_, errSeek := file.Seek(0, 0)
	if errSeek != nil {
		fmt.Println("ERROR seeking data: %V", errSeek)
	}

	errTruncate := file.Truncate(0)
	if errTruncate != nil {
		fmt.Println("ERROR truncating data: %V", errTruncate)
	}

	writer := bufio.NewWriter(file)
	_, errWrite := writer.Write(bytes)
	if errWrite != nil {
		fmt.Println("ERROR writing data: %V", errWrite)
	}

	errFlush := writer.Flush()
	if errFlush != nil {
		fmt.Println("ERROR flushing data: %V", errFlush)
	}

}
