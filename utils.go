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