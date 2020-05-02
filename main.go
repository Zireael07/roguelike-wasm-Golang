package main

import (
	"log"
	"fmt"
)

//just a stub for now
type game struct {
	Term *terminal;
	Map *gamemap
	entities []*GameEntity ///for ECS
	MessageLog []Message
}

type Message struct {
	text string
	Color Color
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
	player.AddComponent("stats", StatsComponent{hp:20, max_hp:20, power:5})
	player.AddComponent("name", NameComponent{"Player"})
	//NPC!
	npc := &GameEntity{}
	npc.setupComponentsMap()
	npc.AddComponent("position", PositionComponent{Pos:position{X:10, Y:10}})
	npc.AddComponent("renderable", RenderableComponent{Color{255, 0,0,255}, 'h'})
	npc.AddComponent("blocker", BlockerComponent{})
	npc.AddComponent("NPC", NPCComponent{})
	npc.AddComponent("stats", StatsComponent{hp:10, max_hp: 10, power:2})
	npc.AddComponent("name", NameComponent{"Thug"})

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


func (g *game) renderBar(x,y,width int, value, max_value float32, barColor, backColor Color) {
	// draw the bg
	for i := x; i <= x + width; i++ {
		g.Term.SetCell(i,y, ' ', backColor, backColor, false)
	} 
	//fill with color
	barWidth := int((value / max_value * float32(width)));
	if ( barWidth > 0 ) {
		for i := x; i <= x + barWidth; i++ {
			g.Term.SetCell(i,y, ' ', barColor, barColor, false)
		}
	}
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

	//render GUI
	g.Term.SetCell(25,1, 'H', Color{255,255,255,255}, Color{0,0,0,255}, false)
	g.Term.SetCell(26,1, 'P', Color{255,255,255,255}, Color{0,0,0,255}, false)
	g.renderBar(27,1,10, float32(g.entities[0].Components["stats"].(StatsComponent).hp), float32(g.entities[0].Components["stats"].(StatsComponent).max_hp), 
	Color{255,115, 155, 255}, Color{128,0,0,255})

	//draw message log
	y := 21
	for _, msg := range g.MessageLog {
		g.Term.DrawColoredText(0, y, msg.text, msg.Color)
		y++
	}
}

func (g *game) describePosition(pos position) {
	//log.Printf("Describing position... %v", pos)
	if !pos.isValid(g){
		return
	}

	g.Term.DrawText(25, 2, fmt.Sprintf("X:%d Y:%d", pos.X, pos.Y))

	if !g.Map.tiles[pos.X][pos.Y].explored{
		return
	}

	txt := ""

	for _, e := range g.entities {
		if e != nil {
			if e.HasComponents([]string{"position", "name"}) {
				pos_c, _ := e.Components["position"].(PositionComponent)
				if pos_c.Pos.X == pos.X && pos_c.Pos.Y == pos.Y {
					name, _ := e.Components["name"].(NameComponent)
					txt = name.name
					break
				}
			}
		}
	}

	g.Term.DrawText(25, 3, txt)

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
	//swap because the path is returned inverted?!
	to := e.Components["position"].(PositionComponent).Pos
	from := g.entities[0].Components["position"].(PositionComponent).Pos

	mp := &Path{game: g}

	path, _, found := AstarPath(mp, from, to)
	if !found {
		return
	}
	log.Printf("Path: %v", path);

	posComponent, _ := e.Components["position"].(PositionComponent)
	//#0, as usual, is our own position
	//log.Printf("Closest point: %v", path[1])
	if path[1].X != from.X && path[0].Y != from.Y {
		//path only takes into account walkable tiles, so no need for other checking
		posComponent.Pos = path[1]
		//bit of a dance because we're not using pointers to Components
		e.RemoveComponent("position")
		e.AddComponent("position", posComponent)
	} else {
		//log.Printf("Enemy kicks at your shins")
		statsComp := g.entities[0].Components["stats"].(StatsComponent)
		damage := e.Components["stats"].(StatsComponent).power
		old_hp := g.entities[0].Components["stats"].(StatsComponent).hp
		//log.Printf("Enemy dealt %d damage to player", damage)
		g.logMessage(fmt.Sprintf("Enemy dealt %d damage to player", damage)) //, Color{255,0,0,255})
		statsComp.hp = old_hp - damage
		//bit of a dance because we're not using pointers to Components
		g.entities[0].RemoveComponent("stats")
		g.entities[0].AddComponent("stats", statsComp)
		//dead
		if g.entities[0].Components["stats"].(StatsComponent).hp <= 0 {
			log.Printf("Player dead")
			//remove from ECS
			//for now, remove all components
			g.entities[0].RemoveComponents([]string{"position", "renderable", "blocker", "stats"})
		}
	}



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
	var blocker *GameEntity
	blocker = g.getAllBlockers(tg)
	if blocker != nil {
		//combat goes here
		statsComp := blocker.Components["stats"].(StatsComponent)
		damage := ent.Components["stats"].(StatsComponent).power
		old_hp := blocker.Components["stats"].(StatsComponent).hp
		g.logMessage(fmt.Sprintf("Player dealt %d damage to enemy", damage)) // Color{0,255,0,255})
		//bit of a dance because we're not using pointers to Components
		statsComp.hp = old_hp - damage
		blocker.RemoveComponent("stats")
		blocker.AddComponent("stats", statsComp)
		//dead
		if blocker.Components["stats"].(StatsComponent).hp <= 0 {
			g.logMessage("Enemy is killed.")
			//remove from ECS
			//for now, remove all components
			blocker.RemoveComponents([]string{"position", "renderable", "blocker", "stats", "NPC", "name"})
		}
		//log.Printf("The enemy growls at you!")
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

func (g *game) logMessage(msg string) {
	if len(g.MessageLog) >= 4 {
		// Throw away any messages that exceed our total queue size
		g.MessageLog = g.MessageLog[1:]
		//g.MessageLog = g.MessageLog[:len(g.MessageLog)-1]
	}
	message := Message{msg, Color{255,255,255,255}}
	g.MessageLog = append(g.MessageLog, message)

	// Prepend the message
	//g.MessageLog = append([]Message{message}, g.MessageLog...)
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
			g.describePosition(pos)
			g.Term.Flush()
		//left
		case 0:
			// move player
			pl_posComponent, _ := g.entities[0].Components["position"].(PositionComponent) 
			if (pl_posComponent.Pos.Distance(pos) < 2){
				dir := pos.sub(pl_posComponent.Pos)
				//log.Printf("direction: %v", dir)
				g.Term.Clear()
				g.MovePlayer(g.entities[0], dir)
				g.render()
				g.Term.Flush()
			} else {
				g.Term.Clear()
				g.render()
				g.Term.highlightPos(pos)
				g.describePosition(pos)
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
