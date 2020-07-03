package main

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
