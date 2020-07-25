// run N-queen problems

package main

const (
	maxQueens uint8 = 8
	printOut  bool  = true
)

var (
	queens  [maxQueens]uint8
	counter uint64
)

func main() {
	println("Calculating", maxQueens, "queens puzzle...")
	placeQueen(0)
	println(counter, " result(s)")
}

func placeQueen(columnPos uint8) {
	if columnPos >= maxQueens {
		counter++
		if printOut {
			printBoard()
		}
	} else {
		for i := uint8(0); i < maxQueens; i++ {
			if verifyPos(columnPos, i) {
				queens[columnPos] = i
				placeQueen(columnPos + 1)
			}
		}
	}
}

func verifyPos(checkPos uint8, newPos uint8) bool {
	for i := uint8(0); i < checkPos; i++ {
		if (queens[i] == newPos) || (abs(int16(checkPos)-int16(i)) == abs(int16(queens[i])-int16(newPos))) {
			return false
		}
	}
	return true
}

func printBoard() {
	print(counter, " [")
	for i := uint8(0); i < maxQueens; i++ {
		print(queens[i] + 1)
		if i < maxQueens-1 {
			print(", ")
		}
	}
	print("]")
	println("")
}

func abs(x int16) int16 {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
