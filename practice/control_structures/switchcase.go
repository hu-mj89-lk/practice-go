package control_structures

import (
	"errors"
	"math/rand"
)

var err_templates = []string{
	"Oppsy daizies !! You've given an incorrect month",
	"You've reached dead end",
	"Season unknown",
}

func Season(month int) (string, error) {
	switch month {
	case 1, 2, 12:
		return "Winter", nil
	case 3, 4, 5:
		return "Spring", nil
	case 6, 7, 8:
		return "Summer", nil
	case 9, 10, 11:
		return "Autumn", nil
	default:
		return "", errors.New(err_templates[rand.Intn(3)])
	}
}
