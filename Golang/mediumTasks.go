package main

// 1041. Robot Bounded In Circle
// https://leetcode.com/problems/robot-bounded-in-circle/
func isRobotBounded(instructions string) bool {

	// 0, 1 - coordinates,
	// 2 - direction (0 - north, 1 - east, 2 - south, 3 - west)
	var position [3]int

	for _, instruction := range instructions {

		switch instruction {
		case 'G':
			switch position[2] {
			case 0:
				position[0]++
			case 1:
				position[1]++
			case 2:
				position[0]--
			case 3:
				position[1]--
			}

		case 'L':
			if position[2] < 3 {
				position[2]++
			} else {
				position[2] = 0
			}
		case 'R':
			if position[2] > 0 {
				position[2]--
			} else {
				position[2] = 3
			}
		}
	}

	return position[0] == 0 && position[1] == 0 || position[2] != 0
}
