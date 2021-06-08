package main

import (
	cw "duskerscli/console_wrapper"
	"strconv"
	"strings"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func euclideanDistance(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func toUnitVector(vx, vy int) (int, int) {
	if vx != 0 {
		vx = vx/abs(vx)
	}
	if vy != 0 {
		vy = vy/abs(vy)
	}
	return vx, vy
}

func stringBeginsWith(str, beg string) bool {
	return strings.Index(strings.ToLower(str), strings.ToLower(beg)) == 0
}

func drawIntegerMap(arr [][]int) {
	cw.SetFgColor(cw.CYAN)
	for x := range arr {
		for y := range arr[x] {
			cw.PutString(strconv.Itoa(arr[x][y]), rend.roomSizeX*4+1+x, y)
		}
	}
	cw.Flush_console()
}

func createFloodFilledIntegerMap(sizex, sizey, integers int) [][]int {
	arr := make([][]int, sizex)
	for i := range arr {
		arr[i] = make([]int, sizey)
	}
	zeros := sizex*sizey
	for currInt := 1; currInt <= integers; currInt++ {
		for {
			x, y := rnd.Rand(sizex), rnd.Rand(sizey)
			if arr[x][y] == 0 {
				arr[x][y] = currInt
				zeros--
				break
			}
		}
	}
	// now "grow" all the numbers
	for zeros > 0 {
		for currInt := 1; currInt <= integers; currInt++ {
			growables := make([][2]int, 0)
			for x := 0; x < sizex; x++ {
				for y := 0; y < sizey; y++ {
					if arr[x][y] == currInt {
						if x < sizex-1 && arr[x+1][y] == 0 {
							growables = append(growables, [2]int{x+1, y})
						}
						if x > 0 && arr[x-1][y] == 0 {
							growables = append(growables, [2]int{x-1, y})
						}
						if y < sizey-1 && arr[x][y+1] == 0 {
							growables = append(growables, [2]int{x, y+1})
						}
						if y > 0 && arr[x][y-1] == 0 {
							growables = append(growables, [2]int{x, y-1})
						}
					}
				}
			}
			if len(growables) == 0 {
				continue
			}
			coords := growables[rnd.Rand(len(growables))]
			arr[coords[0]][coords[1]] = currInt
			zeros--
		}
	}
	return arr
}
