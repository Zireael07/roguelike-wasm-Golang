// +build js

// build js annotation REQUIRED to use syscall/js

package main

import (
	//"fmt"
	"log"

	"syscall/js"
	"errors"
	"fmt" //for file loading
	"net/http"
	"io/ioutil"
)

//init the game
func main() {
	//create a new terminal
	term := &terminal{}
	err := term.Init()
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	defer term.Close()

	//RAF = the main driver
	go func() {
		for {
			term.ReqAnimFrame()
		}
	}()
	//for {
	newGame(term)
	//}
}

func newGame(term *terminal) {
	g := &game{}
	term.DrawBufferInit()

	//term.g = g //if we need access to g down the line
	g.Term = term
	g.GameInit()
	//recalc FOV
	g.clearFOV()
	//g.FOV.SetAllInvisible()
	// var opaque VB = func(x,y int32) bool {
	// 	//paranoia
	// 	if x >= 0 && y >= 0 && x <= int32(g.Map.width) && y <= int32(g.Map.height) {
	// 		return g.Map.tiles[x][y].IsWall() 
	// 	} else 
	// 	{ return true } 
	// }
	// var visit VE = func(x,y int32) {
	// 	//paranoia
	// 	if x >= 0 && y >= 0 && x <= int32(g.Map.width) && y <= int32(g.Map.height) {
	// 		g.Map.tiles[x][y].visible = true
	// 		g.Map.tiles[x][y].explored = true
	// 	}
	// }
	// var inmap IM = func(x,y int32) bool {
	// 	if x >= 0 && y >= 0 && x <= int32(g.Map.width) && y <= int32(g.Map.height){
	// 		return true
	// 	} else { return false } 
	// }
	//FOV for player position (player is always #0)
	pl_pos := g.entities[0].Components["position"].(PositionComponent).Pos
	//g.pp_FOV(int32(pl_pos.X), int32(pl_pos.Y), 5, opaque, visit, inmap)
	g.FOV.RayCast(pl_pos.X, pl_pos.Y, g.Map)
	//camera!!
	g.camera.update(pl_pos)
	//make the game run!
	g.gameeventLoop()

	//initial draw
	term.Clear()
	g.render()
	term.Flush()
	//block
	term.PressAnyKey()

}

var SaveError string

type terminal struct {
	//g 			*game
	DrawBuffer          []TermCell
	//to avoid drawing what hasn't changed
	drawBackBuffer      []TermCell
	display   js.Value
	cache     map[TermCell]js.Value
	ctx       js.Value
	width     int
	height    int
	mousepos  position
}



//JS frontend specific stuff
func (term *terminal) InitElements() error {
	canvas := js.Global().Get("document").Call("getElementById", "gamecanvas")
	canvas.Call("addEventListener", "contextmenu", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := args[0]
		e.Call("preventDefault")
		return nil
	}), false)
	canvas.Call("setAttribute", "tabindex", "1")
	term.ctx = canvas.Call("getContext", "2d")
	term.ctx.Set("imageSmoothingEnabled", false)
	term.width = 16
	term.height = 24
	canvas.Set("height", 24*TermHeight)
	canvas.Set("width", 16*TermWidth)
	term.cache = make(map[TermCell]js.Value)
	return nil
}

func (term *terminal) Draw(cell TermCell, x, y int) {
	var canvas js.Value
	if cv, ok := term.cache[cell]; ok {
		canvas = cv
	} else {
		canvas = js.Global().Get("document").Call("createElement", "canvas")
		canvas.Set("width", 16)
		canvas.Set("height", 24)
		ctx := canvas.Call("getContext", "2d")
		ctx.Set("imageSmoothingEnabled", false)
		buf := getImage(cell).Pix
		ua := js.Global().Get("Uint8Array").New(js.ValueOf(len(buf)))
		//actually sends over image pixels
		js.CopyBytesToJS(ua, buf)
		ca := js.Global().Get("Uint8ClampedArray").New(ua)
		imgdata := js.Global().Get("ImageData").New(ca, 16, 24)
		ctx.Call("putImageData", imgdata, 0, 0)
		term.cache[cell] = canvas
	}
	//old familiar JS canvas drawImage...
	term.ctx.Call("drawImage", canvas, x*term.width, term.height*y)
}

