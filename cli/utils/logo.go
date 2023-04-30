package utils

import (
	"fmt"
)

// Print the Hermes ACSII Logo
func PrintLogoWithPort(port int) {
	fmt.Println("Hermes")
	fmt.Println("Serving port:", port)
}
