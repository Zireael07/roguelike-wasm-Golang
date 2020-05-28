package main

import (
	"log"
	"fmt"
)

//all the menu rendering functions are here

func (g *game) renderActionMenu(pos position){
	if !pos.isValid(g){
		return
	}

	if !g.Map.tiles[pos.X][pos.Y].explored{
		return
	}

	txt := ""
	if g.GetItemsAtPos(pos) != nil {
		txt = "Get an item"
	}

	//this loop has to be named, otherwise engine gets stuck?
	menuloop:
	for {
		g.Term.DrawColoredText(pos.X+1, pos.Y, txt, Color{0,255,255,255})
		g.Term.Flush()
		in := g.Term.PollEvent()
		//log.Printf("Event: %v", in);
		if in.mouse {
			m_pos := position{X: in.mouseX, Y: in.mouseY}
			switch in.button {
				case -1:
					// do nothing
				//left
				case 0:
					if m_pos.X >= pos.X && m_pos.X <= 20 && m_pos.Y == pos.Y {
						log.Printf("Selected get")
						// player on the same tile
						pl_pos := g.entities[0].Components["position"].(PositionComponent)
						if pl_pos.Pos.X == pos.X && pl_pos.Pos.Y == pos.Y {
							//Actually get the item
							it := g.GetItemsAtPos(pos)
							if it != nil {
								it.AddComponent("backpack", InBackpackComponent{})
								g.logMessage("Picked up item")
								//AI gets to move
								for _, e := range g.entities {
									if e != nil {
										if e.HasComponents([]string{"NPC"}) {
											g.takeTurn(e)
										}
									}
								}
							}
						}
						
						
						break menuloop
					}
					//exit either way
					break menuloop
				//right
				case 2:
					//exit the menu
					break menuloop
			}
		}
	}
	log.Printf("Exited loop")
	
}


func (g *game) renderInventory() {
	//log.Printf("render inventory...")
	items := g.GetItemsInventory()

	if len(items) < 1{
		log.Printf("No items in inventory...")
		return
	}
		

	x := 10
	y := 1
	for _, it := range items {
		name_c, _ := it.Components["name"].(NameComponent)
		g.Term.DrawColoredText(x,y, name_c.Name, Color{0,255,255,255})
		y++
	}

	//this loop has to be named, otherwise engine gets stuck?
	menuloop:
	for {
		g.Term.Flush()
		in := g.Term.PollEvent()
		//log.Printf("Event: %v", in);
		if in.mouse {
			m_pos := position{X: in.mouseX, Y: in.mouseY}
			switch in.button {
				case -1:
					// do nothing
				//left
				case 0:
					if m_pos.X >= 10 && m_pos.X <= 20 {
						id := m_pos.Y-1
						// draw submenu
						g.Term.DrawColoredText(x+8,1, "Use", Color{0,255,255,255})
						g.Term.DrawColoredText(x+8,2, "Drop", Color{0,255,255,255})
						submenuloop:
							for {
								g.Term.Flush()
								in := g.Term.PollEvent()
								if in.mouse {
									m_pos := position{X: in.mouseX, Y: in.mouseY}
									switch in.button {
									case -1:
										//do nothing
									//left
									case 0:
										if m_pos.X >= 18 && m_pos.Y <= 25 {
											//use
											if m_pos.Y == 1 {
												log.Printf("Selected use")
												if items[id].HasComponent("medkit") {
													//use the medkit
													heal := items[id].Components["medkit"].(MedkitComponent).Heal
													//heal the player
													statsComp := g.entities[0].Components["stats"].(StatsComponent)
													old_hp := g.entities[0].Components["stats"].(StatsComponent).Hp
													g.logMessage(fmt.Sprintf("Medkit healed %d damage", heal)) //, Color{255,0,0,255})
													statsComp.Hp = old_hp + heal
													//bit of a dance because we're not using pointers to Components
													g.entities[0].RemoveComponent("stats")
													g.entities[0].AddComponent("stats", statsComp)
													//nuke the medkit
													items[id].RemoveComponents([]string{"position", "renderable", "item", "medkit", "backpack"})
													//end turn
													//AI gets to move
													for _, e := range g.entities {
														if e != nil {
															if e.HasComponents([]string{"NPC"}) {
																g.takeTurn(e)
															}
														}
													}
													//break menuloop
												}
												if items[id].HasComponent("range") {
													shootloop:
														for {
															in := g.Term.PollEvent()
															//log.Printf("Event: %v", in);
															if in.mouse {
																m_pos := position{X: in.mouseX, Y: in.mouseY}
																g.Term.highlightPos(m_pos)
																g.describePosition(m_pos)
																g.Term.Flush()
																switch in.button {
																	case -1:
																		// do nothing
																	//left
																	case 0:
																		
																		//target
																		dist := items[id].Components["range"].(RangeComponent).dist
																		pl_pos := g.entities[0].Components["position"].(PositionComponent).Pos
					
																		//target in range
																		if m_pos.X <= pl_pos.X + dist && m_pos.Y <= pl_pos.Y + dist{
																			//are we targeting something?
																			blocker := g.getAllBlockers(m_pos)
																			if blocker != nil{
																				//damage it!
																				statsComp := blocker.Components["stats"].(StatsComponent)
																				damage := 6 //dummy
																				old_hp := blocker.Components["stats"].(StatsComponent).Hp
																				g.logMessage(fmt.Sprintf("Player shoots enemy for %d damage", damage)) // Color{0,255,0,255})
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
					
																				//nuke the gun
																				//we're kind, we only nuke it if we actually hit something
																				items[id].RemoveComponents([]string{"position", "renderable", "item", "range", "backpack"})
																				
																			} else {
																				g.logMessage("Nothing to shoot at here")
																			}

																			//end turn
																			//AI gets to move
																			for _, e := range g.entities {
																				if e != nil {
																					if e.HasComponents([]string{"NPC"}) {
																						g.takeTurn(e)
																					}
																				}
																			}
																		}
																		break shootloop
																	//right
																	case 2:
																		//cancel
																		break shootloop
																}
															}
														}
														break submenuloop
												}

											//drop
											} else if m_pos.Y == 2 {
												log.Printf("Select drop")
												//paranoia
												if items[id].HasComponent("backpack") {
													items[id].RemoveComponent("backpack")
													pl_pos := g.entities[0].Components["position"].(PositionComponent).Pos
													posComp := PositionComponent{Pos:position{X:pl_pos.X, Y:pl_pos.Y}}
													//bit of a dance because we're not using pointers to Components
													items[id].RemoveComponent("position")
													items[id].AddComponent("position", posComp)
													g.logMessage("Player dropped item")

													//end turn
													//AI gets to move
													for _, e := range g.entities {
														if e != nil {
															if e.HasComponents([]string{"NPC"}) {
																g.takeTurn(e)
															}
														}
													}
												}
											}
											
										}
										break submenuloop
									//right
									case 2:
										break submenuloop
									}

								}
							}

					}
					// break out of menu loop (item actions are done!)
					break menuloop
				//right
				case 2:
					//exit the menu
					break menuloop
			}
		}
	}
	log.Printf("Exited ui loop")
	g.Term.Flush()
}