func (term *terminal) GetMousePos(evt js.Value) (int, int) {
	canvas := js.Global().Get("document").Call("getElementById", "gamecanvas")
	rect := canvas.Call("getBoundingClientRect")
	scaleX := canvas.Get("width").Float() / rect.Get("width").Float()
	scaleY := canvas.Get("height").Float() / rect.Get("height").Float()
	x := (evt.Get("clientX").Float() - rect.Get("left").Float()) * scaleX
	y := (evt.Get("clientY").Float() - rect.Get("top").Float()) * scaleY
	return (int(x) - 1) / term.width, (int(y) - 1) / term.height
}


func (term *terminal) Init() error {
	canvas := js.Global().Get("document").Call("getElementById", "gamecanvas")
	//gamediv := js.Global().Get("document").Call("getElementById", "gamediv")
	js.Global().Get("document").Call(
		//add JS key listener
		"addEventListener", "keydown", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			e := args[0]
			if !e.Get("ctrlKey").Bool() && !e.Get("metaKey").Bool() {
				e.Call("preventDefault")
			} else {
				return nil
			}
			s := e.Get("key").String()
			if s == "Unidentified" {
				s = e.Get("code").String()
			}
			if len(ch) < cap(ch) {
				ch <- TermInput{key: s}
			}
			return nil
		}))
	//mouse listeners
	canvas.Call(
		"addEventListener", "mousedown", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			e := args[0]
			x, y := term.GetMousePos(e)
			if len(ch) < cap(ch) {
				ch <- TermInput{mouse: true, mouseX: x, mouseY: y, button: e.Get("button").Int()}
			}
			return nil
		}))
	canvas.Call(
		"addEventListener", "mousemove", js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			e := args[0]
			x, y := term.GetMousePos(e)
			if x != term.mousepos.X || y != term.mousepos.Y {
				term.mousepos.X = x
				term.mousepos.Y = y
				//log.Printf("Mouse pos: %d, %d", x,y);
				if len(ch) < cap(ch) {
					ch <- TermInput{mouse: true, mouseX: x, mouseY: y, button: -1}
				}
			}
			return nil
		}))

	
	term.InitElements()
	return nil
}


//goroutines and channels here!
var ch chan TermInput
var interrupt chan bool

func init() {
	ch = make(chan TermInput, 5) //buffered
	interrupt = make(chan bool)
	Flushdone = make(chan bool)
	ReqFrame = make(chan bool)
}

func (term *terminal) Close() {
	// stub
}

func (term *terminal) Flush() {
	ReqFrame <- true
	<-Flushdone
}

// actually draw
var tmark float64


func (term *terminal) ReqAnimFrame() {
	<-ReqFrame
	js.Global().Get("window").Call("requestAnimationFrame",
		js.FuncOf(func(this js.Value, args []js.Value) interface{} { 
			term.FlushCallback(args[0]); return nil 
		}))
}



var Flushdone chan bool
var ReqFrame chan bool

func (term *terminal) DrawFrame() {
	if len(term.drawBackBuffer) != len(term.DrawBuffer) {
		term.drawBackBuffer = make([]TermCell, len(term.DrawBuffer))
	}

	for i := 0; i < len(term.DrawBuffer); i++ {
		//back buffer lets us skip what hasn't changed
		if term.DrawBuffer[i] == term.drawBackBuffer[i] {
			continue
		}
		// draw all cells now
		c := term.DrawBuffer[i]
		x, y := term.GetPos(i)
		term.Draw(c, x, y)

		//add to back buffer
		term.drawBackBuffer[i] = c
	}
	
}


//'t' was args[0] in the callback
func (term *terminal) FlushCallback(t js.Value) {
	//performance tracking
	// now := t.Float()
	// tdiff := now - tmark
	// //NOTE: if we remove fps from page, this NEEDS to be commented out!!!
	// js.Global().Get("document").Call("getElementById", "fps").Set("innerHTML", fmt.Sprintf("FPS: %.01f", 1000/tdiff))
	// tmark = now

	 //do the damn drawing
	 term.DrawFrame()
	
	//stub
	Flushdone <- true
}



// ----------------------------
//input
func (term *terminal) PollEvent() (in TermInput) {
	select {
	case in = <-ch:
	case in.interrupt = <-interrupt:
	}

	//log.Printf("Poll event: %v", in);
	
	return in
}


//blocks, waiting for input such as mouse move (or a key)
func (term *terminal) PressAnyKey() error {
	for {
		e := term.PollEvent()
		if e.interrupt {
			return errors.New("interrupted")
		}
		if e.key != "" || (e.mouse && e.button != -1) {
			return nil
		}
	}
}

//load server files
//https://stackoverflow.com/a/42718113
func getContent(url string) ([]byte, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("GET error: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("Status error: %v", resp.StatusCode)
    }

    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("Read body: %v", err)
    }

    return data, nil
}