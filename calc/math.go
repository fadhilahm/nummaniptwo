package calc

// Add function will run addition process
func Add (numbers ...int) int {
	sum := 0

	for _, num:= range numbers {
		sum = sum + num
	}

	return sum
}