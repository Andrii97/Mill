package main

import "testing"

func TestLineCheck(t *testing.T) {
	init_board()

	board[1][1] = WHITE
	board[1][4] = WHITE
	board[1][7] = WHITE
	if check_line(1, 1) != 1 {
		t.Error("Expected 1, got ", check_line(1, 1))
	}

	board[4][1] = WHITE
	board[7][1] = WHITE

	if check_line(1, 1) != 2 {
		t.Error("Expected 2, got ", check_line(1, 1))
	}

	init_board()
	board[4][1] = BLACK
	board[4][2] = BLACK
	board[4][3] = BLACK

	if check_line(2, 4) != 1 {
		t.Error("Expected 1, got ", check_line(4, 2))
	}

	board[3][3] = BLACK
	board[4][3] = BLACK
	board[5][3] = BLACK
	if check_line(3, 4) != 2 {
		t.Error("Expected 2, got ", check_line(4, 3))
	}

	if check_line(3, 3) != 1 {
		t.Error("Expected 1, got ", check_line(3, 3))
	}
}
