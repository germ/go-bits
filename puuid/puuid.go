// Package puuid returns 25 character ID in the style of UUIDs, 1 / (26^25) chance of collision. Good enough.
package puuid

import (
	"math/rand"
	"time"
)

func Generate() (string) {
	var uuid string
	rand.Seed(time.Now().UnixNano())

	//Look at that nasty ASCII, I need to learn to UTF proper
	//TODO: Fix this shit
	for index := 1; index <= 25; index++ {
		if index % 5 == 0 && index != 25 {
			uuid += "-"
		}
		uuid += string('A' + (rand.Int() % 26))
	}

	return uuid
}
