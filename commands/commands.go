package commands

type Command struct {
	Name        string
	Description string
}

var commands = map[string]Command{
	"help": {
		Name:        "help",
		Description: "Show this menu",
	},
	"init": {
		Name:        "init",
		Description: "Create cmake with default parameters",
	},
	"clean": {
		Name:        "clean",
		Description: "Removes cmaker files and folders from the project",
	},
}

func GetCommands() map[string]Command {
	return commands
}
