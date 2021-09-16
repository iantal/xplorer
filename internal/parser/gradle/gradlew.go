package gradle

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/iantal/xplorer/internal/utils"
	"github.com/projectdiscovery/gologger"
)

// Gradlew handles execution of various gradlew comands
type Gradlew struct {
	projectRoot string
}

// NewGradlew created a new Gradlew
func NewGradlew(projectRoot string) *Gradlew {
	return &Gradlew{projectRoot}
}

// Projects parses the output of `gradlew projects`
func (g *Gradlew) Projects() ([]string, error) {
	if !g.hasGradlew() {
		return nil, fmt.Errorf("Could not execute gradlew")
	}

	err, stdout, stderr := utils.CMD("./gradlew", "projects")
	if err != nil {
		return nil, fmt.Errorf("Error executing [gradlew projects] command for %s", g.projectRoot)
	}

	if isFailedCommand(stdout) || len(stderr) != 0 {
		return nil, fmt.Errorf("Received build failure for [gradlew projects] for %s", g.projectRoot)
	}

	return extractProjects(stdout), nil
}

func isFailedCommand(stdout []string) bool {
	for _, line := range stdout {
		if strings.Contains(line, "BUILD SUCCESSFUL") {
			return false
		}
	}
	return true
}

func extractProjects(data []string) []string {
	var res []string
	for _, out := range data {
		p := extractProject(out)
		if p != "" {
			res = append(res, p)
		}
	}
	return res
}

func extractProject(line string) string {
	var re = regexp.MustCompile(`(?m).*Project\s+':(?P<project>.*)'`)

	result := make(map[string]string)
	match := re.FindStringSubmatch(line)
	if match == nil {
		return ""
	}
	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	return result["project"]
}

// Dependencies executes `gradlew [proj]:dependencies`
func (g *Gradlew) Dependencies(project string, isSubproject bool) string {
	gologger.Debug().Str("project", project).Msg("Generating dependency tree")

	if !isSubproject {
		err, stdout, serr := utils.CMD("./gradlew", "dependencies")
		if err != nil {
			gologger.Error().Msgf("Error executing [gradlew dependencies].\n%s", strings.Join(serr, "\n"))
			return ""
		}
		return strings.Join(stdout, "\n")
	}

	c := project + ":dependencies"
	err, stdout, serr := utils.CMD("./gradlew", c)
	if err != nil {
		gologger.Error().Msgf("Error executing [gradlew %s:dependencies].\n%s", project, strings.Join(serr, "\n"))
		return ""
	}
	return strings.Join(stdout, "\n")
}

func (g *Gradlew) hasGradlew() bool {
	os.Chdir(g.projectRoot)
	err, stdout, _ := utils.CMD("ls")
	if err != nil {
		gologger.Error().Str("path", g.projectRoot).Msg(err.Error())
		return false
	}
	for _, out := range stdout {
		if out == "gradlew" {
			return true
		}
	}
	return false
}
