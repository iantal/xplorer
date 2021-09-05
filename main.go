package main

import (
	"github.com/iantal/xplorer/internal/runner"

	"github.com/projectdiscovery/goflags"
)

func main() {
	options := readOptions()
	

	r := runner.New(options)
	r.ListLibraries()
}

func readOptions() *runner.Options {
	options := &runner.Options{}

	flagSet := goflags.NewFlagSet()
	flagSet.StringVarP(&options.Repository, "repository", "repo", "", "Full path of the repository")
	flagSet.StringVarP(&options.OutputDirectory, "outputDir", "o", "", "Full path of the output directory")

	_ = flagSet.Parse()

	return options
}
