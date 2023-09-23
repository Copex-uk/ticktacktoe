//
// "git "
//

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// declair Constants

const vers int = 28

var numberofgames int = 1

// Start of functions

func wait(s time.Duration) {

	time.Sleep(s * time.Millisecond)
}
func scoreboard(ply1 string, ply2 string, score [4]int) {
	// scoreboard
	fmt.Print("\n\n")
	fmt.Printf("+%s+\n", strings.Repeat("-", 45))
	fmt.Printf("| %-20s | %-20s |\n", "Player", "Score Game["+strconv.Itoa(score[3])+"]")
	fmt.Printf("|%s|\n", strings.Repeat("-", 45))
	fmt.Printf("| %-20s | %-20d |\n", ply1, score[0])
	fmt.Printf("| %-20s | %-20d |\n", ply2, score[1])
	fmt.Printf("+%s+\n", strings.Repeat("-", 45))
	fmt.Printf("| %-20s | %-20d |\n", "Games Drawen ", score[2])
	fmt.Printf("|%s|\n", strings.Repeat("-", 45))

}

func drawboard(board []string) {
	// draw the board
	fmt.Print("\033c")
	fmt.Printf("Welcome to Tic Tac Toe ! ( Version  %04d)\n", vers)
	fmt.Print("\n\n")
	fmt.Printf(" %s| %s| %s\n", board[1], board[2], board[3])
	fmt.Printf("__|__|__\n")
	fmt.Printf(" %s| %s| %s\n", board[4], board[5], board[6])
	fmt.Printf("__|__|__\n")
	fmt.Printf(" %s| %s| %s\n", board[7], board[8], board[9])
	fmt.Print("\n\n")

}

func defaultboard(board []string) {
	// clear the board

	for i := 0; i < len(board); i++ {
		board[i] = strconv.Itoa(i)
	}
	board[0] = "-"
	drawboard(board)

}

func updateboard(updateboard []string, pmove int, xoro string) {
	// update the board[] to reflect the players move
	updateboard[pmove] = xoro
}

func checkwin(checkboard []string, player string) bool {

	// check to see if there are 3 consecutive X's or 0's and return true if found

	wins := [...]int{
		0,
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
		1, 4, 7,
		2, 5, 8,
		3, 6, 9,
		1, 5, 9,
		3, 5, 7,
	}

	i := 1

	for {

		if i >= len(wins) {
			return false
		}
		if checkboard[wins[i]] == player && checkboard[wins[i+1]] == player && checkboard[wins[i+2]] == player {
			return true
		}

		i = i + 3
	}

}
func checkdraw(checkboard []string) bool {

	// check if any of the board[] slice contain anything that is no a X or 0 we skip board[0] as it will always contain "-"
	// we retrun true if we get to the end of the list of false if we find any X's or 0'S befor we get to the end

	wins := [...]int{
		0,
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
		1, 4, 7,
		2, 5, 8,
		3, 6, 9,
		1, 5, 9,
		3, 5, 7,
	}

	i := 1

	for {

		if i >= len(wins) {
			return true
		}

		if checkboard[wins[i]] != "X" && checkboard[wins[i]] != "0" {
			return false
		}

		i++
	}

}

func gophermove() int {

	// check computer player and generate move

	wait(50)
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(9) + 1

}

func canwin(plmove string, checkboard []string) int {

	// check if the computer player has 2 of 3 in a wining line and return the empty space to win
	var oplayer string

	oplayer = ""

	if plmove == "X" {
		oplayer = "0"
	} else {
		oplayer = "X"
	}

	wins := [...]int{
		0,
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
		1, 4, 7,
		2, 5, 8,
		3, 6, 9,
		1, 5, 9,
		3, 5, 7,
	}

	i := 1

	check := [3]int{0, 0, 0}

	// code below can be optmised --->

	for {

		if checkboard[wins[i]] == plmove && checkboard[wins[i]] != oplayer {
			check[0] = 1
		}

		if checkboard[wins[i+1]] == plmove && checkboard[wins[i+1]] != oplayer {
			check[1] = 1
		}

		if checkboard[wins[i+2]] == plmove && checkboard[wins[i+2]] != oplayer {
			check[2] = 1
		}

		if check[0]+check[1]+check[2] == 2 {

			if check[0] == 0 && checkboard[wins[i]] != plmove && checkboard[wins[i]] != oplayer {
				return wins[i]
			}

			if check[1] == 0 && checkboard[wins[i+1]] != plmove && checkboard[wins[i+1]] != oplayer {
				return wins[i+1]
			}

			if check[2] == 0 && checkboard[wins[i+2]] != plmove && checkboard[wins[i+2]] != oplayer {
				return wins[i+2]
			}

		}
		check[0] = 0
		check[1] = 0
		check[2] = 0

		i = i + 3

		if i >= len(wins) {
			return 0
		}
	}
}

