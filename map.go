package main


type maptile struct {
	glyph rune
	blocks_move bool
	//FOV
	explored bool
	visible bool
}


//because 'map' in Go is a data structure...
type gamemap struct {
	width int
	height int
	tiles [][]*maptile //2d array/slice of tiles, nothing unusual
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
				m.tiles[x][y] = &maptile{glyph: '#', blocks_move: true, visible: false}
			} else {
				m.tiles[x][y] = &maptile{glyph: '.', blocks_move: false, visible: false}
			}
		}
	}
}
