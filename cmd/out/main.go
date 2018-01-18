package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := fmt.Fprint(os.Stdout, "{}")
	if err != nil {
		log.Fatalf("writing output: %s", err)
	}
}
