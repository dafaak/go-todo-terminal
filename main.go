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

}
