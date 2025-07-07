package tableprint

import (
	"github.com/rodaine/table"
	"github.com/fatih/color"

	"github.com/dmastr/task-manager-cli/internal/domain/entity"
)

type TablePrint struct {
	table.Table
}

func NewTablePrint() *TablePrint {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
  	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("ID", "Completed", "Title")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	
	tp := TablePrint{tbl}
	return &tp
}

func (tp *TablePrint) PrintEntry(task entity.Task) {
	tp.PrintTable([]entity.Task{task})
}

func (tp *TablePrint) PrintTable(tasks []entity.Task) {
	for _, task := range tasks {
		checkBox := tp.ToCheckbox(task.IsCompleted)
		tp.AddRow(task.Id, checkBox, task.Text)
	}

	tp.Print()
}

func (tp *TablePrint) PrintMessage(msg string) {
	println(msg)
}

func (tp *TablePrint) ToCheckbox(isDone bool) string {
	if isDone {
		return string(rune(0x2611))
	}
	return string(rune(0x2610))
}
