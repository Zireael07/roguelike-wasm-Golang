//based on https://github.com/SolarLune/dngn/blob/master/dngn.go

package main

import (
	"math/rand"
	"log"
)

type subrect struct {
	X, Y, W, H int //easier for BSP to work on W, H
}

//helper functions
func subMinSize(subroom subrect) int {
	if subroom.W < subroom.H {
		return subroom.W
	}
	return subroom.H
}

func subSplit(parent subrect, vertical bool, min_size int) (subrect, subrect, bool) {

	splitPercentage := 0.2 + rand.Float32()*0.6

	if vertical {

		splitCX := int(float32(parent.W) * splitPercentage)
		splitCX2 := parent.W - splitCX
		a, b := subrect{parent.X, parent.Y, splitCX, parent.H}, subrect{parent.X + splitCX, parent.Y, splitCX2, parent.H}

		if subMinSize(a) <= min_size || subMinSize(b) <= min_size {
			return a, b, false
		}

		// // Line is attempting to start on a door
		// if doorValue != wallValue && doorValue != 0 && (room.Get(parent.X+splitCX, parent.Y) == doorValue || room.Get(parent.X+splitCX, parent.Y+parent.H) == doorValue) {
		// 	return a, b, false
		// }

		//room.DrawLine(parent.X+splitCX, parent.Y+1, parent.X+splitCX, parent.Y+parent.H-1, wallValue, 1, false)

		// Place door
		// for i := 0; i < 100; i++ {
		// 	ry := parent.Y + 1 + rand.Intn(parent.H-1)
		// 	if room.Get(parent.X+splitCX-1, ry) == wallValue || room.Get(parent.X+splitCX+1, ry) == wallValue {
		// 		continue
		// 	}
		// 	room.Set(parent.X+splitCX, ry, doorValue)
		// 	break
		// }

		return a, b, true
	}

	splitCY := int(float32(parent.H) * splitPercentage)
	splitCY2 := parent.H - splitCY
	a, b := subrect{parent.X, parent.Y, parent.W, splitCY}, subrect{parent.X, parent.Y + splitCY, parent.W, splitCY2}

	if subMinSize(a) <= min_size || subMinSize(b) <= min_size {
		return a, b, false
	}

	// // Line is attempting to start on a door
	// if doorValue != wallValue && doorValue != 0 && (room.Get(parent.X, parent.Y+splitCY) == doorValue || room.Get(parent.X+parent.W, parent.Y+splitCY) == doorValue) {
	// 	return a, b, false
	// }

	//room.DrawLine(parent.X+1, parent.Y+splitCY, parent.X+parent.W-1, parent.Y+splitCY, wallValue, 1, false)

	// Create doors somewhere in the lines
	// for i := 0; i < 100; i++ {
	// 	rx := parent.X + 1 + rand.Intn(parent.W-1)
	// 	if room.Get(rx, parent.Y+splitCY-1) == wallValue || room.Get(rx, parent.Y+splitCY+1) == wallValue {
	// 		continue
	// 	}
	// 	room.Set(rx, parent.Y+splitCY, doorValue)
	// 	break
	// }

	return a, b, true
}

//meat of the generator
func (m *gamemap) GenerateBSP(numSplits int) []subrect {
	rect := Rect{pos1:position{X:0, Y:0}, pos2:position{m.width, m.height}}
	
	//if we have a submap act only in submap
	if len(m.submaps) > 0 {
		rect = m.submaps[0]
	}

	rooms := []subrect{subrect{rect.pos1.X, rect.pos1.Y, rect.pos2.X-rect.pos1.X, rect.pos2.Y-rect.pos1.Y}}

	attemptCount := 10
	doneSplits := 0

	//while equivalent
	for attemptCount > 0 {
		//we're done, break the loop
		if doneSplits == numSplits {
			break
		}

		//random room choice
		splitChoice := rooms[rand.Intn(len(rooms))]

		// random choise
		split_vert := randBool()
		//do we split horizontally or vertically?
		//if width is 25% bigger than height, split vertically, and the converse for horizontal
		if float32(splitChoice.W/splitChoice.H) > 1.25 {
			split_vert = true
		} else {
			split_vert = false
		}

		// Do the split
		a, b, success := subSplit(splitChoice, split_vert, 5)

		//decrement attempts
		attemptCount--

		if !success {
			continue
		} else {
			//add the new room
			rooms = append(rooms, a, b)

			for i, r := range rooms {
				if r == splitChoice {
					rooms = append(rooms[:i], rooms[i+1:]...)
					break
				}
			}

			//count number of splits actually done
			doneSplits++
		}
	}


	log.Printf("BSP rooms: %v ", rooms)
	//return rooms list
	return rooms
}

