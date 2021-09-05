package gradle

import (
	"regexp"
	"strings"

	"github.com/iantal/xplorer/internal/domain"
	"github.com/projectdiscovery/gologger"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(dependencyTree string) {
	libraries := []*domain.Library{}
	set := make(map[string]bool)
	lines := strings.Split(dependencyTree, "\n")

	r := regexp.MustCompile(`(?m)^(\\|\+)---`)

	for _, line := range lines {
		if r.MatchString(line) {
			name := strings.Split(line, " ")[1]
			if !set[name] {
				l := domain.NewLibrary(name, 0)
				libraries = append(libraries, l)
				set[name] = true
			}
		}
	}

	for _, l := range libraries {
		gologger.Info().Msgf("Library: %s", l.Name)
	}
}
