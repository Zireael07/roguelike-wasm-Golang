package main

type Camera struct {
	width int
	height int
	x int
	y int
	top_x int
	top_y int
	//offset position
}

func (c *Camera) update(pos position) {
	c.x = pos.X
	c.y = pos.Y
	c.top_x = c.x - c.width/2
	c.top_y = c.y - c.height/2
}

func (c *Camera) getWidthStart() int {
	if c.top_x > 0 {
		return c.top_x
	} else {
		return 0
	}
}

func (c *Camera) getWidthEnd(m *gamemap) int {
	if c.top_x + c.width <= m.width {
		return c.top_x + c.width
	} else {
		return m.width
	}
}

func (c *Camera) getHeightStart() int {
	if c.top_y > 0 {
		return c.top_y
	} else {
		return 0
	}
}

func (c *Camera) getHeightEnd(m *gamemap) int {
	if c.top_y + c.height <= m.height {
		return c.top_y + c.height
	} else {
		return m.height
	}
}