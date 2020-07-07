package main

import (
	"math"
)

//Based on Gogue https://github.com/gogue-framework/gogue

// FieldOfVision represents an area that an entity can see, defined by the torch radius. The cos and sin tables are
// generated once on instantiation, so we don't have to build them each time we want to calculate visible distances.
type FieldOfVision struct {
	cosTable    map[int]float64
	sinTable    map[int]float64
	torchRadius int
}

// InitializeFOV generates the cos and sin tables, for 360 degrees, for use when raycasting to determine line of sight
func (f *FieldOfVision) InitializeFOV() {

	f.cosTable = make(map[int]float64)
	f.sinTable = make(map[int]float64)

	for i := 0; i < 360; i++ {
		ax := math.Sin(float64(i) / (float64(180) / math.Pi))
		ay := math.Cos(float64(i) / (float64(180) / math.Pi))

		f.sinTable[i] = ax
		f.cosTable[i] = ay
	}
}

// SetTorchRadius sets the radius of the FOVs torch, or how far the entity can see
func (f *FieldOfVision) SetTorchRadius(radius int) {
	if radius > 1 {
		f.torchRadius = radius
	}
}

//Equal to the code that used to live in main's game:clearFOV()
// SetAllInvisible makes all tiles on the gamemap invisible to the player.
func (f *FieldOfVision) SetAllInvisible(m *gamemap) {
	for x := 0; x < m.width; x++ {
		for y := 0; y < m.height; y++ {
			m.tiles[x][y].visible = false
		}
	}
}

// RayCast casts out rays each degree in a 360 circle from the player. If a ray passes over a floor (does not block sight)
// tile, keep going, up to the maximum torch radius (view radius) of the player. If the ray intersects a wall
// (blocks sight), stop, as the player will not be able to see past that. Every visible tile will get the Visible
// and Explored properties set to true.
func (f *FieldOfVision) RayCast(playerX, playerY int, m *gamemap) {

	for i := 0; i < 360; i++ {

		ax := f.sinTable[i]
		ay := f.cosTable[i]

		x := float64(playerX)
		y := float64(playerY)

		// Mark the players current position as explored
		tile := m.tiles[playerX][playerY]
		tile.explored = true
		tile.visible = true

		for j := 0; j < f.torchRadius; j++ {
			x -= ax
			y -= ay

			roundedX := int(round(x))
			roundedY := int(round(y))

			if x < 0 || x > float64(m.width-1) || y < 0 || y > float64(m.height-1) {
				// If the ray is cast outside of the map, stop
				break
			}

			tile := m.tiles[roundedX][roundedY]

			tile.explored = true
			tile.visible = true

			//if m.tiles[roundedX][roundedY].BlocksSight == true {
			if m.tiles[roundedX][roundedY].IsWall() {
				// The ray hit a wall, go no further
				break
			}
		}
	}
}

func round(f float64) float64 {
	return math.Floor(f + .5)
}