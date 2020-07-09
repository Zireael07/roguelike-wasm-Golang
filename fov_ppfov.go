// Based on Lokathor's 2018 repo: https://github.com/Lokathor/roguelike-tutorial-2018
//! Does precise permissive FOV calculations.

package main

import (
  "fmt"
  //"log"
)


//fake a hashset with a map
type HashSet  map[string]struct{}

// Returns a boolean value describing if it exists in the set
func (s HashSet) has(key string) bool {
	_, ok := s[key]
	return ok
}

// Adds to the set
func (s HashSet) add(key string) {
	s[key] = struct{}{}
}


type Line struct {
  xi int32
  yi int32
  xf int32
  yf int32
}

func (l *Line) dx() int32 {
    return l.xf - l.xi;
}

func (l *Line) dy() int32 {
    return l.yf - l.yi;
}

func (l *Line) relative_slope(x, y int32) int32 {
    return (l.dy() * (l.xf - x)) - (l.dx() * (l.yf - y));
}
 
func (l *Line) below_p(x, y int32) bool {
    return l.relative_slope(x, y) > 0
}

func (l *Line) below_or_collinear_p(x, y int32) bool {
    return l.relative_slope(x, y) >= 0
}

func (l *Line) above_p(x,y int32) bool {
    return l.relative_slope(x, y) < 0
}

func (l *Line) above_or_collinear_p(x,y int32) bool {
    return l.relative_slope(x, y) <= 0
}

func (l *Line) collinear_p(x,y int32) bool {
    return l.relative_slope(x, y) == 0
  }
func (l *Line) collinear_line(line Line) bool {
    return l.collinear_p(line.xi, line.yi) && l.collinear_p(line.xf, line.yf)
}

func (l Line) deepcopy() Line {
  n := Line{xi: l.xi, yi: l.yi, xf: l.xf, yf:l.yf}
  return n
}

type intpair struct {
	a int32
	b int32
}

type View struct {
  shallow_line Line
  steep_line Line
  shallow_bumps []intpair
  steep_bumps []intpair
}


func (v *View) add_shallow_bump(x,y int32) {
	v.shallow_line.xf = x;
  v.shallow_line.yf = y;
  v.shallow_bumps = append(v.shallow_bumps, intpair{x, y})
  // move to index 0
  copy(v.shallow_bumps[1:], v.shallow_bumps)
  v.shallow_bumps[0] = intpair{x, y}
  
	//Golang iteration
	for _, bump := range v.steep_bumps {
		if v.shallow_line.above_p(bump.a, bump.b) {
		  v.shallow_line.xi = bump.a;
		  v.shallow_line.yi = bump.b;
		}
	}
}

func (v *View) add_steep_bump(x, y int32) {
	v.steep_line.xf = x;
  v.steep_line.yf = y;
  //v.steep_bump = ViewBump{x,y, v.steep_bump}
  v.steep_bumps = append(v.steep_bumps, intpair{x, y})
  // move to index 0
  copy(v.steep_bumps[1:], v.steep_bumps)
  v.steep_bumps[0] = intpair{x,y}

	for _, bump := range v.shallow_bumps {
		if v.steep_line.below_p(bump.a, bump.b) {
		  v.steep_line.xi = bump.a;
		  v.steep_line.yi = bump.b;
		}
	}
}

//based on https://rosettacode.org/wiki/Deepcopy#Go
func (v View) deepcopy() View {
  n := View{shallow_line: v.shallow_line.deepcopy(), steep_line: v.steep_line.deepcopy(), 
    shallow_bumps: append([]intpair{}, v.shallow_bumps...),
    steep_bumps: append([]intpair{}, v.steep_bumps...)} 
  return n
}



