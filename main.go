package main

import (
	"nuxctl/cmd"
)

func main() {
	var (
		Version = "0.2.1"
	)
	cmd.Execute(Version)
}
