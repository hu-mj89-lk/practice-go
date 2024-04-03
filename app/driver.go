package main

import (
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"runtime"

	"example.com/practice/control_structures"
	"example.com/practice/functions"
	types "example.com/practice/types"
)

func execute(f func()) {
	defer fmt.Println()

	valueF := reflect.ValueOf(f)
	funcF := runtime.FuncForPC(valueF.Pointer())
	nameF := funcF.Name()
	fmt.Printf("##### %s #####\n", nameF)

	f()
}

func main() {
	fmt.Println()

	// season of a month
	execute(seasonOfTheMonth)

	// all seasons
	execute(seasonsOfAllMonths)

	// play with closures
	execute(playWithClosures)

	// play with custom types
	execute(playWithCustomType)

	// play with arrays
	execute(playWithArrays)

	// play with slices
	execute(playWithSlices)
}

// season of the month
func seasonOfTheMonth() {
	month := rand.Intn(20)
	season, err := control_structures.Season(month)
	if err != nil {
		log.Printf("month : %d , error : %s\n", month, err)
	} else {
		fmt.Printf("month : %d , season : %s\n", month, season)
	}
}

func seasonsOfAllMonths() {
	seasons := func() map[int]string {
		result := make(map[int]string)
		for i := 1; i <= 12; i++ {
			result[i], _ = control_structures.Season(i)
		}
		return result
	}()

	fmt.Println(seasons)
}

func playWithClosures() {
	sl := []int{1, 2, 3, 4, 5, 6, 7}

	even, odd := functions.OddEven(sl)
	fmt.Println("Even : ", even, " | Odd : ", odd)

	bigger, smaller := functions.SplitBy4(sl)
	fmt.Println("Bigger : ", bigger, " | Smaller : ", smaller)
}

func playWithCustomType() {
	stud1 := types.Student{Name: "Jitesh", Age: 35}
	fmt.Println(stud1)
}

func playWithArrays() {
	arr1 := [5]int{1, 2, 3}

	fmt.Printf("arr1 | type : %t, value : %v, length : %d, capacity : %d\n", arr1, arr1, len(arr1), cap(arr1))
}

func playWithSlices() {
	var sl1 []int
	fmt.Printf("sl1 | type : %t, value : %v, length : %d, capacity : %d\n", sl1, sl1, len(sl1), cap(sl1))
}
