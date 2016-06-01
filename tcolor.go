package tcolor

// HasColor returns whether the terminal has support for
// terminal colors.
func HasColor(name string) bool {
	// db is sorted, run a binary search
	l := 0
	r := len(db)
	for l <= r {
		m := (l + r) / 2
		switch {
		case db[m] < name:
			l = m + 1
		case name < db[m]:
			r = m - 1
		default:
			return true
		}
	}
	return false
}
