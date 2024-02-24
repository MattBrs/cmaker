package main

import (
	"fmt"
	"os"

	"github.com/cmaker/commands"
)

func listHelp() {
	fmt.Println("Cmaker is a tool to manage cmake projects")
	fmt.Println("\nList of available commands: ")

	commandsMap := commands.GetCommands()

	for key, value := range commandsMap {
		fmt.Println("   ", key, " ", value.Description)
	}
}

func initCache() {
	err := os.Mkdir(".cmaker", 0750)
	if err != nil && os.IsExist(err) {
		fmt.Println(".cmaker folder already exists, skipping")
	} else if err != nil {
		fmt.Println("There was an error in initing phase: ", err)
		return
	}
}

func handleErr(info string, error error, id int) {
	if error != nil {
		fmt.Println(info, error, id)
		// TODO: should also do cleanup
		os.Exit(1)
	}
}

func initCmake() {
	if _, err := os.Stat("CMakeLists.txt"); err == nil {
		fmt.Println("CMakeFile already exist! Quitting to avoid breakage")
		os.Exit(1)
	}

	f, err := os.Create("CMakeLists.txt")
	if err != nil {
		fmt.Println("There was an error during init: ", err)
		return
	}

	var name string
	var minVers string
	var cxxStd string

	fmt.Print("Enter the project name: ")
	fmt.Scan(&name)

	fmt.Print("Enter the cmake minimum version: ")
	fmt.Scan(&minVers)

	fmt.Print("Enter the cxx standard: ")
	fmt.Scan(&cxxStd)

	n, err := f.WriteString(fmt.Sprintf("cmake_minimum_required(VERSION %s)\n", minVers))
	handleErr("Error while writing main cmake: ", err, n)

	n, err = f.WriteString(fmt.Sprintf("project(%s LANGUAGES CXX)\n", name))
	handleErr("Error while writing main cmake: ", err, n)

	n, err = f.WriteString(fmt.Sprintf("set(CMAKE_CXX_STANDARD %s)\n", cxxStd))
	handleErr("Error while writing main cmake: ", err, n)

	n, err = f.WriteString("set(PROJECT_SOURCES main.cpp)\n")
	handleErr("Error while writing main cmake: ", err, n)

	n, err = f.WriteString(fmt.Sprintf("add_executable(%s ${PROJECT_SOURCES})\n", name))
	handleErr("Error while writing main cmake: ", err, n)

	f.Close()
}

func initCommand() {
	initCache()
	initCmake()
}

func execCommand() {
	arg := os.Args[1]
	switch arg {
	case "init":
		initCommand()
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
