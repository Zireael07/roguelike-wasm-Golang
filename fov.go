package main


//First class functions allow the FOV code internals to not care about the map struct at all
// Based on Lokathor's 2018 repo: https://github.com/Lokathor/roguelike-tutorial-2018
type VB func (int32, int32) bool // VB = 'vision_blocked'
type VE func (int32, int32) //VE = 'visit_effect'
//is it in map?
type IM func (int32, int32) bool

type FOV interface {
	computeFOV(origin position, radius int, vision_blocked VB, visit_effect VE)
}