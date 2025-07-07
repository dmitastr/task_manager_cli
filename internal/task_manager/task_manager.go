package taskmanager

import (
	"flag"
	"fmt"

	"github.com/dmastr/task-manager-cli/internal/database/jsonstorage"
	"github.com/dmastr/task-manager-cli/internal/domain/repository"
	"github.com/dmastr/task-manager-cli/internal/domain/tasks_service"
	"github.com/dmastr/task-manager-cli/internal/presentation/printer"
	"github.com/dmastr/task-manager-cli/internal/presentation/table_print"

	_ "github.com/dmastr/task-manager-cli/internal/database/memstorage"
	_ "github.com/dmastr/task-manager-cli/internal/presentation/pretty_print"
)

const (
	addCmdName      string = "add"
	listCmdName     string = "list"
	completeCmdName string = "complete"
	dbPath          string = "internal/database/jsonstorage/data/db.json"
)

type ParsedArgs map[string]*string
type ParsedArgsBool map[string]*bool

type CommandCfg struct {
	name     string
	args     []argument
	argsBool []argumentBool
}

type Command struct {
	parser         *flag.FlagSet
	cfg            CommandCfg
	parsedArgs     ParsedArgs
	parsedArgsBool ParsedArgsBool
}

type argument struct {
	name         string
	defaultValue string
	description  string
}

type argumentBool struct {
	name         string
	defaultValue bool
	description  string
}

type TaskManager struct {
	service  service.Service
	printer  printer.Printer
	commands map[string]*Command
}

func NewTaskManager() *TaskManager {
	// db := memstorage.NewMemStorage()
	db := jsonstorage.NewJsonStorage(dbPath)
	service := tasksservice.NewService(db)
	// printer := prettyprint.NewPrettyPrint("    ")
	printer := tableprint.NewTablePrint()
	app := TaskManager{service: service, printer: printer}
	app.commands = map[string]*Command{}

	app.addCommand(CommandCfg{name: listCmdName})
	app.addCommand(CommandCfg{name: addCmdName, args: []argument{{name: "title"}}})
	app.addCommand(CommandCfg{name: completeCmdName,
		args:     []argument{{name: "id"}},
		argsBool: []argumentBool{{name: "u", defaultValue: false}},
	})

	return &app
}

func (t *TaskManager) addCommand(cm CommandCfg) {
	parsedArgs := make(ParsedArgs)
	parsedArgsBool := make(ParsedArgsBool)
	cmd := Command{parser: flag.NewFlagSet(cm.name, flag.ExitOnError), cfg: cm, parsedArgs: parsedArgs, parsedArgsBool: parsedArgsBool}
	for _, arg := range cm.args {
		cmd.parsedArgs[arg.name] = cmd.parser.String(arg.name, arg.defaultValue, arg.description)
	}
	for _, arg := range cm.argsBool {
		cmd.parsedArgsBool[arg.name] = cmd.parser.Bool(arg.name, arg.defaultValue, arg.description)
	}
	t.commands[cm.name] = &cmd
}

func (t *TaskManager) ParseArgs(cmdName string, args []string) (ParsedArgs, ParsedArgsBool) {
	t.commands[cmdName].parser.Parse(args)
	com := t.commands[cmdName]
	com.parser.Parse(args)
	return com.parsedArgs, com.parsedArgsBool
}

func (t *TaskManager) Run(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("expected subcommand, get none")
	}

	subCommand := args[1]
	cmdArguments := args[2:]

	switch subCommand {
	case listCmdName:
		tasks := t.service.GetAll()
		t.printer.PrintTable(tasks)
	case addCmdName:
		parsedArgs, _ := t.ParseArgs(addCmdName, cmdArguments)
		task, err := t.service.Put(*parsedArgs["title"])
		if err != nil {
			panic(err)
		}
		t.printer.PrintEntry(task)
	case completeCmdName:
		parsedArgs, parsedArgsBool := t.ParseArgs(completeCmdName, cmdArguments)
		task, err := t.service.MarkTask(*parsedArgs["id"], !*parsedArgsBool["u"])
		if err != nil {
			panic(err)
		}
		t.printer.PrintEntry(task)

	default:
		return fmt.Errorf("get unexpected subcommand: %s", subCommand)
	}
	return nil
}
