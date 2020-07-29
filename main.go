package main

import (
	"fmt"
	"os"
	"random-ranker/cmd"
)

func main() {
	cmd.Execute()
}

func exitWithError(err error) {
	fmt.Fprintf(os.Stderr, "\n%v\n", err)
	os.Exit(1)
}
