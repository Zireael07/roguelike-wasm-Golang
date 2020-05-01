package main

import (
	"log"
)

//just a stub for now
type game struct {
	Term *terminal;
	Map *gamemap
	entities []*GameEntity ///for ECS
}

type position struct {
	X int
	Y int
}

func (g *game) GameInit() {
	g.ECSInit()
	m := &gamemap{width: 20, height:20}
	m.InitMap()
	m.generateArenaMap()
	g.Map = m
}

func (g *game) ECSInit() {
	// Create a player Entity, and add them to our slice of Entities
	player := &GameEntity{}
	player.setupComponentsMap() //crucial!
	player.AddComponent("player", PlayerComponent{})
	player.AddComponent("position", PositionComponent{Pos:position{X: 1, Y: 1}})
	player.AddComponent("renderable", RenderableComponent{Color{255,255,255,255}, '@'})
	//NPC!
	npc := &GameEntity{}
	npc.setupComponentsMap()
	npc.AddComponent("position", PositionComponent{Pos:position{X:10, Y:10}})
	npc.AddComponent("renderable", RenderableComponent{Color{255, 0,0,255}, 'h'})
	npc.AddComponent("blocker", BlockerComponent{})
	npc.AddComponent("NPC", NPCComponent{})

	g.entities = append(g.entities, player)
	g.entities = append(g.entities, npc)
}


//seriously, not in std?
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

func (pos position) isValid(g *game) bool {
	if pos.X >= 0 && pos.Y >= 0 && pos.X <= g.Map.width && pos.Y <= g.Map.height {
		return true
	} else { return false } 
}

func (g *game) clearFOV() {
	for x := 0; x < g.Map.width; x++ {
		for y := 0; y < g.Map.height; y++ {
			//log.Printf("Clear %d %d", x, y)
			g.Map.tiles[x][y].visible = false
		}
	}
}

func (g *game) renderMap(){
	for x := 0; x <= g.Map.width; x++ {
		for y := 0; y <= g.Map.height; y++ {
			if (g.Map.tiles[x][y].visible){
				g.Term.SetCell(x, y, g.Map.tiles[x][y].glyph, Color{255,255,255,255},Color{0,0,0,255}, true)
			} else if (g.Map.tiles[x][y].explored) {
				g.Term.SetCell(x, y, g.Map.tiles[x][y].glyph, Color{120,120,120,255},Color{0,0,0,255}, true)
			}
			
		}
	}
}


func (g *game) render(){
	// g.Term.SetCell(2,2,'N',Color{255,0,0, 255}, Color{0,0,0,255}, true)
	// g.Term.SetCell(3,2,'e',Color{88,110,17, 255}, Color{0,0,0,255}, true)
	// g.Term.SetCell(4,2, 'o', Color{255,255,255, 255}, Color{0,0,255,255}, true)
	// g.Term.SetCell(5,2, 'n', Color{0,255, 0, 255}, Color{0,0,0,255}, true)

	g.renderMap()


	// Render all renderable entities to the screen
	for _, e := range g.entities {
		if e != nil {
			if e.HasComponents([]string{"position", "renderable"}) {
				pos, _ := e.Components["position"].(PositionComponent)
				rend, _ := e.Components["renderable"].(RenderableComponent)

				g.Term.SetCell(pos.Pos.X, pos.Pos.Y, rend.Glyph, rend.Color, Color{0,0,0,255}, true)
			}
		}
	}
}

type Path struct {
	game      *game
	neighbors [8]position
}

func (p *Path) Cost(from, to position) int {
	return 1
}

func (p *Path) Estimation(from, to position) int {
	return from.Distance(to)
}

func (p *Path) Neighbors(pos position) []position {
	candidates := [8]position{position{-1,-1}, position{1,-1}, position{-1,1}, position{1,1}, position{0,-1}, position{-1,0}, position{0,1}, position{1,0}}
	var res []position
	for x := 0; x < 8; x++ {
		neighbour := position{pos.X + candidates[x].X, pos.Y + candidates[x].Y};
		if (neighbour.isValid(p.game) && !p.game.Map.tiles[neighbour.X][neighbour.Y].IsWall()){
			//add to array
			res = append(res, neighbour)
		}
	}
	return res;
}


