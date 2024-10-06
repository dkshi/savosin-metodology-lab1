package src

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type LCMGame struct {
	min      int
	max      int
	numCount int
}

func NewLCMGame(min, max, numCount int) *LCMGame {
	return &LCMGame{
		min:      min,
		max:      max,
		numCount: numCount,
	}
}

func (g *LCMGame) startGame() (question string, answer int) {
	fmt.Println("Find the smallest common multiple of given numbers.")

	nums := getRandomIntegerSlice(g.numCount, g.min, g.max)

	answer = lcmMultiple(nums)
	question = strings.Join(lo.Map(nums, func(t int, _ int) string {
		return strconv.Itoa(t)
	}), " ")

	return
}