//VB, VE, IM are functions (see fov.go)
func check_quadrant(visited HashSet, start_x, start_y, dir_x, dir_y, extent_x, extent_y int32, 
  vision_blocked VB, visit_effect VE, in_map IM) {
//   debug_assert!(dir_x == -1 || dir_x == 1);
//   debug_assert!(dir_y == -1 || dir_y == 1);
//   debug_assert!(extent_x > 0);
//   debug_assert!(extent_y > 0);

  shallow_line := Line{0, 1, extent_x, 0};
  steep_line := Line{1, 0, 0, extent_y};
  var active_views []View 
  active_views = append(active_views, View{shallow_line: shallow_line, steep_line: steep_line});

//..= means inclusive on both ends
  //for i in 1..=(extent_x + extent_y) {
	for i := 1; i <= int(extent_x + extent_y); i++ {
    //for j in (i - extent_x).max(0)..=i.min(extent_y) {
    // Golang's Min and Max work on floats so we provide our own for int32 (and then cast to int)
		for j:= int((Max((int32(i)- extent_x), int32(0)))); j <= int(Min(int32(i), extent_y)); j++ {
			if len(active_views) < 1 {
        //log.Printf("No active views")
				return;
		  } else {
        offset_x := i - j;
        offset_y := j;
        visit_coord(
        visited,
        start_x,
        start_y,
        dir_x,
        dir_y,
        vision_blocked,
        visit_effect,
        in_map,
        int32(offset_x),
        int32(offset_y),
        active_views,
        );
		  }
    }
  }
}

func visit_coord(visited HashSet, start_x, start_y, dir_x, dir_y int32, vision_blocked VB, visit_effect VE, in_map IM,
  offset_x, offset_y int32, active_views []View) {
//   debug_assert!(dir_x == -1 || dir_x == 1);
//   debug_assert!(dir_y == -1 || dir_y == 1);
//   debug_assert!(offset_x >= 0);
//   debug_assert!(offset_y >= 0);

    //log.Printf("visiting: sx %d sy %d dx %d dy %d ox %d oy %d", start_x, start_y, dir_x, dir_y, offset_x, offset_y)

  top_left := intpair{offset_x, offset_y + 1};
  bottom_right := intpair{offset_x + 1, offset_y};
  view_index := 0;
  
  //log.Printf("Active views: %d ", len(active_views))
  
  // this is an equivalent of a 'while' loop
  for view_index < len(active_views){
    view_ref := active_views[view_index]
    if view_ref.steep_line.below_or_collinear_p(bottom_right.a, bottom_right.b) {
          view_index += 1;
    } else if view_ref.shallow_line.above_or_collinear_p(top_left.a, top_left.b) {
          return;
    } else {
          break;
    }
  }

  if view_index == len(active_views){
    //coordinate is below all the fields
    return
  }
  
  target := intpair{start_x + (offset_x * dir_x), start_y + (offset_y * dir_y)};

  //prevent going out of map
  //if (in_map(target.a, target.b) == false) {
    //log.Printf("Target %d %d out of map, abort", target.a, target.b)
  //  return
  //}

  //log.Printf("target: %d %d", target.a, target.b)
  key := fmt.Sprintf("%d,%d", target.a, target.b) 
  if !visited.has(key) {
    visited.add(key);
    if (in_map(target.a, target.b) == true) {
      visit_effect(target.a, target.b);
    }
  }

  if vision_blocked(target.a, target.b) {
    if active_views[view_index].shallow_line.above_p(bottom_right.a, bottom_right.b) &&
       active_views[view_index].steep_line.below_p(top_left.a, top_left.b){
        // The shallow line and steep line both intersect this location, and
        // sight is blocked here, so this view is dead.
        //log.Printf("Case 1")
        //remove from slice
        i := view_index
        active_views = append(active_views[:i], active_views[i+1:]...)
 
    } else if active_views[view_index].shallow_line.above_p(bottom_right.a, bottom_right.b) {
        // The shallow line intersects here but the steep line does not, so we
        // add this location as a shallow bump and check our views.
        // shallow line needs to be raised
        //log.Printf("Case 2")
        active_views[view_index].add_shallow_bump(top_left.a, top_left.b);
        check_view(active_views, view_index);
      } else if active_views[view_index].steep_line.below_p(top_left.a, top_left.b) {
        // the steep line intersects here but the shallow line does not, so we
        // add a steep bump at this location and check our views.
        //log.Printf("Case 3")
        active_views[view_index].add_steep_bump(bottom_right.a, bottom_right.b);
        check_view(active_views, view_index);
      } else {
        // Neither line intersects this location but it blocks sight, so we have
        // to split this view into two views.
        
        //FIXME: I think this case doesn't work
        //log.Printf("Case 4")
        //deep copy
        // new_view := active_views[view_index]
        // new_view.shallow_bumps = make([]intpair, len(active_views[view_index].shallow_bumps))
        // copy(new_view.shallow_bumps, active_views[view_index].shallow_bumps)
        // new_view.steep_bumps = make([]intpair, len(active_views[view_index].steep_bumps))
        // copy(new_view.steep_bumps, active_views[view_index].steep_bumps)
        
        new_view := active_views[view_index].deepcopy()
        //log.Printf("%v", new_view)

        //insert the new view
        active_views = append(active_views, new_view)
        copy(active_views[view_index+1:], active_views)
        active_views[view_index] = new_view
        //log.Printf("%v", active_views[view_index])

        // We add the shallow bump on the farther view first, so that if it gets
        // killed we don't have to change where we add the steep bump and check
        active_views[view_index + 1].add_shallow_bump(top_left.a, top_left.b);
        //log.Printf("%v", active_views[view_index+1])
        check_view(active_views, view_index + 1);
        //log.Printf("Check view + 1")
        active_views[view_index].add_steep_bump(bottom_right.a, bottom_right.b);
        check_view(active_views, view_index);
        //log.Printf("%v", active_views[view_index])
        //log.Printf("Check view")
      }
    }
}

