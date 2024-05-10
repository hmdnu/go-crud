package utils

import (
	"math/rand"
)

func GenerateId() int {
	return rand.New(rand.NewSource(5)).Int()
}
