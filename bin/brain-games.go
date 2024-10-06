package main

import (
	"log"

	"github.com/dkshi/savosin-metodology-lab1/src"
)

func main() {
	cli := src.NewCLI()
	err := cli.Greeting()
	if err != nil {
		log.Fatalf("error while greeting: %s", err.Error())
	}

	// Пайплайн игр
	gamesToPlay := []src.Game{
		src.NewLCMGame(1, 10, 3),
		src.NewGeometricProgGame(2, 3, 10),
	}

	err = cli.RunGamesPipeline(gamesToPlay)
	if err != nil {
		log.Fatalf("error running games pipeline: %s", err.Error())
	}
}
