package main

import (
	"fmt"
	//"io/ioutil"
	//"log"
)


type cell_state int 

const (
	EMPTY cell_state = 0
	WHITE cell_state = 5
	BLACK cell_state = -5
	INACCESSIBLE cell_state = -1 // Not available
	BORDER cell_state = -2
)

const n = 9 // board_size

func init_board() [n][n]cell_state{
	var board [n][n]cell_state
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (i < 1 || j < 1 || i > n - 2 || j > n - 2){
				board[i][j] = BORDER
			} else if (i == j || i == n - j - 1 || j == n / 2 || i == n / 2) {
				board[i][j] = EMPTY
			} else {
				board[i][j] = INACCESSIBLE
			}
		}
	}
	board[n/2][n/2] = INACCESSIBLE
	return board
}

func print(arr [n][n]cell_state){
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%4d", arr[i][j])
		}
		fmt.Println()
		fmt.Println()
	}
}

func main(){
	board := init_board()
	symbol := '_'
	for symbol != 'q'{
		print(board)
		fmt.Scanf("%c", &symbol) // _, err := 
	}
}