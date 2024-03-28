package types

type Predicate_int func(input int) bool

// function as a value
var IsEven Predicate_int = func(input int) bool {
	return input%2 == 0
}

// function as a value
var IsBiggerThan4 Predicate_int = func(input int) bool {
	return input > 4
}
