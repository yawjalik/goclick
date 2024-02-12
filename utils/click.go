package utils

import "github.com/go-vgo/robotgo"

// Click infinitely within the coordinate boundaries
func Click(coords []Coord) {
	for {
		x, y := robotgo.Location()
		// y coordinates increase downwards
		t, r, b, l := coords[0], coords[1], coords[2], coords[3]
		if x >= l.x && x <= r.x && y >= t.y && y <= b.y {
			robotgo.Click()
		}
	}
}
