package main

import "slices"

func main() {

}
func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	// robot
	type robot struct {
		hp  int // health
		dir int // direction: -1 is left, 1 is right
		pos int // position
		idx int // initial health index
	}

	var (
		n      = len(positions)
		robots = make([]robot, 0, n) // line
		passed = make([]robot, 0, n) // robots moving to the left and passed all robots moving to the right
		stack  = make([]robot, 0, n) // robots moving to the right
		sIdx   int
	)

	// create robots
	for i, pos := range positions {
		dir := 1
		if directions[i] == 'L' {
			dir = -1
		}
		robots = append(robots, robot{healths[i], dir, pos, i})
	}

	// sort robots (place them on the line correctly)
	// This way collisions can occur only with neighbours
	slices.SortFunc(robots, func(a, b robot) int { return a.pos - b.pos })

	// collide robots
	for i := 0; i < len(robots); i++ {
		if robots[i].dir > 0 {
			// robot moves to the right
			// push it to stack
			if sIdx == len(stack) {
				stack = append(stack, robots[i])
			} else {
				stack[sIdx] = robots[i]
			}
			sIdx++
		} else {
			// robot moves to the left
			// collide it with robot from stack
			switch {
			case sIdx == 0:
				// no robots in stack
				// this robot passes
				passed = append(passed, robots[i])
			case robots[i].hp > stack[sIdx-1].hp:
				// this robot has more hp
				// it survives and robot from the stack is dead
				robots[i].hp--
				sIdx--
				i-- // need to process this robot again in case stack still has robots
			case robots[i].hp < stack[sIdx-1].hp:
				// robot from stack has more hp
				// this robot is dead
				stack[sIdx-1].hp--
			default:
				// both dead
				sIdx--
			}
		}
	}

	// append robots moving right to robots moving left
	// they all survive and cannot collide anymore
	passed = append(passed, stack[:sIdx]...)

	if len(passed) == 0 {
		// all robots are dead
		return []int{}
	}

	// sort robots according to the initial order
	slices.SortFunc(passed, func(a, b robot) int { return a.idx - b.idx })
	// get their healths
	var newHealths = make([]int, len(passed))
	for i, rb := range passed {
		newHealths[i] = rb.hp
	}

	return newHealths
}
