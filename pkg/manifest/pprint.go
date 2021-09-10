package manifest

import (
	"fmt"
	"github.com/deissh/rf-cli/pkg/extension"
)

func PrettyPrint(ext *extension.Extension) {
	fmt.Println("ID:", ext.ID)
	fmt.Println("Name:", ext.Name)
	fmt.Println("Short description:", ext.ShortDescription)
	fmt.Println("Description:", ext.Description)
	fmt.Println("Email:", ext.Email)
	fmt.Println("BaseURL:", ext.BaseURL)
	fmt.Println("Extension user:")
	fmt.Println("    Username:", ext.User.Username)

	fmt.Println("\nCommands:")
	for i, command := range ext.Commands {
		fmt.Printf(" %d. %s\n", i+1, command.Name)
		fmt.Println("    Description:", command.Description)
		fmt.Println("    Group:", command.Group)
		fmt.Printf("    Type: %+v\n", command.Type)

		fmt.Println("\n    ShowRules")
		for j, rule := range command.ShowRules {
			fmt.Printf("     %d. %+v\n", j+1, rule)
		}
		fmt.Println()
	}

	fmt.Println("\nRequired Types:")
	for i, reqType := range ext.RequiredTypes {
		fmt.Printf(" %d. %s\n", i+1, reqType.Name)

		fmt.Println("    Properties")
		for j, prop := range reqType.Properties {
			fmt.Printf("     %d. %s (%s / %s)\n", j+1, prop.Name, prop.Category, prop.Argument)
		}
		fmt.Println()
	}
}
