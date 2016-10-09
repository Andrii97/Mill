package main

import (
	"fmt"
	//"io/ioutil"
	//"log"
)

type cell_state int

var board [n][n]cell_state

const (
	EMPTY        cell_state = 0
	WHITE        cell_state = 5
	BLACK        cell_state = -5
	INACCESSIBLE cell_state = -1 // Not available
	BORDER       cell_state = -2
)

const n = 9 // board_size

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

func main() {
	init_board()
	/*board[4][5] = 5
	board[4][6] = 5
	board[4][7] = 5
	board[4][1] = 5
	board[4][2] = 5
	board[4][3] = 5
	board[3][3] = 5
	board[5][3] = 5
	board[7][1] = 5
	board[1][1] = 5
	board[1][4] = 5
	board[1][7] = 5
	/*symbol := '_'
	for symbol != 'q' {
		print(board)
		fmt.Scanf("%c", &symbol) // _, err :=
	}*/
	print(board)
	fmt.Print(check_line(1, 1))
}
