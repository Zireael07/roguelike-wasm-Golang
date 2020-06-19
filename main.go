package main

import (
	"log"
	"fmt"
	"math/rand" //for RNG
	"time"
	"github.com/yuin/gopher-lua" //Lua
	"layeh.com/gopher-luar"
	//"github.com/yuin/gluamapper"
)

//just a stub for now
type game struct {
	Term *terminal
	Map *gamemap
	camera *Camera
	entities []*GameEntity ///for ECS
	MessageLog []Message
}

type Message struct {
	text string
	Color Color
}

func (g *game) GameInit() {
	//init RNG
	rand.Seed(time.Now().UnixNano())

	m := &gamemap{width: 20, height:20}
	m.InitMap()
	//m.generateArenaMap()
	m.generatePerlinMap()
	g.Map = m

	//camera
	c := &Camera{20,20,1,1,0,0}
	g.camera = c

	g.LuaInit()
	//Lua generates the map, among other things
	g.ECSInit()
}

func (g *game) Add(who *GameEntity) {
    g.entities = append(g.entities, who)
}

func (g *game) LuaInit(){
	L := lua.NewState()
	defer L.Close()
	// doFile doesn't work on WASM
	//if err := L.DoFile("hello.lua"); err != nil {

	//script := `print("hello WASM from lua")`

	//test Lua->Go interop
	//ent := &GameEntity{}
	//ent.setupComponentsMap() //crucial!
	L.SetGlobal("entities", luar.New(L, g))

	L.SetGlobal("Ent", luar.NewType(L, GameEntity{}))
	//L.SetGlobal("ent", luar.New(L, ent))
	//components
	L.SetGlobal("Position", luar.NewType(L, PositionComponent{}))
	L.SetGlobal("Renderable", luar.NewType(L, RenderableComponent{}))
	L.SetGlobal("Stats", luar.NewType(L, StatsComponent{}))
	L.SetGlobal("Name", luar.NewType(L, NameComponent{}))
	L.SetGlobal("Blocker", luar.NewType(L, BlockerComponent{}))
	L.SetGlobal("NPC", luar.NewType(L, NPCComponent{}))
	L.SetGlobal("Item", luar.NewType(L, ItemComponent{}))
	L.SetGlobal("Medkit", luar.NewType(L, MedkitComponent{}))

	//map
	L.SetGlobal("map", luar.New(L, g.Map))
	L.SetGlobal("Color", luar.NewType(L, Color{}))

	//this contains a byteslice
	script := Scripts["hello"]

	if err := L.DoString(string(script)); err != nil {
		panic(err)
	}

	//debug
	log.Printf("Entities: %d", len(g.entities))

	// var data NameComponent
	// if err := gluamapper.Map(L.GetGlobal("data").(*lua.LTable), &data); err != nil {
	// 	panic(err)
	//    }
	// //debug
	// log.Printf("Name: %s", data.Name)  	 
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
	player.SetupComponentsMap() //crucial!
	player.AddComponent("player", PlayerComponent{})

	//closest free position to 1,1
	pos_s := g.Map.FreeGridInRange(20, position{X:1, Y: 1})
	if len(pos_s) < 1 {
		//TODO: assert/bail out if no grid in range
		log.Printf("Could not place player")
	} else {
		log.Printf("Placed player at %v", pos_s[0])
	}
	player.AddComponent("position", PositionComponent{Pos:pos_s[0]})

	//player.AddComponent("position", PositionComponent{Pos:position{X: 1, Y: 1}})
	player.AddComponent("renderable", RenderableComponent{Color{255,255,255,255}, '@'})
	player.AddComponent("stats", StatsComponent{Hp:20, Max_hp:20, Power:5})
	player.AddComponent("name", NameComponent{"Player"})
	
	g.entities = moveToFront(player, g.entities)
	//g.entities = append(g.entities, player)
	//NPCs!
	npcs_num := 2
	for i :=0; i < npcs_num; i++ {
		npc := &GameEntity{}
		npc.SetupComponentsMap()
		//random position in range 1-19
		//rnd_x := g.randRange(1,19)
		//rnd_y := g.randRange(1,19)
		
		//random position in free list
		rnd_pos_id := randRange(1, len(g.Map.freetiles)-1)
		rnd_pos := g.Map.freetiles[rnd_pos_id]
		rnd_x := rnd_pos.pos.X
		rnd_y := rnd_pos.pos.Y
		
		npc.AddComponent("position", PositionComponent{Pos:position{X:rnd_x, Y:rnd_y}})
		npc.AddComponent("renderable", RenderableComponent{Color{255, 0,0,255}, 'h'})
		npc.AddComponent("blocker", BlockerComponent{})
		npc.AddComponent("NPC", NPCComponent{})
		npc.AddComponent("stats", StatsComponent{Hp:10, Max_hp: 10, Power:2})
		npc.AddComponent("name", NameComponent{"Thug"})
		g.entities = append(g.entities, npc)
		log.Printf("Added npc @ %d, %d ...", rnd_x, rnd_y)
	}

	//item
	sel := g.selectRandomItem()
	log.Printf("Sel: %v", sel)
	if sel != "" {
		if sel == "Pistol" {
			//gun
			it := &GameEntity{}
			it.SetupComponentsMap()
			//closest free position
			pos_s := g.Map.FreeGridInRange(20, position{X:4, Y: 4})
			if len(pos_s) < 1 {
				return
			}
			it.AddComponent("position", PositionComponent{Pos:pos_s[0]})
			it.AddComponent("renderable", RenderableComponent{Color{0,255,255,255}, '('})
			it.AddComponent("item", ItemComponent{})
			it.AddComponent("name", NameComponent{"Pistol"})
			it.AddComponent("range", RangeComponent{6})
			g.entities = append(g.entities, it)
		} else if sel == "Medkit" {
			it := &GameEntity{}
			it.SetupComponentsMap()
			//closest free position
			pos_s := g.Map.FreeGridInRange(20, position{X:6, Y: 6})
			if len(pos_s) < 1 {
				return
			}
			it.AddComponent("position", PositionComponent{Pos:pos_s[0]})
			it.AddComponent("renderable", RenderableComponent{Color{255,0,0,255}, '!'})
			it.AddComponent("item", ItemComponent{})
			it.AddComponent("name", NameComponent{"Medkit"})
			it.AddComponent("medkit", MedkitComponent{4})
			g.entities = append(g.entities, it)
		}	
	}
	
	//g.entities = append(g.entities, gun)
	//debug
	log.Printf("Entities: %d", len(g.entities))
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
	//camera
	width_st := g.camera.getWidthStart()
	width_end := g.camera.getWidthEnd(g.Map)
	height_st := g.camera.getHeightStart()
	height_end := g.camera.getHeightEnd(g.Map)


	for x := 0; x <= g.Map.width; x++ {
		for y := 0; y <= g.Map.height; y++ {
			if x >= width_st && x <= width_end && y >= height_st && y <= height_end {
				if (g.Map.tiles[x][y].visible){
					g.Term.SetCell(x, y, g.Map.tiles[x][y].glyph, g.Map.tiles[x][y].fgColor, Color{0,0,0,255}, true)
				} else if (g.Map.tiles[x][y].explored) {
					g.Term.SetCell(x, y, g.Map.tiles[x][y].glyph, Color{120,120,120,255},Color{0,0,0,255}, true)
				}
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

	//camera
	width_st := g.camera.getWidthStart()
	width_end := g.camera.getWidthEnd(g.Map)
	height_st := g.camera.getHeightStart()
	height_end := g.camera.getHeightEnd(g.Map)


	// Render all renderable entities to the screen
	for _, e := range g.entities {
		if e != nil {
			if e.HasComponents([]string{"position", "renderable"}) {
				pos, _ := e.Components["position"].(PositionComponent)
				//if not in camera view
				if pos.Pos.X < width_st || pos.Pos.X > width_end || pos.Pos.Y < height_st || pos.Pos.Y > height_end {
					continue
				}

				if e.HasComponents([]string{"backpack"}){
					continue
				}


				rend, _ := e.Components["renderable"].(RenderableComponent)
				if (g.Map.tiles[pos.Pos.X][pos.Pos.Y].visible) {
					g.Term.SetCell(pos.Pos.X, pos.Pos.Y, rend.Glyph, rend.Color, Color{0,0,0,255}, true)
				}
			}
		}
	}

	//render GUI
	g.Term.SetCell(25,1, 'H', Color{255,255,255,255}, Color{0,0,0,255}, false)
	g.Term.SetCell(26,1, 'P', Color{255,255,255,255}, Color{0,0,0,255}, false)
	g.renderBar(27,1,10, float32(g.entities[0].Components["stats"].(StatsComponent).Hp), float32(g.entities[0].Components["stats"].(StatsComponent).Max_hp), 
	Color{255,115, 155, 255}, Color{128,0,0,255})

	g.Term.DrawColoredText(25, 6, "Inventory", Color{0,255,255,255})

	//draw message log
	y := 21
	for _, msg := range g.MessageLog {
		g.Term.DrawColoredText(0, y, msg.text, msg.Color)
		y++
	}
}

//NOTE: Not very tenable because it requires every color to be listed in terminal.go:20
func (g *game) getTerrainName(terrain *maptile) string {
	ret := ""
	if terrain.glyph == '#' && terrain.fgColor == ColorFg {
		ret = "wall"
	}
	if terrain.glyph == '.' && terrain.fgColor == ColorGray {
		ret = "floor"
	}
	return ret
}

func (g *game) describePosition(pos position) {
	//log.Printf("Describing position... %v", pos)
	if !pos.isValid(g){
		return
	}

	g.Term.DrawText(25, 2, fmt.Sprintf("X:%d Y:%d", pos.X, pos.Y))
	//display directions
	pl_posComponent, _ := g.entities[0].Components["position"].(PositionComponent) 
	if pos.X != pl_posComponent.Pos.X || pos.Y != pl_posComponent.Pos.Y {
		dir := pos.Dir(pl_posComponent.Pos)
		//log.Printf("Direction: %s ", dir.String())
		g.Term.DrawText(34, 2, fmt.Sprintf("(%s)", dir.String()))
	}

	if !g.Map.tiles[pos.X][pos.Y].explored{
		return
	}

	terrain := ""
	terrain = g.getTerrainName(g.Map.tiles[pos.X][pos.Y])
	g.Term.DrawText(25,3, terrain)

	sym_txt := ' '
	sym_col := ColorFg
	txt := ""

	for _, e := range g.entities {
		if e != nil {
			if e.HasComponents([]string{"position", "name", "renderable"}) {
				if !e.HasComponents([]string{"backpack"}) {
					pos_c, _ := e.Components["position"].(PositionComponent)
					if pos_c.Pos.X == pos.X && pos_c.Pos.Y == pos.Y {
						name, _ := e.Components["name"].(NameComponent)
						rend, _ := e.Components["renderable"].(RenderableComponent)
						txt = name.Name
						sym_txt = rend.Glyph
						sym_col = rend.Color
						break
					}
				}
			}
		}
	}

	g.Term.SetCell(int(25), int(4), sym_txt, sym_col, ColorBg, false)
	g.Term.DrawText(25, 5, txt)

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

	//paranoia
	if len(path) < 2 {
		log.Printf("Path too short!")
		return
	}

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
		damage := e.Components["stats"].(StatsComponent).Power
		old_hp := g.entities[0].Components["stats"].(StatsComponent).Hp
		//log.Printf("Enemy dealt %d damage to player", damage)
		g.logMessage(fmt.Sprintf("Enemy dealt %d damage to player", damage)) //, Color{255,0,0,255})
		statsComp.Hp = old_hp - damage
		//bit of a dance because we're not using pointers to Components
		g.entities[0].RemoveComponent("stats")
		g.entities[0].AddComponent("stats", statsComp)
		//dead
		if g.entities[0].Components["stats"].(StatsComponent).Hp <= 0 {
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

func (g *game) MovePlayer(ent *GameEntity, dir position){
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
		damage := ent.Components["stats"].(StatsComponent).Power
		old_hp := blocker.Components["stats"].(StatsComponent).Hp
		g.logMessage(fmt.Sprintf("Player dealt %d damage to enemy", damage)) // Color{0,255,0,255})
		//bit of a dance because we're not using pointers to Components
		statsComp.Hp = old_hp - damage
		blocker.RemoveComponent("stats")
		blocker.AddComponent("stats", statsComp)
		//dead
		if blocker.Components["stats"].(StatsComponent).Hp <= 0 {
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
	//camera!!!
	g.camera.update(posComp.Pos)
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
			//TODO: get player entity via a Player component/tag
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
				if pos.X > 20 && pos.Y >= 6 {
					if pos.Y == 6 {
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
