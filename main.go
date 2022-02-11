package main

import "github.com/devOpifex/cranlogs/cmd"

var (
	version string = "dev"
	// goreleaser will also inject these if desireable to use
	// commit  string = "none"
	// date    string = "unknown"
)

func main() {
	cmd.Execute(version)
}
