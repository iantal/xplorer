package detector

import "github.com/projectdiscovery/gologger"

type BuildTool int

const (
	GRADLE BuildTool = iota
	MAVEN
	SBT
	NPM
)

func toBuildTool(tool string) BuildTool {
	switch {
	case tool == "gradle":
		return GRADLE
	case tool == "maven":
		return MAVEN
	case tool == "sbt":
		return SBT
	case tool == "npm":
		return NPM
	}
	gologger.Warning().Msg("Defaulting to gradle as a build tool")
	return GRADLE
}

func (e BuildTool) String() string {
	switch e {
	case GRADLE:
		return "gradle"
	case MAVEN:
		return "maven"
	case SBT:
		return "sbt"
	case NPM:
		return "npm"
	default:
		return "unknown"
	}
}
