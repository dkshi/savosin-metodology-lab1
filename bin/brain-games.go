package main

import (
	"log"

	"github.com/dkshi/savosin-metodology-lab1/src"
)

func main() {
	ge := src.NewGameEngine()
	err := ge.Greeting()
	if err != nil {
		log.Fatalf("error while greeting: %s", err.Error())
	}

	// Пайплайн игр
	gamePipeline := src.GamePipeline{
		src.NewLCMGame(1, 10, 3),
		src.NewGeometricProgGame(2, 3, 10),
	}

	err = ge.RunGamesPipeline(gamePipeline)
	if err != nil {
		log.Fatalf("error running games pipeline: %s", err.Error())
	}
}
