package main

import (
	"sort"
	"log"
	"math/rand" //for RNG
	"github.com/ojrac/opensimplex-go"
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


//because 'map' in Go is a data structure...
type gamemap struct {
	width int
	height int
	tiles [][]*maptile //2d array/slice of tiles, nothing unusual
	freetiles []*freetile //list of all free tiles
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

func (m *gamemap) generateArenaMap() {
	// Generates a large, empty room, with walls ringing the outside edges
	for x := 0; x <= m.width; x++ {
		for y := 0; y <= m.height; y++ {
			if x == 0 || x == m.width || y == 0 || y == m.height {
				m.tiles[x][y] = &maptile{glyph: '#', fgColor: Color{255,255,255, 255}, blocks_move: true, visible: false}
			} else {
				m.tiles[x][y] = &maptile{glyph: '.', fgColor: Color{255,255,255, 255}, blocks_move: false, visible: false}
				//m.freetiles = append(m.freetiles, m.tiles[x][y]) //add to list of free tiles
			}
		}
	}

	//two random pillars
	for i :=0; i < 2; i++ {
		//random position in range 4-18
		rnd_x := randRange(4,18)
		rnd_y := randRange(4,18)
		m.tiles[rnd_x][rnd_y] = &maptile{glyph: '#', fgColor: Color{255,255,255,255}, blocks_move: true, visible: false}
	}

	//mark free tiles as such
	for x := 0; x <= m.width; x++ {
		for y := 0; y <= m.height; y++ {
			if !m.tiles[x][y].IsWall() {
				free := &freetile{tile: m.tiles[x][y], pos: position{X:x, Y:y}}
				m.freetiles = append(m.freetiles, free) //add to list of free tiles
			}
		}
	}
}

//this one is Lua-exposed
func (m *gamemap) GenerateArenaMapData(wall_glyph, floor_glyph rune, wall_color, floor_color Color) {
	// Generates a large, empty room, with walls ringing the outside edges
	for x := 0; x <= m.width; x++ {
		for y := 0; y <= m.height; y++ {
			if x == 0 || x == m.width || y == 0 || y == m.height {
				m.tiles[x][y] = &maptile{glyph: wall_glyph, fgColor: wall_color, blocks_move: true, visible: false}
			} else {
				m.tiles[x][y] = &maptile{glyph: floor_glyph, fgColor: floor_color, blocks_move: false, visible: false}
				//free := &freetile{tile: m.tiles[x][y], pos: position{X:x, Y:y}}
				//m.freetiles = append(m.freetiles, free) //add to list of free tiles
			}
		}
	}

	//two random pillars
	for i :=0; i < 2; i++ {
		//random position in range 4-18
		rnd_x := randRange(4,18)
		rnd_y := randRange(4,18)
		m.tiles[rnd_x][rnd_y] = &maptile{glyph: wall_glyph, fgColor: wall_color, blocks_move: true, visible: false}
	}

	// test
	m.tiles[1][1] = &maptile{glyph: wall_glyph, fgColor: wall_color, blocks_move: true, visible: false}

	//mark free tiles as such
	for x := 0; x <= m.width; x++ {
		for y := 0; y <= m.height; y++ {
			if !m.tiles[x][y].IsWall() {
				free := &freetile{tile: m.tiles[x][y], pos: position{X:x, Y:y}}
				m.freetiles = append(m.freetiles, free) //add to list of free tiles
			}
		}
	}
}

func (m *gamemap) generatePerlinMap() {
	rnd := rand.Int63()
	log.Printf("Random: %d", rnd)
	noise := opensimplex.New(rnd)
	//heightmap = a 2D array of noise data
	heightmap := make([][]float64, m.width+1)
	for i := range heightmap {
		heightmap[i] = make([]float64, m.height+1)
	}

	for x := 0; x <= m.width; x++ {
		for y := 0; y <= m.height; y++ {
			xFloat := float64(x) / float64(m.width)
			yFloat := float64(y) / float64(m.height)
			heightmap[x][y] = noise.Eval2(xFloat, yFloat) * 255 // because default values are very small
		}
	}

	//actual map
	for x := 0; x <= m.width; x++ {
		for y := 0; y <= m.height; y++ {
			if heightmap[x][y] > 0 {
				m.tiles[x][y] = &maptile{glyph: '#', fgColor: Color{255,255,255, 255}, blocks_move: true, visible: false}
			} else {
				m.tiles[x][y] = &maptile{glyph: '.', fgColor: Color{255,255,255, 255}, blocks_move: false, visible: false}
				free := &freetile{tile: m.tiles[x][y], pos: position{X:x, Y:y}}
				m.freetiles = append(m.freetiles, free) //add to list of free tiles
			}
		}
	}
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