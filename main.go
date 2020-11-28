package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

var winningconditions = [][]int{
	[]int{0, 1, 2},
	[]int{0, 3, 6},
	[]int{0, 4, 8},
	[]int{1, 4, 7},
	[]int{2, 4, 6},
	[]int{2, 5, 8},
	[]int{3, 4, 6},
	[]int{6, 7, 8},
}

// ok ok ill clean it up later
func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	quit := make(chan struct{})
	place := make(chan int)
	playing := false
	board := []string{"_", "_", "_", "_", "_", "_", "_", "_", "_"}
	player := "X"
	clear()
	fmt.Println("TicTacGo - Press Space to start!")

	go func() {
		for {
			ev := termbox.PollEvent()
			switch ev.Type {
				case termbox.EventKey:
					switch ev.Key {
						case termbox.KeyEsc:
							close(quit)
							return
						case termbox.KeySpace:
							if playing == false { 
								playing = true
								clear()
								displayboard(board)
							}
					}
					if ch := ev.Ch; ch > 48 && ch < 58 {
						place <- int(ch - 49)
					}
			}
		}
	}()

loop:
	for {
		select {
			case <-quit:
				break loop
			case p := <-place:
				if playing == true {
					if player == "X" { player = "O" } else { player = "X" }
					if state := play(board, p, player); state == "w" {
						fmt.Printf("\n%s won! Press Space to restart.", player)
						playing = false
						board = []string{"_", "_", "_", "_", "_", "_", "_", "_", "_"}
					}
				}
		}
	}
	termbox.Close()
	fmt.Printf("\x1b[2J")
	fmt.Println("Goodbye!")
}

func play(board []string, place int, player string) string {
	board[place] = player
	displayboard(board)
	for i := 0; i < 8; i++ {
		condition := winningconditions[i];
		a := board[condition[0]];
		b := board[condition[1]];
		c := board[condition[2]];

		if a == "_" || b == "_" || c == "_" { continue }
		if a == b && b == c { return "w" }
	}
	return "p"
}

func displayboard(board []string) {
	clear()
	for i := 0; i < 9; i++ {
		if i != 0 && i % 3 == 0 { fmt.Printf("\n") }
		fmt.Printf("%s ", board[i])
	}
}

func clear() {
	fmt.Printf("\x1b[2J")
	termbox.SetCursor(0, 0)
}