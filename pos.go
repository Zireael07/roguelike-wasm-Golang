package main

import (
	//"log"
	"fmt"
)

type position struct {
	X int
	Y int
}

func (pos position) Distance(to position) int {
	deltaX := Abs(to.X - pos.X)
	deltaY := Abs(to.Y - pos.Y)
	if deltaX > deltaY {
		return deltaX
	}
	return deltaY
}

func (pos position) sub(other position) position {
	return position{pos.X - other.X, pos.Y - other.Y}
}

//direction enum taken from Boohu by anaseto
type direction int

const (
	NoDir direction = iota
	E
	ENE
	NE
	NNE
	N
	NNW
	NW
	WNW
	W
	WSW
	SW
	SSW
	S
	SSE
	SE
	ESE
)

func (pos position) Dir(from position) direction {
	deltaX := Abs(pos.X - from.X)
	deltaY := Abs(pos.Y - from.Y)
	switch {
	case pos.X > from.X && pos.Y == from.Y:
		return E
	case pos.X > from.X && pos.Y < from.Y:
		switch {
		case deltaX > deltaY:
			return ENE
		case deltaX == deltaY:
			return NE
		default:
			return NNE
		}
	case pos.X == from.X && pos.Y < from.Y:
		return N
	case pos.X < from.X && pos.Y < from.Y:
		switch {
		case deltaY > deltaX:
			return NNW
		case deltaX == deltaY:
			return NW
		default:
			return WNW
		}
	case pos.X < from.X && pos.Y == from.Y:
		return W
	case pos.X < from.X && pos.Y > from.Y:
		switch {
		case deltaX > deltaY:
			return WSW
		case deltaX == deltaY:
			return SW
		default:
			return SSW
		}
	case pos.X == from.X && pos.Y > from.Y:
		return S
	case pos.X > from.X && pos.Y > from.Y:
		switch {
		case deltaY > deltaX:
			return SSE
		case deltaX == deltaY:
			return SE
		default:
			return ESE
		}
	default:
		panic(fmt.Sprintf("internal error: invalid position:%+v-%+v", pos, from))
	}
}

//for debugging/display
func (dir direction) String() string {
    // ... operator counts how many items in the array
    names := [...]string{
		"NoDir",
        "E",
		"ENE",
		"NE",
		"NNE",
		"N",
		"NNW",
		"NW",
		"WNW",
		"W",
		"WSW",
		"SW",
		"SSW",
		"S",
		"SSE",
		"SE",
		"ESE"}
   
    // return the name of a constant from the names array above.
    return names[dir]
}