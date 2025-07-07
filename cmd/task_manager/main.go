package main

import (
	"os"

	"github.com/dmastr/task-manager-cli/internal/task_manager"
)


func main() {
	app := taskmanager.NewTaskManager()
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}