package main

import (
	"log"
	"fmt"
	"math/rand" //for RNG
	"time"
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
	//init RNG
	rand.Seed(time.Now().UnixNano())
	g.ECSInit()
	m := &gamemap{width: 20, height:20}
	m.InitMap()
	m.generateArenaMap()
	g.Map = m
}

type randomPick struct {
	chance int
	entry string
}

func (g *game) selectRandomItem() string {
	var chances []randomPick
	chances = []randomPick{ randomPick{25, "Pistol"}, randomPick{50, "Medkit"} }

	sum := 0
	for _, ch := range chances {
		sum = sum + ch.chance
	}
	//log.Print("Sum: %d", sum)
	roll := rand.Intn(sum) 
	log.Printf("Roll: %d", roll)

	//look up result
	res := ""
	for i, _ := range chances {
		low := 0
		high := chances[i].chance
		if i > 0 {
			low = chances[i-1].chance
			high = low+chances[i].chance
		}
		if roll > low && roll <= high {
			return chances[i].entry
		}
	}

	return res
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
	g.entities = append(g.entities, player)
	//NPCs!
	npcs_num := 2
	for i :=0; i < npcs_num; i++ {
		npc := &GameEntity{}
		npc.setupComponentsMap()
		//random position in range 1-19
		rnd_x := rand.Intn(19-1)+1
		rnd_y := rand.Intn(19-1)+1
		npc.AddComponent("position", PositionComponent{Pos:position{X:rnd_x, Y:rnd_y}})
		npc.AddComponent("renderable", RenderableComponent{Color{255, 0,0,255}, 'h'})
		npc.AddComponent("blocker", BlockerComponent{})
		npc.AddComponent("NPC", NPCComponent{})
		npc.AddComponent("stats", StatsComponent{hp:10, max_hp: 10, power:2})
		npc.AddComponent("name", NameComponent{"Thug"})
		g.entities = append(g.entities, npc)
		//log.Printf("Added npc...")
	}

	//item
	sel := g.selectRandomItem()
	log.Printf("Sel: %v", sel)
	if sel != "" {
		if sel == "Pistol" {
			//gun
			it := &GameEntity{}
			it.setupComponentsMap()
			it.AddComponent("position", PositionComponent{Pos:position{X:4, Y:4}})
			it.AddComponent("renderable", RenderableComponent{Color{0,255,255,255}, '('})
			it.AddComponent("item", ItemComponent{})
			it.AddComponent("name", NameComponent{"Pistol"})
			it.AddComponent("range", RangeComponent{6})
			g.entities = append(g.entities, it)
		} else if sel == "Medkit" {
			it := &GameEntity{}
			it.setupComponentsMap()
			it.AddComponent("position", PositionComponent{Pos:position{X:6, Y:6}})
			it.AddComponent("renderable", RenderableComponent{Color{255,0,0,255}, '!'})
			it.AddComponent("item", ItemComponent{})
			it.AddComponent("name", NameComponent{"Medkit"})
			it.AddComponent("medkit", MedkitComponent{4})
			g.entities = append(g.entities, it)
		}	
	}
	
	//g.entities = append(g.entities, gun)
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
				if !e.HasComponents([]string{"backpack"}){
					pos, _ := e.Components["position"].(PositionComponent)
					rend, _ := e.Components["renderable"].(RenderableComponent)
					if (g.Map.tiles[pos.Pos.X][pos.Pos.Y].visible) {
						g.Term.SetCell(pos.Pos.X, pos.Pos.Y, rend.Glyph, rend.Color, Color{0,0,0,255}, true)
					}
				}
			}
		}
	}

	//render GUI
	g.Term.SetCell(25,1, 'H', Color{255,255,255,255}, Color{0,0,0,255}, false)
	g.Term.SetCell(26,1, 'P', Color{255,255,255,255}, Color{0,0,0,255}, false)
	g.renderBar(27,1,10, float32(g.entities[0].Components["stats"].(StatsComponent).hp), float32(g.entities[0].Components["stats"].(StatsComponent).max_hp), 
	Color{255,115, 155, 255}, Color{128,0,0,255})

	g.Term.DrawColoredText(25, 5, "Inventory", Color{0,255,255,255})

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
				if !e.HasComponents([]string{"backpack"}) {
					pos_c, _ := e.Components["position"].(PositionComponent)
					if pos_c.Pos.X == pos.X && pos_c.Pos.Y == pos.Y {
						name, _ := e.Components["name"].(NameComponent)
						txt = name.name
						break
					}
				}
			}
		}
	}

	g.Term.DrawText(25, 3, txt)

}

