package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"strings"
)

type GeneratorConfig struct {
	ToolName    string
	Description string
	Fields      []Field
}

type Field struct {
	Name        string
	Type        string
	Description string
}

func main() {
	fmt.Println("Welcome to the Responsive CLI Tool Generator!")
	fmt.Println("Please answer the following questions to generate your tool:")

	config := GeneratorConfig{}

	// Get tool name
	prompt := &promptui.Prompt{
		Label: "Enter the name of your tool",
	}
	name, err := prompt.Run()
	if err != nil {
		fmt.Printf(" Prompt failed %v\n", err)
		return
	}
	config.ToolName = name

	// Get tool description
	prompt = &promptui.Prompt{
		Label: "Enter a brief description of your tool",
	}
	description, err := prompt.Run()
	if err != nil {
		fmt.Printf(" Prompt failed %v\n", err)
		return
	}
	config.Description = description

	// Get fields
	prompt = &promptui.Prompt{
		Label: "Enter the number of fields for your tool (0 to finish)",
	}
	for {
		numFields, err := prompt.Run()
		if err != nil {
			fmt.Printf(" Prompt failed %v\n", err)
			return
		}
		num, err := strconv.Atoi(numFields)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}
		if num == 0 {
			break
		}
		for i := 0; i < num; i++ {
			field := Field{}
			prompt = &promptui.Prompt{
				Label: fmt.Sprintf("Enter the name of field %d", i+1),
			}
			name, err := prompt.Run()
			if err != nil {
				fmt.Printf(" Prompt failed %v\n", err)
				return
			}
			field.Name = name

			prompt = &promptui.Prompt{
				Label: fmt.Sprintf("Enter the type of field %s (string, int, etc.)", field.Name),
			}
		 tipo, err := prompt.Run()
			if err != nil {
				fmt.Printf(" Prompt failed %v\n", err)
				return
			}
			field.Type = tipo

			prompt = &promptui.Prompt{
				Label: fmt.Sprintf("Enter a brief description of field %s", field.Name),
			}
			desc, err := prompt.Run()
			if err != nil {
				fmt.Printf(" Prompt failed %v\n", err)
				return
			}
			field.Description = desc

			config.Fields = append(config.Fields, field)
		}
	}

	fmt.Println("Generating your tool...")
	generateTool(config)
}

func generateTool(config GeneratorConfig) {
	fmt.Printf("package main\n\n")
	fmt.Printf("import \"fmt\"\n\n")
	fmt.Printf("type %s struct {\n", strings.Title(config.ToolName))
	for _, field := range config.Fields {
		fmt.Printf("\t%s %s // %s\n", strings.Title(field.Name), field.Type, field.Description)
	}
	fmt.Printf("}\n\n")
	fmt.Printf("func main() {\n")
	fmt.Printf("\ttool := &%s{\n", strings.Title(config.ToolName))
	for _, field := range config.Fields {
		fmt.Printf("\t\t%s: \"\",\n", strings.Title(field.Name))
	}
	fmt.Printf("\t}\n")
	fmt.Printf("\tfmt.Printf(\"%%+v\\n\", tool)\n")
	fmt.Printf("}\n")
}