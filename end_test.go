package main

import "testing"

func TestEndingCheck(t *testing.T) {
	init_board()

	board[1][1] = WHITE
	board[1][4] = WHITE
	board[1][7] = WHITE

	if check_end(WHITE) != 0 {
		t.Error("Expected 0, got ", check_end(WHITE))
	}

	if check_end(BLACK) != 1 {
		t.Error("Expected 1, got ", check_end(BLACK))
	}

	init_board()

	board[1][7] = BLACK
	board[4][7] = BLACK
	board[7][7] = BLACK

	if check_end(WHITE) != 2 {
		t.Error("Expected 2, got ", check_end(WHITE))
	}

	if check_end(BLACK) != 0 {
		t.Error("Expected 0, got ", check_end(BLACK))
	}

	init_board()

	board[1][7] = BLACK
	board[4][7] = BLACK
	board[4][1] = WHITE
	board[4][2] = WHITE

	if check_end(WHITE) != 2 {
		t.Error("Expected 2, got ", check_end(WHITE))
	}

	if check_end(BLACK) != 1 {
		t.Error("Expected 1, got ", check_end(BLACK))
	}

	board[7][7] = WHITE
	board[4][3] = BLACK

	if check_end(WHITE) != 0 {
		t.Error("Expected 0, got ", check_end(WHITE))
	}

	if check_end(BLACK) != 0 {
		t.Error("Expected 0, got ", check_end(BLACK))
	}

	init_board()

	board[1][1] = BLACK
	board[1][4] = WHITE
	board[4][1] = WHITE

	if check_end(BLACK) != 1 {
		t.Error("Expected 1, got ", check_end(BLACK))
	}

	board[1][7] = BLACK
	board[2][4] = WHITE
	board[3][4] = BLACK
	board[7][1] = BLACK
	board[4][2] = BLACK
	board[2][2] = BLACK
	board[2][6] = BLACK

	if check_end(WHITE) != 2 {
		t.Error("Expected 2, got ", check_end(WHITE))
	}
}
