package main

import (
	"fmt"
	"github.com/MattBrs/cmaker/commands"
	"github.com/MattBrs/cmaker/commands/clean"
	"github.com/MattBrs/cmaker/commands/init"
	"os"
)

func listHelp() {
	fmt.Println("Cmaker is a tool to manage cmake projects")
	fmt.Println("\nList of available commands: ")

	commandsMap := commands.GetCommands()

	for key, value := range commandsMap {
		fmt.Println("   ", key, " ", value.Description)
	}
}

func execCommand() {
	arg := os.Args[1]
	switch arg {
	case "init":
		initCommand.Exec()
	case "help":
		listHelp()
	case "clean":
		cleanCommand.Exec()
	default:
		fmt.Println("Command", arg, "not recognized, terminating")
		return
	}
}

func main() {
	if len(os.Args) == 1 {
		listHelp()
		return
	}

	execCommand()
}
