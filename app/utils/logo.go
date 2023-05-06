package utils

import (
	"fmt"
)

// Print the Hermes ACSII Logo
func PrintLogoWithPort(port int) {
	fmt.Printf(`
╔╗ ╔╗                   
║║ ║║                   
║╚═╝║╔══╗╔═╗╔╗╔╗╔══╗╔══╗
║╔═╗║║╔╗║║╔╝║╚╝║║╔╗║║══╣
║║ ║║║║═╣║║ ║║║║║║═╣╠══║
╚╝ ╚╝╚══╝╚╝ ╚╩╩╝╚══╝╚══╝

Serving http://localhost:%d`, port)
}
