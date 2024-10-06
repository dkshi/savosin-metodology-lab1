package src

import (
	"math/rand"
)

// gcd функция для нахождения наибольшего общего делителя (GCD)
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// lcm функция для нахождения наименьшего общего кратного (LCM) двух чисел
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// lcmMultiple функция для нахождения LCM нескольких чисел
func lcmMultiple(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = lcm(result, numbers[i])
	}
	return result
}

// getRandomIntegerBetween вернет случайное число в диапазоне [min, max]
func getRandomIntegerBetween(min, max int) int {
	return rand.Intn(max-min) + min
}

func getRandomIntegerSlice(n, min, max int) (out []int) {
	out = make([]int, 0, n)
	for i := 0; i < n; i++ {
		out = append(out, getRandomIntegerBetween(min, max))
	}

	return
}

// geometricProgression геометрическая прогрессия
func geometricProgression(numCount, first, ratio int) []int {
	progression := make([]int, numCount)
	progression[0] = first

	for i := 1; i < numCount; i++ {
		progression[i] = progression[i-1] * ratio
	}

	return progression
}
