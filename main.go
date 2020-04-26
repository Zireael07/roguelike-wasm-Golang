package main

import (
	"log"
)

//just a stub for now
type game struct {
	Term *terminal;
}

type position struct {
	X int
	Y int
}

func (g *game) render(){
	g.Term.SetCell(2,2,'N',Color{255,0,0, 255}, Color{0,0,0,255}, true)
	g.Term.SetCell(3,2,'e',Color{88,110,17, 255}, Color{0,0,0,255}, true)
	g.Term.SetCell(4,2, 'o', Color{255,255,255, 255}, Color{0,0,255,255}, true)
	g.Term.SetCell(5,2, 'n', Color{0,255, 0, 255}, Color{0,0,0,255}, true)
}

func (g *game) HandlePlayerEvent() () {
	in := g.Term.PollEvent()
	//log.Printf("Event: %v", in);
	if in.mouse {
		pos := position{X: in.mouseX, Y: in.mouseY}

		g.Term.Clear()
		g.render()
		g.Term.highlightPos(pos)
		g.Term.Flush()
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
