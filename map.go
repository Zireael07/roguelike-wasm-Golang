package main

import (
	"sort"
	//"log"
)

type maptile struct {
	glyph rune
	fgColor Color
	blocks_move bool
	//FOV
	explored bool
	visible bool
}

type freetile struct {
	tile *maptile
	pos position
}

//TODO: put in map_common.go or some such?
type Rect struct {
	pos1 position
	pos2 position
}


//because 'map' in Go is a data structure...
type gamemap struct {
	width int
	height int
	tiles [][]*maptile //2d array/slice of tiles, nothing unusual
	freetiles []*freetile //list of all free tiles
	submaps []Rect
}

func (t *maptile) IsWall() bool {
	if t.blocks_move {
		return true
	}

	return false
}

var (
	MapTileNum = 0
)


func (m *gamemap) InitMap() {
	// Initialize a 2d array that will represent the current game map (of dimensions Width x Height)
	m.tiles = make([][]*maptile, m.width+1)
	for i := range m.tiles {
		m.tiles[i] = make([]*maptile, m.height+1)
	}
	MapTileNum = m.width*m.height;
}


type distpos struct {
	pos position
	dist int
}

func (m *gamemap) findGridInRange(dist int, pos position) []distpos {
	var coords []distpos

	for x := pos.X; x <= pos.X + dist; x++ {
		for y := pos.Y; y <= pos.Y + dist; y++{
			if x > 0 && x <= m.width && y > 0 && y <= m.height {
				cand := position{X:x, Y:y}
				distance := cand.Distance(pos)
				coord := distpos{pos:cand, dist:distance}
				coords = append(coords, coord)
			}
		}
	}

	//sort
	sort.Slice(coords, func(i, j int) bool { return coords[i].dist < coords[j].dist } )

	return coords;
}

func (m *gamemap) FreeGridInRange(dist int, pos position) []position {
	coords := m.findGridInRange(dist, pos)

	free := m.freetiles

	var out []position

	for _, coord := range coords {
		for _, fre := range free {
			if fre.pos.X == coord.pos.X && fre.pos.Y == coord.pos.Y {
				out = append(out, coord.pos)
			}
		}
	}

	return out
}