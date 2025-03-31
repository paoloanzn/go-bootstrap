package main

import (
	"fmt"
	"log"
	"os"

	"github.com/paoloanzn/go-bootstrap/bootstrap"
	"github.com/paoloanzn/go-bootstrap/config"
	"github.com/paoloanzn/go-bootstrap/parsing"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "init":
		if len(os.Args) < 3 {
			os.Exit(1)
		}

		jsonTemplate, err := parsing.ParseTemplate(os.Args[2])
		if err != nil {
			log.Fatalf("Fatal: %v\n", err)
		}

		err = bootstrap.Bootstrap(jsonTemplate)
		if err != nil {
			log.Fatalf("Fatal: %v\n", err)
		}

	default:
		fmt.Printf("version %s\n", config.VERSION)
	}
}
