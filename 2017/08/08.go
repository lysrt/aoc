package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Operation struct {
	register     string
	action       string
	actionValue  int
	condRegister string
	condSymbol   string
	condValue    int
}

func parseLine(line string) Operation {
	parts := strings.Fields(line)

	identifier := parts[0]
	action := parts[1]
	if action != "inc" && action != "dec" {
		panic("Wrong action: " + action)
	}
	number, err := strconv.Atoi(parts[2])
	if err != nil {
		panic(err)
	}
	if parts[3] != "if" {
		print("Missing if: " + parts[3])
	}
	identifier2 := parts[4]
	symbol := parts[5]
	number2, err := strconv.Atoi(parts[6])
	if err != nil {
		panic(err)
	}

	return Operation{identifier, action, number, identifier2, symbol, number2}
}

func (op Operation) run(registers map[string]int, allTimeMax int) int {
	// Check if condition is fullfilled
	var conditionOk bool

	fmt.Printf("---- Run op: %s %s %d if %s %s %d\n", op.register, op.action, op.actionValue, op.condRegister, op.condSymbol, op.condValue)
	fmt.Printf("Before: %s = %d\n", op.register, registers[op.register])

	actual := registers[op.condRegister]
	expected := op.condValue

	if op.condSymbol == "==" {
		if actual == expected {
			conditionOk = true
		}
	} else if op.condSymbol == "!=" {
		if actual != expected {
			conditionOk = true
		}
	} else if op.condSymbol == "<" {
		if actual < expected {
			conditionOk = true
		}
	} else if op.condSymbol == "<=" {
		if actual <= expected {
			conditionOk = true
		}
	} else if op.condSymbol == ">" {
		if actual > expected {
			conditionOk = true
		}
	} else if op.condSymbol == ">=" {
		if actual >= expected {
			conditionOk = true
		}
	} else {
		panic("Unknown symbol: " + op.condSymbol)
	}

	if !conditionOk {
		fmt.Println("DO NOT RUN.")
		return allTimeMax
	}

	if op.action == "inc" {
		registers[op.register] += op.actionValue
	} else {
		registers[op.register] -= op.actionValue
	}

	fmt.Printf("Condition OK -> %s = %d\n", op.register, registers[op.register])

	if registers[op.register] > allTimeMax {
		allTimeMax = registers[op.register]
	}
	return allTimeMax
}

func readAllOperations() []Operation {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var operations []Operation

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		line := s.Text()
		operation := parseLine(line)
		operations = append(operations, operation)
	}

	return operations
}

func main() {
	operations := readAllOperations()

	// Initialize registers
	registers := make(map[string]int)
	for _, op := range operations {
		if _, ok := registers[op.register]; !ok {
			registers[op.register] = 0
		}
	}

	// Run all operations
	allTimeMax := math.MinInt32
	for _, op := range operations[0:10] {
		allTimeMax = op.run(registers, allTimeMax)
	}

	// Display the largest value in the registers
	max := 0
	for _, value := range registers {
		if value > max {
			max = value
		}
	}
	fmt.Println("Max: ", max)
	fmt.Println("All time max: ", allTimeMax)
}
