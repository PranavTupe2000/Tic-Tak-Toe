package main

import (
	"fmt"
)

type ticTacToe struct {
	position []string
	player   string
}

func (t *ticTacToe) inti() {
	for i := 0; i < 9; i++ {
		t.position[i] = string(i + 49)
	}
}

func (t ticTacToe) printBoard() {
	println("\n")
	for i := 0; i < 9; i = i + 3 {
		println(t.position[i], "|", t.position[i+1], "|", t.position[i+2])
		if i < 6 {
			println("---------")
		}
	}
}

func (t *ticTacToe) swapPlayer() {
	if t.player == "X" {
		t.player = "O"
	} else {
		t.player = "X"
	}
}

func (t *ticTacToe) play() {
	var i int
	fmt.Print("\nPlayer ", t.player, " select your choice: ")
	fmt.Scan(&i)
	for {
		if t.position[i-1] != "X" && t.position[i-1] != "O" {
			t.position[i-1] = t.player
			break
		} else {
			fmt.Print("Invalid Choice!!! Enter again: ")
			fmt.Scan(&i)
		}
	}
}

func (t ticTacToe) checkWinner() string {
	var winner string = ""
	switch {
	case t.position[0] == t.position[1] && t.position[1] == t.position[2]: //123
		winner = t.player
	case t.position[3] == t.position[4] && t.position[4] == t.position[5]: //456
		winner = t.player
	case t.position[6] == t.position[7] && t.position[7] == t.position[8]: //789
		winner = t.player
	case t.position[0] == t.position[3] && t.position[3] == t.position[6]: //147
		winner = t.player
	case t.position[1] == t.position[4] && t.position[4] == t.position[7]: //258
		winner = t.player
	case t.position[2] == t.position[5] && t.position[5] == t.position[8]: //369
		winner = t.player
	case t.position[0] == t.position[4] && t.position[4] == t.position[8]: //159
		winner = t.player
	case t.position[2] == t.position[4] && t.position[4] == t.position[6]: //357
		winner = t.player
	}
	return winner
}

func main() {
	t := ticTacToe{make([]string, 9), "X"}
	t.inti()

	for i := 0; i < 9; i++ {
		t.printBoard()
		t.play()
		if "" == t.checkWinner() {
			t.swapPlayer()
		} else {
			t.printBoard()
			fmt.Println("\nPlayer ", t.player, " is the Winner")
			break
		}
	}
	if "" == t.checkWinner() {
		t.printBoard()
		fmt.Println("\nIt's a draw")
	}

}
