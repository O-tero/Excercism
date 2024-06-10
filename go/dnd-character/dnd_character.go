package dndcharacter

import (
	"math"
	"math/rand"
	"sort"
	"time"
)

// Character represents a D&D character with various abilities and hitpoints
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
	return int(math.Floor(float64(score-10) / 2))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	rand.Seed(time.Now().UnixNano())
	dice := []int{
		rand.Intn(6) + 1,
		rand.Intn(6) + 1,
		rand.Intn(6) + 1,
		rand.Intn(6) + 1,
	}
	sort.Ints(dice)
	return dice[1] + dice[2] + dice[3] // sum of the highest three rolls
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	constitution := Ability()
	return Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: constitution,
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
		Hitpoints:    10 + Modifier(constitution),
	}
}
