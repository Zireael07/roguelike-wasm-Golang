package main

import (
	"math/rand" //for RNG
)

//seriously, not in std?
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
  
//Golang's min and max work on float, so we provide our own for integers
func Min(x, y int32) int32 {
	if x < y {
	  return x
	}
	return y
   }
   
func Max(x, y int32) int32 {
	if x > y {
		return x
	}
	return y
}

func randRange(min, max int) int {
	//random position in range 1-19
	return rand.Intn(max-min)+min
}

// from https://github.com/golang/go/wiki/SliceTricks
// moveToFront moves needle to the front of haystack, in place if possible.
func moveToFront(needle *GameEntity, haystack []*GameEntity) []*GameEntity {
	if len(haystack) == 0 || haystack[0] == needle {
		return haystack
	}
	var prev *GameEntity
	for i, elem := range haystack {
		switch {
		case i == 0:
			haystack[0] = needle
			prev = elem
		case elem == needle:
			haystack[i] = prev
			return haystack
		default:
			haystack[i] = prev
			prev = elem
		}
	}
	return append(haystack, prev)
}