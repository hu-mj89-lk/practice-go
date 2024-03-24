package main

import (
	"fmt"
	"log"
	"math/rand"

	"example.com/practice/control_structure"
)

func main() {
	// season of a month
	executeSeasonOfTheMonth()

}

// season of the month
func executeSeasonOfTheMonth() {
	month := rand.Intn(20)
	season, err := control_structure.Season(month)
	if err != nil {
		log.Fatalf("month : %d , error : %s", month, err)
	}
	fmt.Printf("month : %d , season : %s", month, season)
}
