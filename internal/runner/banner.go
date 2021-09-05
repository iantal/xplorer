package runner

import (
	"fmt"

	"github.com/projectdiscovery/gologger"
)

var banner = fmt.Sprint(`
               __                    
   _  ______  / /___  ________  _____
  | |/_/ __ \/ / __ \/ ___/ _ \/ ___/
 _>  </ /_/ / / /_/ / /  /  __/ /    
/_/|_/ .___/_/\____/_/   \___/_/     
    /_/                              
`)

func showBanner() {
	gologger.Print().Msg(banner + "\n\n")
}