package main

import (
	//"log"
)

//just a stub for now
type game struct {
	Term *terminal;
	player position;
}

type position struct {
	X int
	Y int
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


func (g *game) render(){
	// g.Term.SetCell(2,2,'N',Color{255,0,0, 255}, Color{0,0,0,255}, true)
	// g.Term.SetCell(3,2,'e',Color{88,110,17, 255}, Color{0,0,0,255}, true)
	// g.Term.SetCell(4,2, 'o', Color{255,255,255, 255}, Color{0,0,255,255}, true)
	// g.Term.SetCell(5,2, 'n', Color{0,255, 0, 255}, Color{0,0,0,255}, true)
	g.Term.SetCell(g.player.X, g.player.Y, '@', Color{255, 255, 255, 255}, Color{0,0,0,255}, true)
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
			if (g.player.Distance(pos) < 2){
				g.player = pos
				g.Term.Clear()
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
