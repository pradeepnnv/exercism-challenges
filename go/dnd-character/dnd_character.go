package dndcharacter

import (
	"math"
	"math/rand"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	di := float64(score - 10)
	fr := float64(di / 2.0) // when dividing variables, both of them have to be float for answer to be in float
	rn := math.Floor(fr)
	return int(rn)
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	for {
		v := rand.Intn(19)
		if v >= 3 && v <= 18 {
			return v
		}
	}
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	a := []int{} // combo of make and append will cause zeros to be available in the beginning. Hence switched to empty slice.
	for i := 0; i < 6; i++ {
		a = append(a, Ability(), Ability(), Ability())
	}
	c := Character{a[0], a[1], a[2], a[3], a[4], a[5], 10 + Modifier(a[2])}
	return c
}
