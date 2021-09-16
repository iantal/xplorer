package detector

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/projectdiscovery/gologger"
	"golang.org/x/xerrors"
)

type BuildToolDetector struct {
	repositoryAbsPath string
	bTypes            Types
}

func New(repoAbsPath string) *BuildToolDetector {
	conf, err := LoadConfig()
	if err != nil {
		gologger.Error().Msgf("Could not load config file\n%s", err.Error())
	}

	return &BuildToolDetector{
		repositoryAbsPath: repoAbsPath,
		bTypes: conf,
	}
}

func (d *BuildToolDetector) Detect() []BuildTool {
	r := []BuildTool{}
	for _, t := range d.bTypes.BuildTypes {
		for _, file := range t.Files {
			if err := search(file, d.repositoryAbsPath); err == nil {
				r = append(r, toBuildTool(t.Name))
			}
		}
	}
	return r
}

func search(file, directory string) error {
	found := false

	err := filepath.Walk(directory,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if strings.Contains(path, file) {
				found = true
			}

			return nil
		})

	if err != nil {
		return err
	}

	if !found {
		return xerrors.Errorf("Unable to detect the build tool. Please check detector/config.yml for supported build tools")
	}

	return nil
}
