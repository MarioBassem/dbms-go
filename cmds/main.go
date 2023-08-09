package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	StatementInsert = iota
	StatementSelect
)

type Statement struct {
	typ int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}

		line := scanner.Text()
		if line == "" {
			continue
		}

		if line[0] == '.' {
			exit, err := doMetaCommand(line)
			if err != nil {
				fmt.Printf("meta error: %s\n", err.Error())
			}

			if exit {
				os.Exit(0)
			}

			continue
		}

		statement, err := prepareStatement(line)
		if err != nil {
			fmt.Printf("sql error: %s\n", err.Error())
			continue
		}

		if err := executeStatement(statement); err != nil {
			fmt.Printf("exec error: %s\n", err.Error())
		}
	}

}

func doMetaCommand(command string) (exit bool, err error) {
	switch command {
	case ".exit":
		return true, nil
	default:
		return false, fmt.Errorf("Unrecognized command %s", command)
	}
}

func prepareStatement(command string) (Statement, error) {
	switch command {
	case "insert":
		return Statement{typ: StatementInsert}, nil
	case "select":
		return Statement{typ: StatementSelect}, nil
	default:
		return Statement{}, fmt.Errorf("unrecognized statement %s", command)
	}
}

func executeStatement(statement Statement) error {
	switch statement.typ {
	case StatementInsert:
		fmt.Printf("this is where an insert is handled")
	case StatementSelect:
		fmt.Print("this is where a select is handled")
	}

	return nil
}
