package utils

import (
	"math/rand"
)

type RandomSequenceGenerator struct {
	numRunes int
	runes    []rune
}

func NewRandomSequenceGenerator() *RandomSequenceGenerator {
	runes := []rune(`abcdefghijklmnopqrstuvwxyz1234567890@#$^&*()_-=+`)
	return &RandomSequenceGenerator{
		runes:    runes,
		numRunes: len(runes),
	}
}

func (g *RandomSequenceGenerator) Next(keyLen int) string {
	key := make([]rune, g.numRunes)

	for i := range g.runes {
		key[i] = g.runes[rand.Intn(g.numRunes)]
	}

	return string(key)
}