func (g *game) takeTurn(e *GameEntity){
	from := e.Components["position"].(PositionComponent).Pos
	to := g.entities[0].Components["position"].(PositionComponent).Pos

	mp := &Path{game: g}

	path, _, found := AstarPath(mp, from, to)
	if !found {
		return
	}
	log.Printf("Path: %v", path);

	//return path
}

func (g *game) getAllBlockers(tg position) *GameEntity {
	// := aka walrus aka type inference doesn't work for nil
	var ret *GameEntity = nil
	for _, e := range g.entities {
		if e != nil {
			if e.HasComponents([]string{"position", "blocker"}) {
				pos, _ := e.Components["position"].(PositionComponent)
				if pos.Pos.X == tg.X && pos.Pos.Y == tg.Y {
					ret = e
					break
				}
			}
		}
	}
	return ret
}

func (g *game) MovePlayer (ent *GameEntity, dir position){
	//log.Printf("Move %v", dir)
	posComponent, _ := ent.Components["position"].(PositionComponent)
	tg := position{posComponent.Pos.X+dir.X, posComponent.Pos.Y+dir.Y}
	//check for blocked tiles
	if g.Map.tiles[tg.X][tg.Y].IsWall(){
		return
	}

	//check for blocking entities
	if g.getAllBlockers(tg) != nil {
		log.Printf("The enemy growls at you!")
		return
	}

	posComponent.Pos = tg
	//bit of a dance because we're not using pointers to Components
	ent.RemoveComponent("position")
	ent.AddComponent("position", posComponent)
	g.onPlayerMove(posComponent)
}

func (g *game) onPlayerMove(posComp PositionComponent){
	//AI gets to move
	for _, e := range g.entities {
		if e != nil {
			if e.HasComponents([]string{"NPC"}) {
				g.takeTurn(e)
			}
		}
	}


	//g.player = pos
	g.Term.Clear()
	//recalc FOV
	g.clearFOV()
	var opaque VB = func(x,y int32) bool {
		//paranoia
		if x >= 0 && y >= 0 && x <= int32(g.Map.width) && y <= int32(g.Map.height) {
			return g.Map.tiles[x][y].IsWall() 
		} else 
		{ return true } 
	}
	var visit VE = func(x,y int32) {
		//paranoia
		if x >= 0 && y >= 0 && x <= int32(g.Map.width) && y <= int32(g.Map.height) {
			g.Map.tiles[x][y].visible = true
			g.Map.tiles[x][y].explored = true
		}
	}
	var inmap IM = func(x,y int32) bool {
		if x >= 0 && y >= 0 && x <= int32(g.Map.width) && y <= int32(g.Map.height){
			return true
		} else { return false } 
	}
	g.pp_FOV(int32(posComp.Pos.X), int32(posComp.Pos.Y), 5, opaque, visit, inmap)	

}

func (g *game) HandlePlayerEvent() () {
	in := g.Term.PollEvent()
	//log.Printf("Event: %v", in);
	if in.mouse {
		pos := position{X: in.mouseX, Y: in.mouseY}
		switch in.button {
		//no button
		case -1:
			g.Term.Clear()
			g.render()
			g.Term.highlightPos(pos)
			g.Term.Flush()
		//left
		case 0:
			// move player
			pl_posComponent, _ := g.entities[0].Components["position"].(PositionComponent) 
			if (pl_posComponent.Pos.Distance(pos) < 2){
				dir := pos.sub(pl_posComponent.Pos)
				//log.Printf("direction: %v", dir)
				g.MovePlayer(g.entities[0], dir)
				g.render()
				g.Term.Flush()
			} else {
				g.Term.Clear()
				g.render()
				g.Term.highlightPos(pos)
				g.Term.Flush()
			}
		}

		
	}
}

//MAIN LOOP
//NOTE: method has () first
func (g *game) gameeventLoop() {
//loop:
	for {
		//TODO: if dead break loop
		g.HandlePlayerEvent()
	}
}
