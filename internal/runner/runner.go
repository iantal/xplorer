package runner

import (
	"strings"

	"github.com/iantal/xplorer/internal/detector"
	"github.com/iantal/xplorer/internal/parser/gradle"
	"github.com/projectdiscovery/gologger"
)

type Runner struct {
	options *Options
}

func New(options *Options) *Runner {
	ParseOptions(options)
	return &Runner{
		options: options,
	}
}

func (r *Runner) ListLibraries() {
	buildTools := r.detectBuildTools()
	r.generateDependencyTrees(buildTools)

}

func (r *Runner) detectBuildTools() []detector.BuildTool {
	gologger.Info().Msg("Detecting build tools")
	detector := detector.New(r.options.Repository)
	buildTools := detector.Detect()
	logBuildTools(buildTools)
	return buildTools
}

func logBuildTools(buildTools []detector.BuildTool) {
	asStringArray := []string{}
	for _, tool := range buildTools {
		asStringArray = append(asStringArray, tool.String())
	}
	gologger.Info().Msgf("Following build tools were detected: %s", strings.Join(asStringArray, ", "))
}

func (r *Runner) generateDependencyTrees(buildTools []detector.BuildTool) {
	for _, bt := range buildTools {
		switch {
		case bt == detector.GRADLE:
			r.handleGradle()
		case bt == detector.MAVEN:
			r.handleMaven()
		}
	}
}

func (r *Runner) handleGradle() {
	gp := gradle.NewParser()
	gw := gradle.NewGradlew(r.options.Repository)
	projects, err := gw.Projects()
	if err != nil {
		gologger.Error().Msg("Unable to obtain the list of projects")
		return
	}

	if len(projects) == 1 {
		gw.Dependencies(projects[0], false)

	} else {
		for _, p := range projects {
			dt := gw.Dependencies(p, true)
			gp.Parse(dt)
		}
	}
}

func (r *Runner) handleMaven() {
	// TODO: implement method
}
