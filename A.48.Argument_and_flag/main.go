package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var argsRaw = os.Args
	fmt.Printf("%#v\n", argsRaw)

	var name = flag.String("name", "anonymous", "type your name")
	var age = flag.Int("age", 22, "type your age")

	flag.Parse()
	fmt.Printf("name\t: %s\n", *name)
	fmt.Printf("age\t: %d\n", *age)
}
