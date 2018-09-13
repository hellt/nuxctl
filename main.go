package main

import (
	"nuxctl/cmd"
)

func main() {
	var (
		Version = "0.2.0"
	)
	cmd.Execute(Version)
}
