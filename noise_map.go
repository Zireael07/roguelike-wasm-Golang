package main

import (
	"log"
	"math/rand" //for RNG
	"github.com/ojrac/opensimplex-go"
)

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
				m.tiles[x][y] = &maptile{glyph: '♣', fgColor: Color{0,255,0, 255}, blocks_move: true, visible: false}
			} else {
				m.tiles[x][y] = &maptile{glyph: '.', fgColor: Color{255,255,255, 255}, blocks_move: false, visible: false}
				free := &freetile{tile: m.tiles[x][y], pos: position{X:x, Y:y}}
				m.freetiles = append(m.freetiles, free) //add to list of free tiles
			}
		}
	}
}
