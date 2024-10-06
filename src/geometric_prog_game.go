package src

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type GeometricProgGame struct {
	first    int
	ratio    int
	numCount int
}

func NewGeometricProgGame(first, ratio, numCount int) *GeometricProgGame {
	return &GeometricProgGame{
		first:    first,
		ratio:    ratio,
		numCount: numCount,
	}
}

func (g *GeometricProgGame) startGame() (question string, answer int) {
	fmt.Println("What number is missing in the progression")

	nums := geometricProgression(g.numCount, g.first, g.ratio)
	randomMissing := getRandomIntegerBetween(0, g.numCount-1)
	answer = nums[randomMissing]

	question = strings.Join(lo.Map(nums, func(t int, _ int) string {
		if t == answer {
			return "_"
		}
		return strconv.Itoa(t)
	}), " ")

	return
}
