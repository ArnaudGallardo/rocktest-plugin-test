package main

import (
	"fmt"

	core "github.com/ArnaudGallardo/rocktest-core"
)

var PluginMetadata core.PluginInterface = core.PluginInterface{
	Name:    "TestPlugin",
	Version: "1.0.0",
	Author:  "ArnaudGallardo",
	Methods: []string{
		"Test",
		"DumpTest",
	},
}

func Test(params map[string]interface{}, scenario *core.Scenario) error {
	fmt.Printf("Hello\n")
	return nil
}

func DumpTest(params map[string]interface{}, scenario *core.Scenario) error {
	fmt.Printf("START\n")
	mod, _ := scenario.GetContext("module")

	fmt.Printf("Variables for context %s\n", mod)

	for k, v := range scenario.GetCurrentContext() {
		fmt.Printf("  %s = %v\n", k, v)
	}
	return nil
}
