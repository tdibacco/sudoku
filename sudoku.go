package sudoku

import "fmt"

const (
	Unset uint8 = 0
	N           = 9
)

type Grid [N][N]uint8

func Valid(g *Grid) bool {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			val := g[i][j]
			if val == Unset {
				continue
			}
			g[i][j] = Unset
			valid := validMove(g, i, j, val)
			g[i][j] = val
			if !valid {
				return false
			}
		}
	}
	return true
}

func Solve(g *Grid) bool { return solve(g, 0, 0) }

func solve(g *Grid, row, col int) bool {
	if row == N {
		return true
	}

	if col == N {
		return solve(g, row+1, 0)
	}

	if g[row][col] != Unset {
		return solve(g, row, col+1)
	}

	for i := 1; i <= N; i++ {
		val := uint8(i)

		if !validMove(g, row, col, val) {
			continue
		}

		// set value
		g[row][col] = val

		// recursively attempt to solve the rest of the cells
		if solve(g, row, col+1) {
			return true
		}
		// else backtrack
		g[row][col] = Unset
	}

	return false
}

func validMove(g *Grid, row, col int, val byte) bool {
	// validate row
	if rowContains(g, row, val) {
		return false
	}

	// validate column
	if colContains(g, col, val) {
		return false
	}

	// validate box
	if boxContains(g, row, col, val) {
		return false
	}

	return true
}

func rowContains(g *Grid, row int, val byte) bool {
	for _, cell := range g[row] {
		if cell == val {
			return true
		}
	}
	return false
}

func colContains(g *Grid, col int, val byte) bool {
	for row := 0; row < N; row++ {
		cell := g[row][col]
		if cell == val {
			return true
		}
	}
	return false
}

func boxContains(g *Grid, row, col int, val byte) bool {
	r := int(row/3) * 3
	c := int(col/3) * 3
	for i, m := r, r+3; i < m; i++ {
		for j, n := c, c+3; j < n; j++ {
			cell := g[i][j]
			if cell == val {
				return true
			}
		}
	}
	return false
}

func Print(g *Grid) {
	for i, row := range g {
		if i != 0 {
			println()
			if i%3 == 0 {
				println()
			}
		}

		for j, cell := range row {
			if j != 0 && j%3 == 0 {
				print("  ")
			}
			fmt.Printf("%d  ", cell)
		}
	}
}
