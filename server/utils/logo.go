package utils

import (
	"fmt"
	"os"
)

// Print the Hermes ACSII Logo
func PrintLogoWithPort(port int) error {
	// Read the logo file
	var logo, err = os.ReadFile("/static/logo_ascii.txt")
	if err != nil {
		return err
	}

	// Print the logo and port
	fmt.Println(string(logo))
	fmt.Println("Serving port:", port)

	// Return no error
	return nil
}
