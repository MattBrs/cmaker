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
		Description: "Create cmake with default parameter",
	},
}

func GetCommands() map[string]Command {
	return commands
}
