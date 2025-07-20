// cmd/artgenerator/main.go
package main

import (
	"flag"
	"log"
	"os"

	"artgenerator/internal/artgenerator"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := artgenerator.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
