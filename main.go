package main

import (
	"nuxctl/cmd"
)

func main() {
	var (
		Version = "0.1.0"
	)
	cmd.Execute(Version)
}
