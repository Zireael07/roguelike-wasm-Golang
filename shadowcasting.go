//Golang port of the version presented at http://www.adammil.net/blog/v125_Roguelike_Vision_Algorithms.html

package main

// represents the slope Y/X as a rational number
type Slope struct {
	Y int
	X int
}

func (g *game) computeFOV(origin position, radius int, vision_blocked VB, visit_effect VE) {
	visit_effect(int32(origin.X), int32(origin.Y))
	//https://stackoverflow.com/a/13384288
	for octant := uint(0); octant < 8; octant++ {
		Compute(octant, origin, radius, 1, Slope{1,1}, Slope{0,1}, vision_blocked, visit_effect)
	}
}

func GetDistance(origin position, tx int, ty int) int {
	return origin.Distance(position{X:tx, Y:ty})
}

//helper functions because they wouldn't work inside
func inRange(radius int, tx int, ty int, origin position) bool {
	return radius < 0 || GetDistance(origin, tx, ty) <= radius
}

func isOpaque(tx int, ty int, inRange bool, vision_blocked VB) bool {
	return !inRange || vision_blocked(int32(tx),int32(ty))
}


func Compute(octant uint, origin position, radius int, x int, top Slope, bottom Slope, vision_blocked VB, visit_effect VE) {
	for x := uint(0); uint(x) <= uint(radius); x++ { // rangeLimit < 0 || x <= rangeLimit
		// compute the Y coordinates where the top vector leaves the column (on the right) and where the bottom vector
		// enters the column (on the left). this equals (x+0.5)*top+0.5 and (x-0.5)*bottom+0.5 respectively, which can
		// be computed like (x+0.5)*top+0.5 = (2(x+0.5)*top+1)/2 = ((2x+1)*top+1)/2 to avoid floating point math
		
		//int topY = top.X == 1 ? x : ((x*2+1) * top.Y + top.X - 1) / (top.X*2); // the rounding is a bit tricky, though
		//int bottomY = bottom.Y == 0 ? 0 : ((x*2-1) * bottom.Y + bottom.X) / (bottom.X*2);
		

		//no ternary operators in Golang, so this gets a little wordy
		topY := 0
		if top.X == 1 {
			topY = int(x)
		} else {
			topY = (int((x*2+1)) * top.Y + top.X - 1) / (top.X*2)
		}

		bottomY := 0
		if bottom.Y == 0 {
			bottomY = 0
		} else {
			bottomY = (int((x*2-1)) * bottom.Y + bottom.X) / (bottom.X*2)
		}

		wasOpaque := -1; // 0:false, 1:true, -1:not applicable
		
		for y := topY; y >= bottomY; y-- {
			tx := origin.X
			ty := origin.Y
			switch octant { // translate local coordinates to map coordinates
			case 0: 
				tx += int(x); ty -= int(y); break;
			case 1: 
				tx += int(y); ty -= int(x); break;
			case 2: 
				tx -= int(y); ty -= int(x); break;
			case 3: 
				tx -= int(x); ty -= int(y); break;
			case 4: 
				tx -= int(x); ty += int(y); break;
			case 5: 
				tx -= int(y); ty += int(x); break;
			case 6: 
				tx += int(y); ty += int(x); break;
			case 7: 
				tx += int(x); ty += int(y); break;
			}
			
			//NOTE: tx and ty are MAP coordinates!

			//bool inRange = rangeLimit < 0 || GetDistance(tx, ty) <= radius;
			//inRange := func (radius int, tx int, ty int) bool { return radius < 0 || GetDistance(tx, ty) <= radius }
			inRange := inRange(radius, tx, ty, origin) // for perf
			
			//non-symmetrical
			// if(inRange) {
			// 	visit_effect(int32(tx), int32(ty));
			// }
			// NOTE: use the next line instead if you want the algorithm to be symmetrical
			if(inRange && (y != topY || top.Y*int(x) >= top.X*y) && (y != bottomY || bottom.Y*int(x) <= bottom.X*y)) {
				visit_effect(int32(tx), int32(ty));
			}
			

			//bool isOpaque = !inRange || BlocksLight(tx, ty);
			//isOpaque := func (inRange bool, vision_blocked VB) bool { return !inRange || vision_blocked(tx,ty) }
			
			if(int(x) != radius) {
			if(isOpaque(tx, ty, inRange, vision_blocked)) {            
				if(wasOpaque == 0) { // if we found a transition from clear to opaque, this sector is done in this column, so
								// adjust the bottom vector upwards and continue processing it in the next column.
					newBottom := Slope{int(y*2+1), int(x*2-1)} // (x*2-1, y*2+1) is a vector to the top-left of the opaque tile
					if(!inRange || y == bottomY) { 
						bottom = newBottom; 
						break; 
						// don't recurse unless we have to
					} else { Compute(octant, origin, radius, int(x+1), top, newBottom, vision_blocked, visit_effect); }
				}
				wasOpaque = 1;
				
			} else { // adjust top vector downwards and continue if we found a transition from opaque to clear
				// (x*2+1, y*2+1) is the top-right corner of the clear tile (i.e. the bottom-right of the opaque tile)
				if(wasOpaque > 0) {
					top = Slope{y*2+1, int(x*2+1)}
				}
				wasOpaque = 0;
				}
			}
		}

		if(wasOpaque != 0) {
			break; // if the column ended in a clear tile, continue processing the current sector
		}
		
	}
}