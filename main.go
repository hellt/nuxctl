package main

import (
	"nuxctl/cmd"
)

func main() {
	var (
		Version = "0.3.0"
	)
	cmd.Execute(Version)
}
