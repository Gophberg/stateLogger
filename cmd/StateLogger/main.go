package main

import (
	"fmt"
	stateLogger "stateLogger/internal/app/StateLogger"
)

func main() {
	config := stateLogger.NewConfig()
	fmt.Println(config)
}
