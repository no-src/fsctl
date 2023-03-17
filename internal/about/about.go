package about

import (
	"github.com/no-src/fsctl/internal/version"
	"github.com/no-src/log"
)

const logo = `
 ________ ________  ________ _________  ___          
|\  _____\\   ____\|\   ____\\___   ___\\  \         
\ \  \__/\ \  \___|\ \  \___\|___ \  \_\ \  \        
 \ \   __\\ \_____  \ \  \       \ \  \ \ \  \       
  \ \  \_| \|____|\  \ \  \____   \ \  \ \ \  \____  
   \ \__\    ____\_\  \ \_______\  \ \__\ \ \_______\
    \|__|   |\_________\|_______|   \|__|  \|_______|
            \|_________|                             

`
const (
	openSourceUrl    = "https://github.com/no-src/fsctl"
	documentationUrl = "https://pkg.go.dev/github.com/no-src/fsctl@" + version.VERSION
	releaseUrl       = "https://github.com/no-src/fsctl/releases"
	dockerImageUrl   = "https://hub.docker.com/r/nosrc/fsctl"
)

// PrintAbout print the program logo and basic info
func PrintAbout() {
	log.Log(logo)
	log.Log("The fsctl is a configuration-based file operation and validation tool")
	log.Log("Open source repository at: <%s>", openSourceUrl)
	log.Log("Download the latest version at: <%s>", releaseUrl)
	log.Log("The docker image repository address at: <%s>", dockerImageUrl)
	log.Log("Full documentation at: <%s>", documentationUrl)
}
