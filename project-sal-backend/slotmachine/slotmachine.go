package slotmachine

import (
	"math/rand"

	"github.com/davecusatis/project-sal-backend/project-sal-backend/models"
)

const MaxValue = 8

func GenerateRandomScore() models.Score {
	num1 := (rand.Intn(MaxValue) + 1)
	num2 := (rand.Intn(MaxValue) + 1) * 10
	num3 := (rand.Intn(MaxValue) + 1) * 100
	return models.Score{
		Score:     num1 + num2 + num3,
		ChannelID: "RIGdavethecust",
		UserName:  "davethecust",
	}
}
