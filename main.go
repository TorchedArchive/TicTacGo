package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

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
					play(board, p, player)
				}
		}
	}
	termbox.Close()
	fmt.Printf("\x1b[2J")
	fmt.Println("Goodbye!")
}

func play(board []string, place int, player string) {
	board[place] = player
	displayboard(board)
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