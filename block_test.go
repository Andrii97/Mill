package main

import "testing"

func TestBlock(t *testing.T) {
	init_board()
	board[1][1] = WHITE
	board[1][4] = WHITE

	if isBlocked(1, 1) {
		t.Error("Expected false, got ", isBlocked(1, 1))
	}

	board[4][1] = BLACK

	if !isBlocked(1, 1) {
		t.Error("Expected true, got ", isBlocked(1, 1))
	}

	board[4][3] = BLACK
	board[3][3] = WHITE

	if isBlocked(3, 4) {
		t.Error("Expected false, got ", isBlocked(3, 4))
	}

	board[5][3] = WHITE
	board[4][2] = BLACK

	if !isBlocked(3, 4) {
		t.Error("Expected true, got ", isBlocked(3, 4))
	}
}
