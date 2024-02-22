package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Ciaoo!")
	out, err := exec.Command("make clean").Output()

	if err != nil {
		fmt.Println("The command did not complete successfully")
	} else {
		fmt.Println("The output is: ", string(out))
	}
}
