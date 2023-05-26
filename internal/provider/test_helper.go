package provider

import (
	"fmt"
	"os"
)

// HelperTestAccTFExampleConfig is intended for us to more easily pass in just the
// corresponding resource folder and file
// Example: HelperTestAccTFExampleConfig("waypoint_project/project.tf")
func HelperTestAccTFExampleConfig(filename string) string {
	fileContent, err := os.ReadFile("../../examples/resources/" + filename)
	if err != nil {
		fmt.Print(err)
	}

	return string(fileContent)
}
