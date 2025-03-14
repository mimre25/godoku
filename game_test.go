package main

import "testing"

func TestNewGame(t *testing.T) {
	g := NewGame()
	for i := range 9 {
		for j := range 9 {
			if g.rows[i][j] != g.columns[j][i] {
				t.Errorf("Row %d item %d not equal to column %d item %d\n", i, j, j, i)
			}
		}
	}
	if g.blocks[0][0] != g.rows[0][0] ||
		g.blocks[0][1] != g.rows[0][1] ||
		g.blocks[0][2] != g.rows[0][2] ||
		g.blocks[0][3] != g.rows[1][0] ||
		g.blocks[0][4] != g.rows[1][1] ||
		g.blocks[0][5] != g.rows[1][2] ||
		g.blocks[0][6] != g.rows[2][0] ||
		g.blocks[0][7] != g.rows[2][1] ||
		g.blocks[0][8] != g.rows[2][2] {
		t.Error("Error in block 1")
	}
	if g.blocks[1][0] != g.rows[0][3] ||
		g.blocks[1][1] != g.rows[0][4] ||
		g.blocks[1][2] != g.rows[0][5] ||
		g.blocks[1][3] != g.rows[1][3] ||
		g.blocks[1][4] != g.rows[1][4] ||
		g.blocks[1][5] != g.rows[1][5] ||
		g.blocks[1][6] != g.rows[2][3] ||
		g.blocks[1][7] != g.rows[2][4] ||
		g.blocks[1][8] != g.rows[2][5] {
		t.Error("Error in block 2")
	}
	if g.blocks[2][3] != g.rows[1][6] ||
		g.blocks[2][4] != g.rows[1][7] ||
		g.blocks[2][5] != g.rows[1][8] ||
		g.blocks[2][0] != g.rows[0][6] ||
		g.blocks[2][1] != g.rows[0][7] ||
		g.blocks[2][2] != g.rows[0][8] ||
		g.blocks[2][6] != g.rows[2][6] ||
		g.blocks[2][7] != g.rows[2][7] ||
		g.blocks[2][8] != g.rows[2][8] {
		t.Error("Error in block 3")
	}

	if g.blocks[3][0] != g.rows[3][0] ||
		g.blocks[3][1] != g.rows[3][1] ||
		g.blocks[3][2] != g.rows[3][2] ||
		g.blocks[3][3] != g.rows[4][0] ||
		g.blocks[3][4] != g.rows[4][1] ||
		g.blocks[3][5] != g.rows[4][2] ||
		g.blocks[3][6] != g.rows[5][0] ||
		g.blocks[3][7] != g.rows[5][1] ||
		g.blocks[3][8] != g.rows[5][2] {
		t.Error("Error in block 9")
	}
	if g.blocks[4][0] != g.rows[3][3] ||
		g.blocks[4][1] != g.rows[3][4] ||
		g.blocks[4][2] != g.rows[3][5] ||
		g.blocks[4][3] != g.rows[4][3] ||
		g.blocks[4][4] != g.rows[4][4] ||
		g.blocks[4][5] != g.rows[4][5] ||
		g.blocks[4][6] != g.rows[5][3] ||
		g.blocks[4][7] != g.rows[5][4] ||
		g.blocks[4][8] != g.rows[5][5] {
		t.Error("Error in block 9")
	}

	if g.blocks[5][0] != g.rows[3][6] ||
		g.blocks[5][1] != g.rows[3][7] ||
		g.blocks[5][2] != g.rows[3][8] ||
		g.blocks[5][3] != g.rows[4][6] ||
		g.blocks[5][4] != g.rows[4][7] ||
		g.blocks[5][5] != g.rows[4][8] ||
		g.blocks[5][6] != g.rows[5][6] ||
		g.blocks[5][7] != g.rows[5][7] ||
		g.blocks[5][8] != g.rows[5][8] {
		t.Error("Error in block 6")
	}

	if g.blocks[6][0] != g.rows[6][0] ||
		g.blocks[6][1] != g.rows[6][1] ||
		g.blocks[6][2] != g.rows[6][2] ||
		g.blocks[6][3] != g.rows[7][0] ||
		g.blocks[6][4] != g.rows[7][1] ||
		g.blocks[6][5] != g.rows[7][2] ||
		g.blocks[6][6] != g.rows[8][0] ||
		g.blocks[6][7] != g.rows[8][1] ||
		g.blocks[6][8] != g.rows[8][2] {
		t.Error("Error in block 9")
	}

	if g.blocks[7][0] != g.rows[6][3] ||
		g.blocks[7][1] != g.rows[6][4] ||
		g.blocks[7][2] != g.rows[6][5] ||
		g.blocks[7][3] != g.rows[7][3] ||
		g.blocks[7][4] != g.rows[7][4] ||
		g.blocks[7][5] != g.rows[7][5] ||
		g.blocks[7][6] != g.rows[8][3] ||
		g.blocks[7][7] != g.rows[8][4] ||
		g.blocks[7][8] != g.rows[8][5] {
		t.Error("Error in block 9")
	}

	if g.blocks[8][0] != g.rows[6][6] ||
		g.blocks[8][1] != g.rows[6][7] ||
		g.blocks[8][2] != g.rows[6][8] ||
		g.blocks[8][3] != g.rows[7][6] ||
		g.blocks[8][4] != g.rows[7][7] ||
		g.blocks[8][5] != g.rows[7][8] ||
		g.blocks[8][6] != g.rows[8][6] ||
		g.blocks[8][7] != g.rows[8][7] ||
		g.blocks[8][8] != g.rows[8][8] {
		t.Error("Error in block 9")
	}
}

func TestEvaluateOptions(t *testing.T) {
	g := NewGame()

	for i := range 81 {
		if g._data[i].options != 0b1111111110 {
			t.Errorf("Incorect evaluation after new game at index %d, %b", i, g.rows[0][i].options)
		}
	}
	g.rows[0][0].num = 9
	g.rows[0][1].num = 6

	evaluateOptions(g.rows[0])
	if g.rows[0][0].options != 0b0 {
		t.Errorf("Incorect evaluation of square [0,0]. Should be 0, but is:  %b", g.rows[0][0].options)
	}
	if g.rows[0][1].options != 0b0 {
		t.Errorf("Incorect evaluation of square [0,1]. Should be 0, but is:  %b", g.rows[0][1].options)
	}
	for i := range 7 {
		if g.rows[0][i+2].options != 0b0110111110 {
			t.Errorf("Incorect evaluation after setting for square [0,%d]. %b", i+2, g.rows[0][i].options)
		}
	}

}
func TestEvaluateNumbers(t *testing.T) {
	g := NewGame()

	for i := range 9 {
		g.rows[0][i].options = 1 << (i + 1)
	}
	evaluateNumbers(g.rows[0])

	if g.rows[0][0].num != 1 {
		t.Errorf("Number not set for [0, %d]: %d", 1, g.rows[0][1].num)
	}
}
