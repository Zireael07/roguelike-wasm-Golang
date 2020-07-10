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


func (m *gamemap) paintRects(rects []Rect) {
	for _, r := range rects {
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
}