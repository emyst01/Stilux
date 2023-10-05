package main

import (
	"fmt"
	"regexp"
	"stilux/packages/read"
	"stilux/packages/variable"
	"strconv"
	"strings"

	"github.com/Knetic/govaluate"
)

var StringVariables []variable.Str
var IntegerVariables []variable.Integer
var Float1Variables []variable.Float1
var Float2Variables []variable.Float2
var BooleanVariables []variable.Boolean

func doAbsolutelyNothing(something string) {}

func main() {
	c, err := read.Read()
	if err != nil {
		print(err.Error())
		return
	}
	instructions := strings.Split(c, "-|")
	execute(instructions)
}

func execute(instructions []string) {

	for i, instruction := range instructions {
		if strings.HasPrefix(instruction, "out ") {
			instruction = strings.TrimPrefix(instruction, "out ")

			if strings.Contains(instruction, "${") {
				re := regexp.MustCompile(`\${([^}]+)}`)
				instruction = re.ReplaceAllStringFunc(instruction, func(match string) string {
					parts := strings.Split(strings.Trim(match, "${}"), "-")
					if len(parts) < 2 {
						return fmt.Sprintf("Error on instruction %d. You must specify the variable type\n", i+1)
					}
					variableType := parts[0]
					variableName := parts[1]

					switch variableType {
					case "string":
						if strVar, err := variable.FindStrByName(StringVariables, variableName); err == nil {
							return strVar.Content
						}
					case "int":
						if intVar, err := variable.FindIntegerByName(IntegerVariables, variableName); err == nil {
							return strconv.Itoa(intVar.Content)
						}
					case "float1":
						if float1Var, err := variable.FindFloat1ByName(Float1Variables, variableName); err == nil {
							return fmt.Sprintf("%f", float1Var.Content)
						}
					case "float2":
						if float2Var, err := variable.FindFloat2ByName(Float2Variables, variableName); err == nil {
							return fmt.Sprintf("%f", float2Var.Content)
						}
					case "bool":
						if boolVar, err := variable.FindBooleanByName(BooleanVariables, variableName); err == nil {
							return strconv.FormatBool(boolVar.Content)
						}
					default:
						return fmt.Sprintf("Unknown variable type: %s\n", variableType)
					}

					return fmt.Sprintf("\nVariable not found: %s-%s\n", variableType, variableName)
				})
			}

			fmt.Println(instruction)
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
		} else if strings.HasPrefix(instruction, "in ") {
			instruction = strings.TrimPrefix(instruction, "in ")
			instruction = strings.ReplaceAll(instruction, " ", "")
			var userinput string
			fmt.Scanln(&userinput)
			newVar := variable.Str{Name: instruction, Content: userinput}
			StringVariables = append(StringVariables, newVar)
		} else if strings.HasPrefix(instruction, "if ") {
			instruction = strings.TrimPrefix(instruction, "if ")
			instruction = strings.ReplaceAll(instruction, "\n", "")
			instruction = strings.ReplaceAll(instruction, "    ", "")
			condition(instruction)
		}
	}
}

func condition(ifStatement string) {

	ifStatementArr := strings.Split(ifStatement, ":")
	blocks := strings.Split(ifStatementArr[1], "-")

	if len(blocks) > 2 {
		fmt.Println("The if statement has only one false statement")
	} else if len(blocks) < 0 {
		fmt.Println("You must specify what happens when the condition is satisfied")
	}

	trueStatement := strings.Split(blocks[0], ";")
	falseStatement := strings.Split(blocks[1], ";")

	expression, err := govaluate.NewEvaluableExpression(ifStatementArr[0])
	if err != nil {
		fmt.Println("Couldn't create the expression:", err)
		return
	}

	parameters := make(map[string]interface{})

	result, err := expression.Evaluate(parameters)

	if err != nil {
		fmt.Println("Error during expression evaluation:", err)
		return
	}

	if len(blocks) == 1 {
		if result.(bool) {
			execute(trueStatement)
		}
	} else if len(blocks) == 2 {
		if result.(bool) {
			execute(trueStatement)
		} else {
			execute(falseStatement)
		}
	}
}
