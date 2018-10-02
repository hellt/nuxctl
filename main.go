package main

import "github.com/nuagenetworks/nuxctl/cmd"

func main() {
	var (
		Version = "0.4.1"
	)
	cmd.Execute(Version)
}