func (g *game) renderActionMenu(pos position){
	if !pos.isValid(g){
		return
	}

	if !g.Map.tiles[pos.X][pos.Y].explored{
		return
	}

	txt := ""
	if g.GetItemsAtPos(pos) != nil {
		txt = "Get an item"
	}

	//this loop has to be named, otherwise engine gets stuck?
	menuloop:
	for {
		g.Term.DrawColoredText(pos.X+1, pos.Y, txt, Color{0,255,255,255})
		g.Term.Flush()
		in := g.Term.PollEvent()
		//log.Printf("Event: %v", in);
		if in.mouse {
			m_pos := position{X: in.mouseX, Y: in.mouseY}
			switch in.button {
				case -1:
					// do nothing
				//left
				case 0:
					if m_pos.X >= pos.X && m_pos.X <= 20 && m_pos.Y == pos.Y {
						log.Printf("Selected get")
						// player on the same tile
						pl_pos := g.entities[0].Components["position"].(PositionComponent)
						if pl_pos.Pos.X == pos.X && pl_pos.Pos.Y == pos.Y {
							//Actually get the item
							it := g.GetItemsAtPos(pos)
							if it != nil {
								it.AddComponent("backpack", InBackpackComponent{})
								g.logMessage("Picked up item")
								//AI gets to move
								for _, e := range g.entities {
									if e != nil {
										if e.HasComponents([]string{"NPC"}) {
											g.takeTurn(e)
										}
									}
								}
							}
						}
						
						
						break menuloop
					}
					//exit either way
					break menuloop
				//right
				case 2:
					//exit the menu
					break menuloop
			}
		}
	}
	log.Printf("Exited loop")
	
}

func (g *game) GetItemsAtPos(pos position) *GameEntity {
	// := aka walrus aka type inference doesn't work for nil
	var ret *GameEntity = nil
	for _, e := range g.entities {
		if e != nil {
			if e.HasComponents([]string{"position", "item"}) {
				if !e.HasComponents([]string{"backpack"}) {
					pos_c, _ := e.Components["position"].(PositionComponent)
					if pos_c.Pos.X == pos.X && pos_c.Pos.Y == pos.Y {
						ret = e
						break
					}
				}
			}
		}
	}
	return ret
}

func (g *game) GetItemsInventory() []*GameEntity {
	// := aka walrus aka type inference doesn't work here
	var ret [] *GameEntity
	for _, e := range g.entities {
		if e != nil {
			if e.HasComponents([]string{"backpack", "item", "name"}) {
				ret = append(ret, e)
			}
		}
	}
	return ret
}

