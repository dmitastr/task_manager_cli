package printer


import "github.com/dmastr/task-manager-cli/internal/domain/entity"


type Printer interface {
	PrintTable([]entity.Task)
	PrintEntry(entity.Task)
	PrintMessage(string)
}