package functions

import types "example.com/practice/types"

type closure_slice_int func(slice []int) ([]int, []int)

func constructClosure(pPredicate *types.Predicate_int) closure_slice_int {
	// dereference pointer
	predicate := *pPredicate
	// construct closure
	return func(sl []int) (yes, no []int) {
		for _, val := range sl {
			if predicate(val) {
				yes = append(yes, val)
			} else {
				no = append(no, val)
			}
		}
		return
	}
}

var OddEven = constructClosure(&types.IsEven)
var SplitBy4 = constructClosure(&types.IsBiggerThan4)
