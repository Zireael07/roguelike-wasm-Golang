package main

//interface for Components so that any old struct can be recognized as a Component
type Component interface {
	IsAIComponent() bool
}

//components
//NOTE: For Lua data interop to work, struct field names must be capitalized to be accessible
//same goes for function names

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
	Hp int
	Max_hp int
	Power int
}

func (sc StatsComponent) IsAIComponent() bool {
	return false
}

type NameComponent struct {
	Name string
}

func (nc NameComponent) IsAIComponent() bool {
	return false
}

type ItemComponent struct {
}

func (ic ItemComponent) IsAIComponent() bool {
	return false
}

type InBackpackComponent struct {

}

func (ic InBackpackComponent) IsAIComponent() bool {
	return false
}

type MedkitComponent struct {
	Heal int
}

func (mc MedkitComponent) IsAIComponent() bool {
	return false
}

type RangeComponent struct {
	dist int
}

func (rc RangeComponent) IsAIComponent() bool {
	return false
}

// Player Component
type PlayerComponent struct {
}

func (pl PlayerComponent) IsAIComponent() bool {
	return false
}