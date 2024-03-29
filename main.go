package main

import "fmt"

func main() {
	fmt.Println("dcp39 - Conway's Game of Life")
	fmt.Println("Run Unit Tests!")
}

//GetNextBoard will take a gameboard in the form of live cells and return the next gameboard in the form of livecells
func GetNextBoard(liveCells [][2]int) [][2]int {
	nextBoard := [][2]int{}

	rowMin := liveCells[0][0]
	rowMax := liveCells[0][0]
	colMin := liveCells[0][1]
	colMax := liveCells[0][1]

	// First determine the board dimensions
	for i := 0; i < len(liveCells); i++ {
		r := liveCells[i][0]
		c := liveCells[i][1]

		if r < rowMin {
			rowMin = r
		}

		if r > rowMax {
			rowMax = r
		}

		if c < colMin {
			colMin = c
		}

		if c > colMax {
			colMax = c
		}
	}

	// make the next board 2 rows and 2 columns bigger than the existing board, just in case
	for r := rowMin - 1; r <= (rowMax + 1); r++ {
		for c := colMin - 1; c <= (colMax + 1); c++ {
			neighborCount := CountNeighbors(liveCells, r, c)
			if IsCellLive(liveCells, r, c) {
				// if it has < 2 live neighbor, it will die

				// if it has 2-3 live neighbors it lives
				if neighborCount >= 2 && neighborCount <= 3 {
					nextBoard = append(nextBoard, [2]int{r, c})
				}

				// if it has > 3 live neighbors it dies
			} else {
				// if it has 3 live neighbors it will be born
				if neighborCount == 3 {
					nextBoard = append(nextBoard, [2]int{r, c})
				}
			}
		}
	}

	return nextBoard
}

// IsCellLive will test if a cell exists at row,col
func IsCellLive(liveCells [][2]int, row int, col int) bool {
	for _, cell := range liveCells {
		if cell[0] == row && cell[1] == col {
			return true
		}
	}

	return false
}

// CountNeighbors will find out how many alive neighbors there are for a cell at row,col
func CountNeighbors(liveCells [][2]int, row int, col int) int {
	cnt := 0

	for i := 0; i < len(liveCells); i++ {
		if CellsAreNeighbors(row, col, liveCells[i][0], liveCells[i][1]) {
			cnt++
		}
	}

	return cnt
}

// CellsAreNeighbors will test to see if two cell coordinates are neighboring each other
func CellsAreNeighbors(r1 int, c1 int, r2 int, c2 int) bool {
	ret := false
	if r1 == r2 && c1 == c2 {
		return false
	}

	if ((r1 - r2) >= -1) && ((r1 - r2) <= 1) {
		if ((c1 - c2) >= -1) && ((c1 - c2) <= 1) {
			return true
		}
	}

	return ret
}

// given a board represented by boolean values
func PrintBoard(liveCells [][2]int) {

	rowMin := liveCells[0][0]
	rowMax := liveCells[0][0]
	colMin := liveCells[0][1]
	colMax := liveCells[0][1]

	for i := 0; i < len(liveCells); i++ {
		r := liveCells[i][0]
		c := liveCells[i][1]

		if r < rowMin {
			rowMin = r
		}

		if r > rowMax {
			rowMax = r
		}

		if c < colMin {
			colMin = c
		}

		if c > colMax {
			colMax = c
		}
	}

	width := colMax - colMin + 1
	height := rowMax - rowMin + 1

	board := make([][]bool, height)
	for y := 0; y < height; y++ {
		board[y] = make([]bool, width)
	}

	for i := 0; i < len(liveCells); i++ {
		r := liveCells[i][0]
		c := liveCells[i][1]

		board[r-rowMin][c-colMin] = true
	}

	fmt.Printf("height=%d  width=%d\n", height, width)

	fmt.Println(board)

	// now we have a boolean representation of the board
	for row := rowMin; row <= rowMax; row++ {
		for col := colMin; col <= colMax; col++ {
			if board[row-rowMin][col-colMin] {
				fmt.Print("*")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

// given a board represented by boolean values
func PrintCustomBoard(liveCells [][2]int, top int, bottom int, left int, right int) {

	width := right - left + 1
	height := bottom - top + 1

	board := make([][]bool, height)
	for y := 0; y < height; y++ {
		board[y] = make([]bool, width)
	}

	for i := 0; i < len(liveCells); i++ {
		r := liveCells[i][0]
		c := liveCells[i][1]

		if r >= top && r <= bottom && c >= left && c <= right {
			board[r-top][c-left] = true
		}

	}

	//fmt.Printf("height=%d  width=%d\n", height, width)

	//fmt.Println(board)

	// now we have a boolean representation of the board
	for row := top; row <= bottom; row++ {
		for col := left; col <= right; col++ {
			if board[row-top][col-left] {
				fmt.Print("*")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