func check_view(active_views []View, view_index int) {
  //Removes the view in activeViews at index viewIndex if
  //          - The two lines are coolinear
  //          - The lines pass through either extremity
  
    shallow_line := active_views[view_index].shallow_line;
    steep_line := active_views[view_index].steep_line;
    
    if shallow_line.collinear_line(steep_line) && (shallow_line.collinear_p(0, 1) || shallow_line.collinear_p(1, 0)) {
      //remove from slice
      i := view_index
      active_views = append(active_views[:i], active_views[i+1:]...)
      //active_views.remove(view_index);
  }
}

/// Computes field of view according to the "Precise Permissive" technique.
///
/// [See the RogueBasin page](http://www.roguebasin.com/index.php?title=Precise_Permissive_Field_of_View)
//VB, VE, IM are functions
func (g *game) pp_FOV(start_x, start_y, radius int32, vision_blocked VB, visit_effect VE, in_map IM) {
//   debug_assert!(radius >= 0, "ppfov: vision radius must be non-negative, got {}", radius);
//   debug_assert!(
//     start_x.saturating_add(radius) < ::std::i32::MAX,
//     "ppfov: Location ({},{}) with radius {} would cause overflow problems!",
//     start_x,
//     start_y,
//     radius
//   );
//   debug_assert!(
//     start_y.saturating_add(radius) < ::std::i32::MAX,
//     "ppfov: Location ({},{}) with radius {} would cause overflow problems!",
//     start_x,
//     start_y,
//     radius
//   );
//   debug_assert!(
//     start_x.saturating_sub(radius) > ::std::i32::MIN,
//     "ppfov: Location ({},{}) with radius {} would cause underflow problems!",
//     start_x,
//     start_y,
//     radius
//   );
//   debug_assert!(
//     start_y.saturating_sub(radius) > ::std::i32::MIN,
//     "ppfov: Location ({},{}) with radius {} would cause underflow problems!",
//     start_x,
//     start_y,
//     radius
//   );

  visited := HashSet{};
  visit_effect(start_x, start_y);
  //visited.insert((start_x, start_y));
  var key string
  key = fmt.Sprintf("%d,%d", start_x, start_y) 
  visited[key] = struct{}{}

  // q1
  check_quadrant(visited, start_x, start_y, 1, 1, radius, radius, vision_blocked, visit_effect, in_map);
  // q2
  check_quadrant(visited, start_x, start_y, -1, 1, radius, radius, vision_blocked, visit_effect, in_map);
  // q3
  check_quadrant(visited, start_x, start_y, -1, -1, radius, radius, vision_blocked, visit_effect, in_map);
  // q4
  check_quadrant(visited, start_x, start_y, 1, -1, radius, radius, vision_blocked, visit_effect, in_map);
}