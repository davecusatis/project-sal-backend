package slotmachine

import "math/rand"

const MaxValue = 8

func GenerateRandomScore() int {
	num1 := (rand.Intn(MaxValue) + 1)
	num2 := (rand.Intn(MaxValue) + 1) * 10
	num3 := (rand.Intn(MaxValue) + 1) * 100
	return num1 + num2 + num3
}
