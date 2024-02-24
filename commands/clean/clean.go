package cleanCommand

import (
	"fmt"
	"os"
)

func handleExecErr(error error) {
	if error != nil {
		fmt.Println("Error on cmaker cleanup: ", error)
		os.Exit(1)
	}
}

func Exec() {
	userAnswer := ""

	for userAnswer == "" || (userAnswer != "y" && userAnswer != "n") {
		fmt.Print("Do you want to keep your cmake files? (y/n) ")
		fmt.Scan(&userAnswer)
	}

	err := os.RemoveAll(".cmaker")
	handleExecErr(err)

	if userAnswer == "n" {
		err = os.RemoveAll("./cmake")
		handleExecErr(err)

		err = os.Remove("./CMakeLists.txt")
		handleExecErr(err)
	}
}
