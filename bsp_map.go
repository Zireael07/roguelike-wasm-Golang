package main

import (
	"log"
)


//helper
func (m *gamemap) convertBSP(rooms []BSPLeaf) []Rect {
	//convert rooms to rects
	var rects []Rect
	for _,r := range rooms {
		// #0 is the whole area, we don't need it
		//if i == 0 {
		//skip rooms that have children
		if r.has_child {
			continue
		}

		rect := Rect{pos1:position{X:r.subrect.X, Y:r.subrect.Y}, pos2:position{X:r.subrect.X+r.subrect.W, Y:r.subrect.Y+r.subrect.H}}
		rects = append(rects, rect)
	}

	log.Printf("Rects: %v", rects)
	return rects
}

//do we collide with another room?
//based on https://bfnightly.bracketproductions.com/rustbook/chapter_25.html
func (m *gamemap) isPossible(r Rect) bool {
	can_build := true
	for x := r.pos1.X; x < r.pos2.X; x++ {
		for y := r.pos1.Y; y < r.pos2.Y; y++ {
			if m.tiles[x][y].glyph == '#' {
				can_build = false
			}
		}
	}
	return can_build
}


func (m *gamemap) paintRects(rects []Rect) []Rect {
	var buildings []Rect
	for _, r := range rects {
		//check for any overlaps
		if !m.isPossible(r) {
			continue
		}
		//add to buildings list
		buildings = append(buildings, r)

		//fill with wall
		for x := r.pos1.X+1; x < r.pos2.X; x++ {
			for y := r.pos1.Y+1; y < r.pos2.Y; y++ {
				m.tiles[x][y] = &maptile{glyph: '#', fgColor: Color{255,255,255,255}, blocks_move: true, visible: false}
			}
		}

		//floor
		for x := r.pos1.X+2; x < r.pos2.X-1; x++ {
			for y := r.pos1.Y+2; y < r.pos2.Y-1; y++ {
				m.tiles[x][y] = &maptile{glyph: '.', fgColor: Color{0,255,255,255}, blocks_move: false, visible: false}
			}
		}
	}

	return buildings
}

func (m *gamemap) buildDoors(buildings [] Rect) {
	for _,r := range(buildings) {
		m.buildDoor(r)
	}
}

func (m *gamemap) buildDoor(r Rect) {
	cntr := r.center()

	choices := [4]string{"north", "south", "east", "west"}

	// copy it to avoid modifying while iterating
	sel_choices := choices[0:4]

	//TODO: check if exit leads anywhere

	if len(sel_choices) > 0 {
		wall_id := randRange(0,len(sel_choices))
		wall := sel_choices[wall_id]

		log.Printf("Wall: %v", wall)

		//dummy
		door_pos := cntr
		//buildings are one tile smaller than the enclosing rect to ensure separation
		if wall == "north" {
			door_pos = position{X:cntr.X, Y:r.pos1.Y+1}
		} else if wall == "south" {
			door_pos = position{X:cntr.X, Y:r.pos2.Y-1}
		} else if wall == "east" {
			door_pos = position{X:r.pos2.X-1, Y:cntr.Y}
		} else if wall == "west" {
			door_pos = position{X:r.pos1.X+1, Y:cntr.Y}
		}

		m.tiles[door_pos.X][door_pos.Y] = &maptile{glyph: '+', fgColor: Color{99,43,0,255}, blocks_move: false, visible: false}
	}
}