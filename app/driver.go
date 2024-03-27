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
	// all seasons
	executeSeasonsOfAllMonths()
}

// season of the month
func executeSeasonOfTheMonth() {
	month := rand.Intn(20)
	season, err := control_structure.Season(month)
	if err != nil {
		log.Fatalf("month : %d , error : %s\n", month, err)
	}
	fmt.Printf("month : %d , season : %s\n", month, season)
}

func executeSeasonsOfAllMonths() {
	seasons := func() map[int]string {
		result := make(map[int]string)
		for i := 1; i <= 12; i++ {
			result[i], _ = control_structure.Season(i)
		}
		return result
	}()

	fmt.Println(seasons)
}