func playagain() bool {
	// check if there is another game
	//var playagain string
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n(Games Played %d) Would you Like to Play again? (y/n) ->", numberofgames)
	// fmt.Print("play again?")
	//fmt.Scanf("%s", &playagain)
	playagain, _ := reader.ReadString('\n')

	fmt.Println(len(playagain))

	if strings.TrimSpace(strings.ToLower(playagain)) == "y" {
		numberofgames = numberofgames + 1
		return true
	}

	return false
}

// Start of main game (maybe usefull to someone :-) ???

func main() {

	var input, player1, player2 string
	var pause time.Duration
	var score [4]int

	pause = 35
	score = [4]int{0, 0, 0, 0}

	board := []string{"-", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	// waste some time printing / - \ for no reason :-)

	fmt.Print("\033c")
	fmt.Printf("Welcome to Tic Tac Toe ! ( Version  %04d)\n", vers)
	fmt.Println()
	fmt.Print("Setting up the game ....")

	for a := 0; a <= 10; a++ {
		fmt.Print("/")
		wait(pause)
		fmt.Print("\b")
		fmt.Print("-")
		wait(pause)
		fmt.Print("\b")
		fmt.Print("\\")
		wait(pause)
		fmt.Print("\b")
	}

	fmt.Println()

	//fmt.Print("\n\nPlayer one allways go's first (use gopher for cumputer player)\n\n")
	// get palyers names
	fmt.Print("Player 1, What is your Name ? ")
	fmt.Scanf("%s", &player1)

	fmt.Print("Player 2, What is your Name ? ")
	fmt.Scanln(&player2)

	fmt.Printf("Welcome, %s and %s! Let's play!\n", player1, player2)

	// set current player name and current player move to track game, input set to 0 to stop game loop :-)
	// if player one is set to gopher

	plmove := "X"
	cplname := player1
	input = "0"
	computer := "gopher"

	drawboard(board)

	for {

		if cplname != computer {
			fmt.Printf("[%s] What is your move (1-9)? or (q to Quit) >", cplname)
			fmt.Scanln(&input)
		}

		if input == "q" || input == "Q" {
			break
		}

		value, error := strconv.Atoi(input)

		if error != nil && input != "q" {
			fmt.Printf("[%s]Can you not follow instructions ??\n", cplname)
		}

		if cplname == computer {

			canwin := canwin(plmove, board)

			if canwin == 0 {
				value = gophermove()
			} else {
				value = canwin
			}
		}

		if error == nil {

			if value >= 1 && value <= 9 {

				if board[value] != "X" && board[value] != "0" {

					updateboard(board, value, plmove)
					drawboard(board)

					fmt.Printf("[%s] Your move was :- %d \n\n", cplname, value)

					if checkwin(board, plmove) {
						fmt.Printf("Congratulations %s! You won this game.", cplname)

						score[3] = score[3] + 1

						if plmove == "X" {
							score[0] = score[0] + 1
						} else {
							score[1] = score[1] + 1
						}

						scoreboard(player1, player2, score)

						if playagain() {
							defaultboard(board)

						} else {
							fmt.Println("\a\nEnd of Game.......")
							break
						}

					} else if checkdraw(board) {
						fmt.Printf("Oh well %s the game with %s was a draw.....", player1, player2)

						score[2] = score[2] + 1

						scoreboard(player1, player2, score)

						if playagain() {
							defaultboard(board)

						} else {
							fmt.Println("\a\nEnd of Game.......")
							break
						}
					}

					// fmt.Printf("[%s] Your move was :- %d \n\n", cplname, value)

					if plmove == "X" {
						plmove = "0"
						cplname = player2
					} else {
						plmove = "X"
						cplname = player1
					}
				} else {
					fmt.Printf("[%s], The space [%d] is already occupied, try again .... \n", cplname, value)
				}

			} else {
				fmt.Printf("[%s], \a Number out of ramge enter a number between 1 and 9\n", cplname)
				wait(2000)
			}

		}

	}

}
