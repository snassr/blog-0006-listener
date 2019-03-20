package listener

import (
	"fmt"
	"math/rand"
)

// Listener represents a party listening to music and printing their status.
func Listener() {
	if rand.Intn(3)%2 == 0 {
		fmt.Println("Listening Happily 💃🕺")
	} else {
		fmt.Println("Listening, Not interested?...🤔")
	}
}
