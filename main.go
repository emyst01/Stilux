package main

import (
	"fmt"
	"regexp"
	"stilux/packages/read"
	"stilux/packages/variable"
	"strconv"
	"strings"
)

var StringVariables []variable.Str
var IntegerVariables []variable.Integer
var Float1Variables []variable.Float1
var Float2Variables []variable.Float2
var BooleanVariables []variable.Boolean

func doAbsolutelyNothing(something string) {}

func print(msg string) {
	fmt.Println(msg)
}

func main() {
	c, err := read.Read()
	if err != nil {
		print(err.Error())
		return
	}

	instructions := strings.Split(c, "-|")

	for i, instruction := range instructions {
		if strings.HasPrefix(instruction, "out ") {
			instruction = strings.TrimPrefix(instruction, "out ")
			print(instruction)
		} else if strings.HasPrefix(instruction, "var ") {
			instruction = strings.TrimPrefix(instruction, "var ")
			re := regexp.MustCompile(`\s+`)
			parts := re.Split(instruction, -1)
			args := []string{parts[0], parts[1], strings.Join(parts[2:], " ")}
			args[0] = strings.ToLower(args[0])
			args[1] = strings.ToLower(args[1])

			switch args[1] {
			case "string":
				newVar := variable.Str{Name: args[0], Content: args[2]}
				StringVariables = append(StringVariables, newVar)
			case "int", "integer":
				content, err := strconv.Atoi(args[2])
				if err != nil {
					fmt.Printf("Error at instruction %d, The declared variable type doesn't match with the content type %s: %s\n", i+1, args[0], err)
					continue
				}
				newVar := variable.Integer{Name: args[0], Content: content}
				IntegerVariables = append(IntegerVariables, newVar)
			case "float1", "float32", "float":
				content, err := strconv.ParseFloat(args[2], 32)
				if err != nil {
					fmt.Printf("Error at instruction %d, The declared variable type doesn't match with the content type %s: %s\n", i+1, args[0], err)
					continue
				}
				newVar := variable.Float1{Name: args[0], Content: float32(content)}
				Float1Variables = append(Float1Variables, newVar)
			case "float2", "float64":
				content, err := strconv.ParseFloat(args[2], 64)
				if err != nil {
					fmt.Printf("Error at instruction %d, The declared variable type doesn't match with the content type %s: %s\n", i+1, args[0], err)
					continue
				}
				newVar := variable.Float2{Name: args[0], Content: content}
				Float2Variables = append(Float2Variables, newVar)
			case "bool", "boolean":
				content, err := strconv.ParseBool(args[2])
				if err != nil {
					fmt.Printf("Error at instruction %d, The declared variable type doesn't match with the content type %s: %s\n", i+1, args[0], err)
					continue
				}
				newVar := variable.Boolean{Name: args[0], Content: content}
				BooleanVariables = append(BooleanVariables, newVar)
			default:
				fmt.Printf("Error at instruction %d, This variable type doesn't exist %s: %s\n", i+1, args[0], args[1])
			}
		}
	}
}
