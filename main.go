package main

import (
	"fmt"
	//"io/ioutil"
	//"log"
)

var board [n][n]cell_state

type cell_state int

const (
	EMPTY        cell_state = 0
	WHITE        cell_state = 5
	BLACK        cell_state = -5
	INACCESSIBLE cell_state = -1
	BORDER       cell_state = -2
)

type board_state int 

const (
	PUT_WHITE board_state = 1
	PUT_BLACK board_state = 2
	DELETE_WHITE board_state = 3
	DELETE_BLACK board_state = 4
	MOVE_WHITE board_state = 5
	MOVE_BLACK board_state = 6
	CLOSE board_state = 7
)

const n = 9 // board_size

type coordinates struct {
  x int
  y int
}

func init_board() {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i < 1 || j < 1 || i > n-2 || j > n-2 {
				board[i][j] = BORDER
			} else if i == j || i == n-j-1 || j == n/2 || i == n/2 {
				board[i][j] = EMPTY
			} else {
				board[i][j] = INACCESSIBLE
			}
		}
	}
	board[n/2][n/2] = INACCESSIBLE
}

func print(arr [n][n]cell_state) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%4d", arr[i][j])
		}
		fmt.Println()
		fmt.Println()
	}
}

func check_line(h int, w int) int {
	res := 0
	cnt := 0
	if w != n/2 {
		for i := 1; i < n-1; i++ {
			if board[i][w] == board[h][w] {
				cnt++
			}
		}
	} else {
		if h < n/2 {
			for i := 1; i < n/2; i++ {
				if board[i][w] == board[h][w] {
					cnt++
				}
			}
		} else {
			for i := n/2 + 1; i < n-1; i++ {
				if board[i][w] == board[h][w] {
					cnt++
				}
			}
		}
	}

	if cnt == 3 {
		res++
	}
	cnt = 0

	if h != n/2 {
		for i := 1; i < n-1; i++ {
			if board[h][i] == board[h][w] {
				cnt++
			}
		}
	} else {
		if w < n/2 {
			for i := 1; i < n/2; i++ {
				if board[h][i] == board[h][w] {
					cnt++
				}
			}
		} else {
			for i := 5; i < n-1; i++ {
				if board[h][i] == board[h][w] {
					cnt++
				}
			}
		}
	}

	if cnt == 3 {
		res++
	}
	return res
}

func put_white() coordinates {
	fmt.Println("Turn white_player")
	fmt.Println("Write coordinates (x,y) to put piece")
	coord := coordinates{1, 1}
	return coord
}

func put_black() coordinates {
	fmt.Println("Turn black_player")
	fmt.Println("Write coordinates (x,y) to put piece")
	coord := coordinates{1, 1}
	return coord
}

func delete_white() coordinates {
	fmt.Println("Turn black_player")
	fmt.Println("Write coordinates (x,y) to delete_black")
	coord := coordinates{1, 1}
	return coord
}

func delete_black() coordinates {
	fmt.Println("Turn white_player")
	fmt.Println("Write coordinates (x,y) to delete_black")
	coord := coordinates{1, 1}
	return coord
}

func move_white() coordinates {
	fmt.Println("Turn black_player")
	fmt.Println("Write coordinates (x,y) of piece which you want to move")
	coord := coordinates{1, 1}
	return coord
}

func move_black() coordinates {
	fmt.Println("Turn black_player")
	fmt.Println("Write coordinates (x,y) of piece which you want to move")
	coord := coordinates{1, 1}
	return coord
}

func main() {
	init_board()
	current_state := PUT_WHITE
	for current_state != CLOSE {
		print(board)
		switch current_state {
			case PUT_WHITE:
				coord := put_white()
				if check_line(coord.x, coord.y) == 1 {
					current_state = DELETE_BLACK // todo
				}
			case PUT_BLACK:
				coord := put_black()
				if check_line(coord.x, coord.y) == 1 {
					current_state = DELETE_WHITE // todo
				}
			case DELETE_WHITE:
				delete_white()
				current_state = PUT_WHITE // todo
			case DELETE_BLACK:
				delete_black()
				current_state = PUT_BLACK // todo
			case MOVE_WHITE:
				coord := move_white()
				if check_line(coord.x, coord.y) == 1 {
					current_state = MOVE_BLACK
				}
			case MOVE_BLACK:
				coord := move_black()
				if check_line(coord.x, coord.y) == 1 { // todo for two
					current_state = MOVE_WHITE
				}
		}
		symbol := '_'
		fmt.Scanf("%c", &symbol)
		if symbol == 'q' {
			current_state = CLOSE
		}
	}
	print(board)
	fmt.Print(check_line(1, 1))
}
