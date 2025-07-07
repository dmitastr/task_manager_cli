package prettyprint

import (
	"encoding/json"
	"fmt"

	"github.com/dmastr/task-manager-cli/internal/domain/entity"
)


type PrettyPrint struct {
	indent string
}

func NewPrettyPrint(indent string) *PrettyPrint {
	pp := PrettyPrint{indent: indent}
	return &pp
}

func (pp *PrettyPrint) PrintTable(tasks []entity.Task) {
	res := pp.ToString(tasks)
	fmt.Println(res)
}

func (pp *PrettyPrint) PrintEntry(task entity.Task) {
	res := pp.ToString(task)
	fmt.Println(res)
}

func (pp *PrettyPrint) PrintMessage(msg string) {
	fmt.Println(msg)
}

func (pp *PrettyPrint) ToString(tasks any) string {
	res, err := json.MarshalIndent(tasks, "", pp.indent)
	if err != nil {
		panic(err)
	}
	return string(res)
}