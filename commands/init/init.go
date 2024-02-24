package initCommand

import (
	"fmt"
	"os"
)

func handleInitErr(info string, error error, id int) {
	if error != nil {
		fmt.Println(info, error, id)
		// TODO: should also do cleanup
		os.Exit(1)
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

func initCmake() {
	if _, err := os.Stat("CMakeLists.txt"); err == nil {
		fmt.Println("CMakeFile already exist! Quitting to avoid breakage")
		os.Exit(1)
	}

	f, err := os.Create("CMakeLists.txt")
	if err != nil {
		fmt.Println("There was an error during init: ", err)
		os.Exit(1)
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
	handleInitErr("Error while writing main cmake: ", err, n)

	n, err = f.WriteString(fmt.Sprintf("project(%s LANGUAGES CXX)\n", name))
	handleInitErr("Error while writing main cmake: ", err, n)

	n, err = f.WriteString(fmt.Sprintf("set(CMAKE_CXX_STANDARD %s)\n", cxxStd))
	handleInitErr("Error while writing main cmake: ", err, n)

	n, err = f.WriteString("set(PROJECT_SOURCES main.cpp)\n")
	handleInitErr("Error while writing main cmake: ", err, n)

	n, err = f.WriteString(fmt.Sprintf("add_executable(%s ${PROJECT_SOURCES})\n", name))
	handleInitErr("Error while writing main cmake: ", err, n)

	n, err = f.WriteString("include(${CMAKE_CURRENT_SOURCE_DIR}/cmake/include.cmake)\n")
	handleInitErr("Error while writing main cmake: ", err, n)

	n, err = f.WriteString("include(${CMAKE_CURRENT_SOURCE_DIR}/cmake/link.cmake)\n")
	handleInitErr("Error while writing main cmake: ", err, n)

	f.Close()

	err = os.Mkdir("cmake", 0750)
	if err != nil && os.IsExist(err) {
		fmt.Println("cmake folder already exists, quitting to avoid breakage")
		os.Exit(1)
	} else if err != nil {
		fmt.Println("There was an error in initing phase: ", err)
		os.Exit(1)
	}

	f, _ = os.Create("cmake/include.cmake")
	f.Close()

	f, _ = os.Create("cmake/link.cmake")
	f.Close()
}

func Exec() {
	initCache()
	initCmake()
}
