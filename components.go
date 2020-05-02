package main

//interface for Components so that any old struct can be recognized as a Component
type Component interface {
	IsAIComponent() bool
}

//components

// Position Component
type PositionComponent struct {
	Pos position //bit of indirection, but lets us take advantage of position.Distance
}

func (pc PositionComponent) IsAIComponent() bool {
	return false
}

type RenderableComponent struct {
	Color Color
	Glyph rune
}

func (rc RenderableComponent) IsAIComponent() bool {
	return false
}

type BlockerComponent struct {
}

func (bc BlockerComponent) IsAIComponent() bool {
	return false
}

type NPCComponent struct {

}

func (npc NPCComponent) IsAIComponent() bool {
	return true
}

type StatsComponent struct {
	hp int
	max_hp int
	power int
}

func (sc StatsComponent) IsAIComponent() bool {
	return false
}

type NameComponent struct {
	name string
}

func (nc NameComponent) IsAIComponent() bool {
	return false
}

// Player Component
type PlayerComponent struct {
}

func (pl PlayerComponent) IsAIComponent() bool {
	return false
}