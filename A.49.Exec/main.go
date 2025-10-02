package main

import (
	"fmt"
	"os/exec"
)

func main() {
	output1, err := exec.Command("ls").Output()
	fmt.Printf("ls: \n%s\n", string(output1))
	if err != nil {
		panic(err.Error())
	}

	goVersion, err := exec.Command("go", "version").Output()
	fmt.Printf("go version: \n%s\n", string(goVersion))
	if err != nil {
		panic(err.Error())
	}

	bunVersion, err := exec.Command("bun", "-v").Output()
	fmt.Printf("bun version: \n%s\n", string(bunVersion))
	if err != nil {
		panic(err.Error())
	}

}
