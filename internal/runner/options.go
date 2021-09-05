package runner

import (
	"os"
	"path/filepath"

	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
)

type Options struct {
	Repository      string
	OutputDirectory string
}

func ParseOptions(options *Options) {
	gologger.DefaultLogger.SetMaxLevel(levels.LevelDebug)

	showBanner()

	if !filepath.IsAbs(options.OutputDirectory) || !filepath.IsAbs(options.Repository) {
		gologger.Error().Msg("Please specify absolute (full) paths for the repository and output directories!!")
		os.Exit(1)
	}
}
