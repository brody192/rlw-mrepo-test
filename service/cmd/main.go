package main

import (
	"fmt"

	"github.com/DanielTitkov/rlw-mrepo-test/pkg/lib"
	"github.com/google/uuid"
)

func main() {
	fmt.Println("Run", uuid.New()) // random external dependency

	lib.DoSomething() // local dependency

	fmt.Println("Goodbye, World!")
}
