package main

import (
	"fmt"
	"log"
	"math/rand"

	"example.com/practice/control_structures"
	"example.com/practice/functions"
	types "example.com/practice/types"
)

func main() {
	fmt.Println()
	
	// season of a month
	fmt.Println("##### executeSeasonOfTheMonth #####")
	executeSeasonOfTheMonth()

	// all seasons
	fmt.Println("##### executeSeasonsOfAllMonths #####")
	executeSeasonsOfAllMonths()

	// try closures
	fmt.Println("##### executeClosures #####")
	executeClosures()

	// try closures
	fmt.Println("##### createSliceOfCustomType #####")
	createSliceOfCustomType()

}

// season of the month
func executeSeasonOfTheMonth() {
	defer fmt.Println()

	month := rand.Intn(20)
	season, err := control_structures.Season(month)
	if err != nil {
		log.Printf("month : %d , error : %s\n", month, err)
	} else {
		fmt.Printf("month : %d , season : %s\n", month, season)
	}
}

func executeSeasonsOfAllMonths() {
	defer fmt.Println()
	
	seasons := func() map[int]string {
		result := make(map[int]string)
		for i := 1; i <= 12; i++ {
			result[i], _ = control_structures.Season(i)
		}
		return result
	}()

	fmt.Println(seasons)
}

func executeClosures() {
	defer fmt.Println()
	
	sl := []int{1, 2, 3, 4, 5, 6, 7}

	even, odd := functions.OddEven(sl)
	fmt.Println("Even : ", even, " | Odd : ", odd)

	bigger, smaller := functions.SplitBy4(sl)
	fmt.Println("Bigger : ", bigger, " | Smaller : ", smaller)
}

func createSliceOfCustomType() {
	defer fmt.Println()
	
	var slStudents []types.Student = []types.Student{
		2: {Name: "Jitesh", Age: 35},
		4: {Name: "Robin", Age: 34},
	}
	fmt.Print(slStudents)
}
