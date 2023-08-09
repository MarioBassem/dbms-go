package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}

		switch scanner.Text() {
		case ".exit":
			os.Exit(0)
		default:
			fmt.Printf("Unrecognized command %s.\n", scanner.Text())
		}
	}

}
