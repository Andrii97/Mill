package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
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
	PUT_WHITE    board_state = 1
	PUT_BLACK    board_state = 2
	DELETE_WHITE board_state = 3
	DELETE_BLACK board_state = 4
	MOVE_WHITE   board_state = 5
	MOVE_BLACK   board_state = 6
	CLOSE        board_state = 7
)

const n = 9 // board_size

type coordinates struct {
	x int
	y int
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
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
	board[n/2][n/2] = BORDER
}

func print(arr [n][n]cell_state) {
	fmt.Printf("    ")

	for i := 1; i < n - 1; i++ {
		fmt.Printf("%4d", i)
	}
	fmt.Printf("\n")
	for i := 1; i < n - 1; i++ {
		fmt.Printf("%4d", i)
		for j := 1; j < n - 1; j++ {
			a := arr[i][j]
			if (a == BORDER || a == INACCESSIBLE) {
				fmt.Printf("    ")
			} else {
				fmt.Printf("%4d", arr[i][j])
			}
		}
		fmt.Println()
		fmt.Println()
	}
}

func check_line(w int, h int) int {
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
			for i := n/2 + 1; i < n-1; i++ {
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

func put(current_player cell_state) coordinates {
	var color string
	if current_player == WHITE {
		color = "White"
	} else {
		color = "Black"
	}
	var x, y int = 0, 0
	flag := true
	for flag {
		fmt.Println(color + "Player turn")
		fmt.Println("Write coordinates (x,y) to put piece")
		fmt.Print("X: ")
		fmt.Scanf("%d", &x)
		fmt.Print("Y: ")
		fmt.Scanf("%d", &y)
		if x >= n || y >= n || x < 0 || y < 0 || board[y][x] != EMPTY {
			fmt.Println("Wrong coordinates")
		} else {
			flag = false
		}
	}
	board[y][x] = current_player
	coord := coordinates{x, y}
	return coord
}

func delete(current_opponent cell_state) coordinates {
	var color string
	if current_opponent == BLACK {
		color = "White"
	} else {
		color = "Black"
	}
	var x, y int = 0, 0
	flag := true
	for flag {
		fmt.Println(color + "Player turn")
		fmt.Println("Write coordinates (x,y) to delete opponent piece")
		fmt.Print("X: ")
		fmt.Scanf("%d", &x)
		fmt.Print("Y: ")
		fmt.Scanf("%d", &y)
		if x >= n || y >= n || x < 0 || y < 0 || board[y][x] != current_opponent {
			fmt.Println("Wrong coordinates")
		} else {
			flag = false
		}
	}
	board[y][x] = EMPTY
	coord := coordinates{x, y}
	return coord
}

func move(current_player cell_state) coordinates {
	var x, y, new_x, new_y int
	command := ""
	var color string

	if current_player == WHITE {
		color = "White"
	} else {
		color = "Black"
	}

	flag := true
	for flag {
		fmt.Println(color + "Player turn")
		fmt.Println("Write coordinates (x,y) of piece which you want to move")
		fmt.Print("X: ")
		fmt.Scanf("%d", &x)
		fmt.Print("Y: ")
		fmt.Scanf("%d", &y)
		if x >= n || y >= n || x < 0 || y < 0 || board[y][x] != current_player {
			fmt.Println("Wrong coordinates")
		} else {
			flag = false
		}
	}
OUTER:
	for strings.Compare(command, "yes") != 0 {
		fmt.Println(color + "Player turn")
		fmt.Println("Write where you want to move your piece (up, down, right, left) or type 'change' if you want to choose other piece")
		way := ""
		fmt.Scanf("%s", &way)
		switch way {
		case "up":
			for i := 1; i <= y; i++ {
				if board[y-i][x] == BLACK || board[y-i][x] == WHITE {
					fmt.Println("Cell is occupied")
					continue OUTER
				} else if board[y-i][x] == BORDER || (y-i == n/2 && x == n/2) {
					fmt.Println("Wrong way")
					continue OUTER
				} else if board[y-i][x] == EMPTY {
					new_y = y - i
					new_x = x
					break
				}

			}
		case "down":
			for i := 1; i <= n-y; i++ {
				if board[y+i][x] == BLACK || board[y+i][x] == WHITE {
					fmt.Println("Cell is occupied")
					continue OUTER
				} else if board[y+i][x] == BORDER || (y+i == n/2 && x == n/2) {
					fmt.Println("Wrong way")
					continue OUTER
				} else if board[y+i][x] == EMPTY {
					new_y = y + i
					new_x = x
					break
				}

			}
		case "right":
			for i := 1; i <= n-x; i++ {
				if board[y][x+i] == BLACK || board[y][x+i] == WHITE {
					fmt.Println("Cell is occupied")
					continue OUTER
				} else if board[y][x+i] == BORDER || (y == n/2 && x+i == n/2) {
					fmt.Println("Wrong way")
					continue OUTER
				} else if board[y][x+i] == EMPTY {
					new_x = x + i
					new_y = y
					break
				}
			}
		case "left":
			for i := 1; i <= x; i++ {
				if board[y][x-i] == BLACK || board[y][x-i] == WHITE {
					fmt.Println("Cell is occupied")
					continue OUTER
				} else if board[y][x-i] == BORDER || (y == n/2 && x-i == n/2) {
					fmt.Println("Wrong way")
					continue OUTER
				} else if board[y][x-i] == EMPTY {
					new_x = x - i
					new_y = y
					break
				}
			}
		case "change":
			return move(current_player)
		default:
			fmt.Println("Wrong way")
			continue OUTER
		}

		board[new_y][new_x] = current_player //To show the result of the move (not necessary)
		board[y][x] = 0
		print(board)

		fmt.Println(color + "Player turn")
		fmt.Println("Type 'yes' to confirm your move")

		fmt.Scanf("%s", &command)
		if strings.Compare(command, "yes") != 0 {
			board[new_y][new_x] = 0 //Recovering board state if the move is not confirmed
			board[y][x] = current_player
			print(board)
		}
	}

	coord := coordinates{new_x, new_y}
	return coord
}

func is_blocked(x int, y int) bool {
	cnt := 0
	i := 1

	for board[y][x+i] == INACCESSIBLE {
		i++
	}
	if board[y][x+i] == EMPTY {
		cnt++
	}
	i = 1

	for board[y][x-i] == INACCESSIBLE {
		i++
	}
	if board[y][x-i] == EMPTY {
		cnt++
	}
	i = 1

	for board[y+i][x] == INACCESSIBLE {
		i++
	}
	if board[y+i][x] == EMPTY {
		cnt++
	}
	i = 1

	for board[y-i][x] == INACCESSIBLE {
		i++
	}
	if board[y-i][x] == EMPTY {
		cnt++
	}
	i = 1

	if cnt == 0 {
		return true
	} else {
		return false
	}
}

func check_end(next_player cell_state) bool {
	pieces_cnt := 0
	moves_cnt := 0

	for i := 1; i < n-1; i++ {
		for j := 1; j < n-1; j++ {
			if board[i][j] == next_player {
				pieces_cnt++
				if !is_blocked(j, i) {
					moves_cnt++
				}
			}
		}
	}

	if pieces_cnt < 3 || moves_cnt == 0 {
		return true
	}
	return false
}

func main() {
	clear()
	init_board()
	current_state := PUT_WHITE
	number_of_piece := 9
	for current_state != CLOSE {
		print(board)
		if check_end(WHITE) && current_state != PUT_BLACK && current_state != PUT_WHITE && 
		current_state != DELETE_BLACK && current_state != DELETE_WHITE{
			fmt.Println("Black player has won. The game will be restarted.")
			init_board()
			print(board)
			current_state = PUT_WHITE
		} else if check_end(BLACK) && current_state != PUT_BLACK && current_state != PUT_WHITE &&
		 current_state != DELETE_BLACK && current_state != DELETE_WHITE{
			fmt.Println("White player has won. The game will be restarted.")
			init_board()
			print(board)
			current_state = PUT_WHITE
		}
		switch current_state {
		case PUT_WHITE:
			coord := put(WHITE)
			cnt := check_line(coord.x, coord.y)
			if cnt > 0 {
				if cnt == 2 {
					clear()
					print(board)
					delete(BLACK)
				}
				current_state = DELETE_BLACK
			} else {
				if number_of_piece > 0 {
					current_state = PUT_BLACK
				} else {
					current_state = MOVE_BLACK
				}
			}
		case PUT_BLACK:
			coord := put(BLACK)
			cnt := check_line(coord.x, coord.y)
			number_of_piece--
			if cnt > 0 {
				if cnt == 2 {
					clear()
					print(board)
					delete(WHITE)
				}
				current_state = DELETE_WHITE
			} else {
				if number_of_piece > 0 {
					current_state = PUT_WHITE
				} else {
					current_state = MOVE_WHITE
				}
			}
		case DELETE_WHITE:
			delete(WHITE)
			if number_of_piece > 0 {
				current_state = PUT_WHITE
			} else {
				current_state = MOVE_WHITE
			}
		case DELETE_BLACK:
			delete(BLACK)
			if number_of_piece > 0 {
				current_state = PUT_BLACK
			} else {
				current_state = MOVE_BLACK
			}
		case MOVE_WHITE:
			coord := move(WHITE)
			cnt := check_line(coord.x, coord.y)
			if cnt > 0 {
				if cnt == 2 {
					clear()
					print(board)
					delete(BLACK)
				}
				current_state = DELETE_BLACK
			} else {
				current_state = MOVE_BLACK
			}
		case MOVE_BLACK:
			coord := move(BLACK)
			cnt := check_line(coord.x, coord.y)
			if cnt > 0 {
				if cnt == 2 {
					clear()
					print(board)
					delete(WHITE)
				}
				current_state = DELETE_WHITE
			} else {
				current_state = MOVE_WHITE
			}
		}
		symbol := '_'
		fmt.Scanf("%c", &symbol)
		if symbol == 'q' {
			current_state = CLOSE
		}
		clear()
	}
}
