package main

import "testing"

func TestLineCheck(t *testing.T) {
	init_board()

	board[1][1] = 5
	board[1][4] = 5
	board[1][7] = 5
	if check_line(1, 1) != 1 {
		t.Error("Expected 1, got ", check_line(1, 1))
	}

	board[4][1] = 5
	board[7][1] = 5

	if check_line(1, 1) != 2 {
		t.Error("Expected 2, got ", check_line(1, 1))
	}

	init_board()
	board[4][1] = -5
	board[4][2] = -5
	board[4][3] = -5

	if check_line(4, 2) != 1 {
		t.Error("Expected 1, got ", check_line(4, 2))
	}

	board[3][3] = -5
	board[4][3] = -5
	board[5][3] = -5
	if check_line(4, 3) != 2 {
		t.Error("Expected 2, got ", check_line(4, 3))
	}

	if check_line(3, 3) != 1 {
		t.Error("Expected 1, got ", check_line(3, 3))
	}
}
