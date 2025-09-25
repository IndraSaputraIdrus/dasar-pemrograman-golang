package library2

func Sum(values []int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}
