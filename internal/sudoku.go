package sodoku

import "fmt"

const (
	Unset uint8 = 0
	N           = 9
)

type Grid [N][N]uint8

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

func Solve(g *Grid) bool {
	return solve(g, 0, 0)
}

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

outer:
	for i := 1; i <= N; i++ {
		val := uint8(i)

		// validate row
		for _, c := range g[row] {
			if c == val {
				continue outer
			}
		}

		// validate column
		for j := 0; j < N; j++ {
			c := g[j][col]
			if c == val {
				continue outer
			}
		}

		// validate box
		j := int(row/3) * 3
		k := int(col/3) * 3
		for ji, jn := j, j+3; ji < jn; ji++ {
			for ki, kn := k, k+3; ki < kn; ki++ {
				c := g[ji][ki]
				if c == val {
					continue outer
				}
			}
		}

		// Backtracking... Solution

		// set value
		g[row][col] = val

		// recursively attempt to solve the rest of the cells.
		if solve(g, row, col+1) {
			return true
		}

		// unset value
		g[row][col] = Unset
	}

	return false
}
