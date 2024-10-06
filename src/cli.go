package src

import "fmt"

type Game interface {
	startGame() (question string, answer int)
}

type CLI struct {
	username string
}

func NewCLI() *CLI {
	return &CLI{}
}

func (c *CLI) Greeting() (err error) {
	fmt.Println("Welcome to the Brain Games!")
	fmt.Print("May I have your Name? ")
	_, err = fmt.Scan(&c.username)
	if err != nil {
		return err
	}

	fmt.Printf("Hello, %s", c.username)

	return nil
}

func (c *CLI) RunGamesPipeline(games []Game) (err error) {
	for _, game := range games {
		err = c.playGame(game)
		if err != nil {
			return err
		}
	}
	fmt.Printf("Congratulations, %s!", c.username)
	return nil
}

func (c *CLI) playGame(game Game) (err error) {
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
		fmt.Printf("Let's try again, %s\n", c.username)
		return c.playGame(game)
	}

	fmt.Println("Correct!")
	return nil
}
