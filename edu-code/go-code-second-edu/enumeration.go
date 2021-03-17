package main //열거형

func main() {
	const (
		Sunday = 0
		Monday = 1
		Tuesday = 2
		Thursday = 3
		Friday = 4
		Saturday = 5
	)

	// same code.

	const (
		Sundays = iota
		Mondays
		Tuesdays 
		Thursdays 
		Fridays
		Saturdays 
	)
}