func (g *game) renderInventory() {
	//log.Printf("render inventory...")
	items := g.GetItemsInventory()

	if len(items) < 1{
		log.Printf("No items in inventory...")
		return
	}
		

	x := 10
	y := 1
	for _, it := range items {
		name_c, _ := it.Components["name"].(NameComponent)
		g.Term.DrawColoredText(x,y, name_c.name, Color{0,255,255,255})
		y++
	}

	//this loop has to be named, otherwise engine gets stuck?
	menuloop:
	for {
		g.Term.Flush()
		in := g.Term.PollEvent()
		//log.Printf("Event: %v", in);
		if in.mouse {
			m_pos := position{X: in.mouseX, Y: in.mouseY}
			switch in.button {
				case -1:
					// do nothing
				//left
				case 0:
					if m_pos.X >= 10 && m_pos.X <= 20 {
						id := m_pos.Y-1
						// draw submenu
						g.Term.DrawColoredText(x+8,1, "Use", Color{0,255,255,255})
						g.Term.DrawColoredText(x+8,2, "Drop", Color{0,255,255,255})
						submenuloop:
							for {
								g.Term.Flush()
								in := g.Term.PollEvent()
								if in.mouse {
									m_pos := position{X: in.mouseX, Y: in.mouseY}
									switch in.button {
									case -1:
										//do nothing
									//left
									case 0:
										if m_pos.X >= 18 && m_pos.Y <= 25 {
											//use
											if m_pos.Y == 1 {
												log.Printf("Selected use")
												if items[id].HasComponent("medkit") {
													//use the medkit
													heal := items[id].Components["medkit"].(MedkitComponent).heal
													//heal the player
													statsComp := g.entities[0].Components["stats"].(StatsComponent)
													old_hp := g.entities[0].Components["stats"].(StatsComponent).hp
													g.logMessage(fmt.Sprintf("Medkit healed %d damage", heal)) //, Color{255,0,0,255})
													statsComp.hp = old_hp + heal
													//bit of a dance because we're not using pointers to Components
													g.entities[0].RemoveComponent("stats")
													g.entities[0].AddComponent("stats", statsComp)
													//nuke the medkit
													items[id].RemoveComponents([]string{"position", "renderable", "item", "medkit", "backpack"})
													//end turn
													//AI gets to move
													for _, e := range g.entities {
														if e != nil {
															if e.HasComponents([]string{"NPC"}) {
																g.takeTurn(e)
															}
														}
													}
													//break menuloop
												}
												if items[id].HasComponent("range") {
													shootloop:
														for {
															in := g.Term.PollEvent()
															//log.Printf("Event: %v", in);
															if in.mouse {
																m_pos := position{X: in.mouseX, Y: in.mouseY}
																g.Term.highlightPos(m_pos)
																g.describePosition(m_pos)
																g.Term.Flush()
																switch in.button {
																	case -1:
																		// do nothing
																	//left
																	case 0:
																		
																		//target
																		dist := items[id].Components["range"].(RangeComponent).dist
																		pl_pos := g.entities[0].Components["position"].(PositionComponent).Pos
					
																		//target in range
																		if m_pos.X <= pl_pos.X + dist && m_pos.Y <= pl_pos.Y + dist{
																			//are we targeting something?
																			blocker := g.getAllBlockers(m_pos)
																			if blocker != nil{
																				//damage it!
																				statsComp := blocker.Components["stats"].(StatsComponent)
																				damage := 6 //dummy
																				old_hp := blocker.Components["stats"].(StatsComponent).hp
																				g.logMessage(fmt.Sprintf("Player shoots enemy for %d damage", damage)) // Color{0,255,0,255})
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
					
																				//nuke the gun
																				//we're kind, we only nuke it if we actually hit something
																				items[id].RemoveComponents([]string{"position", "renderable", "item", "range", "backpack"})
																				
																			} else {
																				g.logMessage("Nothing to shoot at here")
																			}

																			//end turn
																			//AI gets to move
																			for _, e := range g.entities {
																				if e != nil {
																					if e.HasComponents([]string{"NPC"}) {
																						g.takeTurn(e)
																					}
																				}
																			}
																		}
																		break shootloop
																	//right
																	case 2:
																		//cancel
																		break shootloop
																}
															}
														}
														break submenuloop
												}

											//drop
											} else if m_pos.Y == 2 {
												log.Printf("Select drop")
												//paranoia
												if items[id].HasComponent("backpack") {
													items[id].RemoveComponent("backpack")
													pl_pos := g.entities[0].Components["position"].(PositionComponent).Pos
													posComp := PositionComponent{Pos:position{X:pl_pos.X, Y:pl_pos.Y}}
													//bit of a dance because we're not using pointers to Components
													items[id].RemoveComponent("position")
													items[id].AddComponent("position", posComp)
													g.logMessage("Player dropped item")

													//end turn
													//AI gets to move
													for _, e := range g.entities {
														if e != nil {
															if e.HasComponents([]string{"NPC"}) {
																g.takeTurn(e)
															}
														}
													}
												}
											}
											
										}
										break submenuloop
									//right
									case 2:
										break submenuloop
									}

								}
							}

					}
					// break out of menu loop (item actions are done!)
					break menuloop
				//right
				case 2:
					//exit the menu
					break menuloop
			}
		}
	}
	log.Printf("Exited ui loop")
	g.Term.Flush()
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
	if path[1].X != from.X || path[1].Y != from.Y {
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
				//did we click the right panel menu?
				if pos.X > 20 && pos.Y >= 5 {
					if pos.Y == 5 {
						g.Term.Clear()
						g.render()
						g.renderInventory();
						g.Term.Flush()
					}
				} else {
					//do nothing
					g.Term.Clear()
					g.render()
					g.Term.highlightPos(pos)
					g.describePosition(pos)
					g.Term.Flush()
				}
				
			}
		//right
		case 2:
			g.Term.Clear()
			g.render()
			g.Term.highlightPos(pos)
			//draw an action menu here
			g.renderActionMenu(pos)
			g.render()
			g.Term.Flush()
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
