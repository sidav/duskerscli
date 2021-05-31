package main

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func euclideanDistance(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}
