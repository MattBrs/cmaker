package main

import (
	"fmt"
	"os"

	"github.com/cmaker/commands"
	"github.com/cmaker/commands/init"
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

	// out, err := exec.Command("make clean").Output()
	//
	// if err != nil {
	// 	fmt.Println("The command did not complete successfully")
	// } else {
	// 	fmt.Println("The output is: ", string(out))
	// }
}
