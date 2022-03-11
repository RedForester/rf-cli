package view

import (
	"fmt"
	"github.com/deissh/rf-cli/pkg/rf"
	"github.com/jaytaylor/html2text"
	"io"
)

type Extension struct {
	Data   *rf.Extension
	Writer io.Writer
}

func (e Extension) Render() error {
	ext := e.Data

	fmt.Println("ID:", ext.ID)
	fmt.Println("Name:", ext.Name)
	fmt.Println("Short description:", ext.ShortDescription)

	e.renderDescription()

	fmt.Println("Email:", ext.Email)
	if ext.BaseURL != nil {
		fmt.Println("BaseURL:", *ext.BaseURL)
	}
	fmt.Println("Extension user:")
	fmt.Println("    Username:", ext.User.Username)

	e.renderCommands()
	e.renderTypes()

	return nil
}

func (e Extension) renderDescription() {
	fmt.Println("Description:")
	text, err := html2text.FromString(e.Data.Description, html2text.Options{PrettyTables: true})
	if err != nil {
		// todo: show warn
		fmt.Println(e.Data.Description)
	}

	fmt.Println(text)
}

func (e Extension) renderCommands() {
	fmt.Println("\nCommands:")
	for i, command := range e.Data.Commands {
		fmt.Printf(" %d. %s\n", i+1, command.Name)
		fmt.Println("    Description:", command.Description)
		if command.Group != nil {
			fmt.Println("    Group:", command.Group)
		}
		switch {
		case command.Type.URL != nil:
			fmt.Printf("    Type: url=%v", *command.Type.URL)
		case command.Type.Action != nil:
			fmt.Printf("    Type: action=%v", *command.Type.Action)
		}

		fmt.Println("\n    ShowRules")
		for j, rule := range command.ShowRules {
			fmt.Printf("     %d. ", j+1)
			switch {
			case rule.Root:
				fmt.Println("Only on root node")
			case rule.AllNodes:
				fmt.Println("All nodes")
			case rule.SelfType != nil:
				fmt.Printf("On nodes of type \"%v\"\n", *rule.SelfType)
			case rule.DescendantOfType != nil:
				fmt.Printf("Node descendants of type \"%v\"\n", *rule.DescendantOfType)
			}
		}
	}
}

func (e Extension) renderTypes() {
	fmt.Println("\nRequired Types:")
	for i, reqType := range e.Data.RequiredTypes {
		fmt.Printf(" %d. \"%s\"\n", i+1, reqType.Name)
		if len(reqType.Properties) == 0 {
			continue
		}
		fmt.Println("    Properties:")
		for j, prop := range reqType.Properties {
			fmt.Printf("     %d. \"%s\" (%s / %s)\n", j+1, prop.Name, prop.Category, prop.Argument)
		}
	}
}
