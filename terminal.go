package main

//all of those imports are for images
import (
	"bytes"
	"encoding/base64"
	"image"
	//"image/color"
	"image/draw"
	"image/png"
	"log"
)

var (
	TermWidth                = 80
	TermHeight               = 25
)

var (
	ColorFg = Color{255,255,255, 255}
	ColorBg = Color{0,0,0, 255}
	ColorHighlightDark = Color{7, 54, 66, 255} //solarized base02
	ColorGray = Color{ 200,200,200, 255}
)

//generic interface
// type terminal struct {
// 	//g 			*game
// 	DrawBuffer          []TermCell
// 	//to avoid drawing what hasn't changed
// 	drawBackBuffer      []TermCell
// 	width     int
// 	height    int
// 	mousepos  position
// }

// Color represents a ARGB color in the console
type Color struct {
	R byte
	G byte
	B byte
	A byte
}

// New creates a new color from R,G,B values
func New(r, g, b byte) Color {
	return Color{r, g, b, 255}
}

// NewTransparent creates a new color from R,G,B,A values
func NewTransparent(r, g, b, a byte) Color {
	return Color{r, g, b, a}
}

// RGBA returns the color values as uint32s
func (c Color) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 8
	g = uint32(c.G)
	g |= g << 8
	b = uint32(c.B)
	b |= b << 8
	a = uint32(c.A)
	a |= a << 8
	return
}

type TermCell struct {
	Fg    Color
	Bg    Color
	R     rune
	InMap bool
}

type TermInput struct {
	key       string
	mouse     bool
	mouseX    int
	mouseY    int
	button    int
	interrupt bool
}

//Generic terminal stuff
//VSCode erroneously highlights all those for some reason but it builds and works fine
func (term *terminal) GetIndex(x, y int) int {
	return y*TermWidth + x
}

func (term *terminal) GetPos(i int) (int, int) {
	return i - (i/TermWidth)*TermWidth, i / TermWidth
}

func (term *terminal) Clear() {
	for i := 0; i < TermHeight*TermWidth; i++ {
		x, y := term.GetPos(i)
		term.SetCell(x, y, ' ', ColorFg, ColorBg, false)
	}
}

func (term *terminal) DrawText(x, y int, text string) {
	term.DrawColoredText(x, y, text, ColorFg)
}

func (term *terminal) DrawColoredText(x, y int, text string, fg Color) {
	term.DrawColoredTextOnBG(x, y, text, fg, ColorBg)
}

func (term *terminal) DrawColoredTextOnBG(x, y int, text string, fg, bg Color) {
	col := 0
	for _, r := range text {
		if r == '\n' {
			y++
			col = 0
			continue
		}
		if x+col >= 80 {
			break
		}
		term.SetCell(x+col, y, r, fg, bg, false)
		col++
	}
}

//'rune' is Go's UTF-8 friendly equivalent of char
func (term *terminal) SetCell(x, y int, r rune, fg, bg Color, inmap bool) {
	//prevent drawing out of bounds
	i := term.GetIndex(x, y)
	if i >= TermHeight*TermWidth {
		return
	}
	c := TermCell{R: r, Fg: fg, Bg: bg, InMap: inmap}
	term.DrawBuffer[i] = c
}

//original
func (term *terminal) GetCell(x,y int) TermCell {
	var c TermCell
	i := term.GetIndex(x, y)
	if i >= TermHeight*TermWidth {
		//dummy
		c = TermCell{R: ' ', Fg: ColorFg, Bg: ColorBg, InMap: false}
		return c
	}

	return term.DrawBuffer[i]
}

func (term *terminal) DrawBufferInit() {
	if len(term.DrawBuffer) == 0 {
		term.DrawBuffer = make([]TermCell, TermHeight*TermWidth)
	} else if len(term.DrawBuffer) != TermHeight*TermWidth {
		term.DrawBuffer = make([]TermCell, TermHeight*TermWidth)
	}
}

func (term *terminal) highlightPos(pos position) {
	c := term.GetCell(pos.X, pos.Y)
	term.SetCell(pos.X, pos.Y, c.R, c.Fg, ColorHighlightDark, c.InMap)
	//term.SetCell(pos.X, pos.Y, 'H', Color{255, 255, 255, 255}, Color{7, 54, 66, 255}, true)
}

// ------------------ tiles stuff

//TileImages saved as byte to work everywhere
var TileImgs map[string][]byte

var MapNames = map[rune]string{
	'¤':  "frontier",
	'√':  "hit",
	'Φ':  "magic",
	'☻':  "dreaming",
	'♫':  "footsteps",
	'#':  "hash",
	'@':  "player",
	'§':  "fog",
	'+':  "door",
	'.':  "dot",
	'♣':  "tree",
	'"':  "foliage",
	'•':  "tick",
	'●':  "rock",
	'×':  "times",
	',':  "comma",
	'}':  "rbrace",
	'%':  "percent",
	':':  "colon",
	'\\': "backslash",
	'~':  "tilde",
	'☼':  "sun",
	'*':  "asterisc",
	'—':  "hbar",
	'/':  "slash",
	'|':  "vbar",
	'∞':  "kill",
	' ':  "space",
	'[':  "lbracket",
	']':  "rbracket",
	')':  "rparen",
	'(':  "lparen",
	'>':  "stairs",
	'Δ':  "portal",
	'!':  "exclamation",
	';':  "semicolon",
	'_':  "stone",
}


//both tiles and ascii are essentially images
func getImage(cell TermCell) *image.RGBA {
	var pngImg []byte
	pngImg = TileImgs["map-notile"]
	if im, ok := TileImgs["letter-"+string(cell.R)]; ok {
		pngImg = im
		//handle longer names
	} else if im, ok := TileImgs["letter-"+MapNames[cell.R]]; ok {
		pngImg = im
	}
		//Go writes else on the same line
	// } else {
	// 	log.Printf("Could not find tile: %v", cell.R);
	// }

	buf := make([]byte, len(pngImg))
	base64.StdEncoding.Decode(buf, pngImg) // TODO: check error
	br := bytes.NewReader(buf)
	img, err := png.Decode(br)
	if err != nil {
		log.Printf("Could not decode png: %v", err)
	}
	rect := img.Bounds()
	rgbaimg := image.NewRGBA(rect)
	draw.Draw(rgbaimg, rect, img, rect.Min, draw.Src)
	bgc := cell.Bg
	fgc := cell.Fg
	for y := 0; y < rect.Max.Y; y++ {
		for x := 0; x < rect.Max.X; x++ {
			c := rgbaimg.At(x, y)
			r, _, _, _ := c.RGBA()
			if r == 0 {
				rgbaimg.Set(x, y, bgc)
			} else {
				rgbaimg.Set(x, y, fgc)
			}
		}
	}
	return rgbaimg
}

var Scripts map[string][]byte