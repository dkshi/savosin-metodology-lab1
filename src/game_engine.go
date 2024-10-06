package src

import "fmt"

type GamePipeline []Game

type Game interface {
	startGame() (question string, answer int)
}

type GameEngine struct {
	username string
}

func NewGameEngine() *GameEngine {
	return &GameEngine{}
}

func (e *GameEngine) Greeting() (err error) {
	fmt.Println("Welcome to the Brain Games!")
	fmt.Print("May I have your Name? ")
	_, err = fmt.Scan(&e.username)
	if err != nil {
		return err
	}

	fmt.Printf("Hello, %s", e.username)

	return nil
}

func (e *GameEngine) RunGamesPipeline(games GamePipeline) (err error) {
	for _, game := range games {
		err = e.playGame(game)
		if err != nil {
			return err
		}
	}
	fmt.Printf("Congratulations, %s!", e.username)
	return nil
}

func (e *GameEngine) playGame(game Game) (err error) {
	question, answer := game.startGame()

	var userAnswer int
	fmt.Printf("Question: %v\n", question)
	fmt.Print("Your answer: ")
	_, err = fmt.Scan(&userAnswer)
	if err != nil {
		return err
	}

	if userAnswer != answer {
		fmt.Printf("'%d' is wrong answer ;(. Correct answer was '%d'\n", userAnswer, answer)
		fmt.Printf("Let's try again, %s\n", e.username)
		return e.playGame(game)
	}

	fmt.Println("Correct!")
	return nil
}
