package main

import (
	"nuxctl/cmd"
)

func main() {
	var (
		Version = "0.3.1"
	)
	cmd.Execute(Version)
}
