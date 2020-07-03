package main

import (
	"sort"
	"log"
)

//step one of finding biggest area of floor in matrix
func (m *gamemap) NumUnbrokenFloors_columns() [][]int {
	num_floors := make([][]int, m.width+1)
	for i := range num_floors {
		num_floors[i] = make([]int, m.height+1)
	}

	//actual values
	for x := 0; x <= m.width; x++ {
		for y := 0; y <= m.height; y++ {
			north := position{X:x, Y:y-1}
			//Golang has no ternary operator
			add := 0
			if y == 0 {
				add = 0
			} else {
				add = num_floors[north.X][north.Y]
			}

			val := 0
			if !m.tiles[x][y].IsWall(){
				val = 1 + add
			}
			num_floors[x][y] = val
		}
	}

	return num_floors
}

//parse it nicely
func (m *gamemap) GetUnbrokenFloors(num_floors [][]int) [][]int {
	var floors [][]int
	len_row := len(num_floors[0])-1
	//log.Printf("len num floors x %d", len_row)
	for y := 0; y <= len(num_floors)-1; y++ {
		var row []int
		for x := 0; x < len_row; x++ {
			row = append(row, num_floors[x][y])
		}

		//log.Printf("row length %d", len(row))
		floors = append(floors, row)
	}	

	return floors
}


type RectArea struct {
	area int
	rect Rect
}

//step two of finding biggest rectangle
func (m *gamemap) LargestAreaRect(floors [][]int) RectArea {
	var rects []RectArea

	//reverse order
	for y := len(floors)-1; y >= 0; y-- {
		//max rectangle
		rectdata := largestRectangleArea(floors[y], y)
		//log.Printf("Rect: %v, y: %d", rectdata, y)
		rects = append(rects, rectdata)
	}

	//sort
	sort.Slice(rects, func(i, j int) bool { return rects[i].area > rects[j].area } )

	return rects[0]
}

//based on my Rust code https://github.com/Zireael07/rust-web-roguelike-reboot/blob/master/src/map_builders/rectangle_builder.rs
func largestRectangleArea(heights []int, row int) RectArea {
	max := 0
	answer := Rect{pos1:position{X:0,Y:0}, pos2:position{X:0,Y:0}}
	stack := make([]int,0)
	
	// populate stack with 0 = (index of first element in heights) to avoid checking for empty stack
	stack = append(stack, 0)
	// insert -1 from both sides so that we don't have to test for corner cases
	// trick described in e.g. http://shaofanlai.com/post/85
	heights = append(heights, -1)
	heights = prependInt(heights, -1)

    for i,h := range heights {
		//if higher than previous, push to stack
		if h > heights[stack[len(stack)-1]] {
		//if (len(stack) == 0 || heights[i] > heights[stack[len(stack) -1]]){
			stack = append(stack, i); //i is our x position
            //i++
        } else{
			//if this bar is lower, pop from stack
			for h < heights[stack[len(stack) -1]]{
				pop := stack[len(stack)-1] 
				stack = stack[:len(stack)-1]
				h := heights[pop]
	
				width := i - stack[len(stack)-1] -1
	
				area := h * width
				if area > max {
					max = area
					// this algo is bottom-up, so deduce height from y to get the top
					//it goes left->right, so the i (position in histogram) is the right end
					answer = Rect{pos1:position{X:i-width,Y:row-h+1}, pos2:position{X:i-1,Y:row}}
				}
			}
			stack = append(stack, i)
        }
    }
	
    return RectArea{area:max, rect:answer}
}

func (m *gamemap) RectangleDetect() {
	floors := m.NumUnbrokenFloors_columns()
	row_floors := m.GetUnbrokenFloors(floors)
	//log.Printf("%v", row_floors)

	largest := m.LargestAreaRect(row_floors)
	log.Printf("%v", largest.rect)
	m.debugLargestArea(largest.rect)

	//add to submaps
	m.submaps = append(m.submaps, largest.rect)
	log.Printf("Submaps: %v", m.submaps)
}

func (m *gamemap) debugLargestArea(rect Rect) {
	for x := rect.pos1.X; x <= rect.pos2.X; x++{
		for y:= rect.pos1.Y; y <= rect.pos2.Y; y++{
			m.tiles[x][y] = &maptile{glyph: ',', fgColor: Color{0,255,0,255}, blocks_move: false, visible: false}
		}
	}
}
