package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type board struct{
	val []string
	players int
	player1 string
	player2 string
}

func main() {
	fmt.Println("======= Tic Tac Toe! =======")
	fmt.Println("Let us play!")
	var b board
	
	b.init()
}

func (b *board) getUserPos(msg string) int{
	userInput := b.getInput(msg)
	// string to int
	userPos, err := strconv.Atoi(userInput)
	handleErr(err)
	if userPos < 0 || userPos > 10 {
		fmt.Println("Invalid answer...")
		userPos = b.getUserPos("Please try again: ")
	}

	if b.val[userPos - 1] == "X" || b.val[userPos - 1] == "O" {
		fmt.Println("Sorry, the position is already taken...")
		userPos = b.getUserPos("Please try again: ")
	} 
	
	return userPos
}

func (b *board) startDoublePlayer(){
	msg1 := "Player 1, select position for 'X'"
	msg2 := "Player 2, select position for 'O'"
	var msg string

	for i := 0; i < 9; i++ {
		b.draw()
		fmt.Println()
		if i % 2 == 0 {
			msg = msg1
		}else{
			msg = msg2
		}

		userPos := b.getUserPos(msg)

		if i % 2 == 0 {
			b.val[userPos - 1] =  b.player1
		}else{
			b.val[userPos - 1] =  b.player2
		}
		fmt.Println()

		if (b.checkForWin()){
			b.draw()
			if i % 2 == 0 {
				fmt.Println("======> Player 1 wins!! <=======")
			}else{
				fmt.Println("======> Player 2 wins!! <=======")
			}
			fmt.Println("Game over!")
			break
		}
	}

}

func (b *board) init(){
	b.player1 = "X"
	b.player2 = "O"
	b.val = append(b.val, "1", "2", "3", "4", "5", "6", "7", "8", "9")
	players := b.getInput("How many players? 1 or 2?")
	switch players {
	case "1":
		b.players = 1
		fmt.Println("You vs Computer!")
		break
	case "2":
		b.players = 2
		b.startDoublePlayer()
		break
	default:
		fmt.Println("Invalid answer, exiting...")
		break
	}
}

// Draw the board
func (b *board) draw(){
	fmt.Println()
	for i, val := range b.val {
		fmt.Print(val, " ")
		if i == 2 || i == 5 {
			fmt.Println("")
			fmt.Println("------")
		}
	}
	fmt.Println()
}

// Get input of the user
func (b *board)getInput(msg string) string{
	fmt.Print(msg, ":")
	var result string
	scanner := bufio.NewScanner(os.Stdin)
	for{
		scanner.Scan()
		input := scanner.Text()
		if len(input)!=0{
			// fmt.Println("received: ", input)
			result = input
			break
		}else{break}
	}
	return result
}

func handleErr(err error){
	if err != nil {
		log.Panic("err", err)
	}
}

func (b *board) checkForWin() bool {
	// check vertically
	if (b.val[0] == b.val[3] && b.val[3] == b.val[6]) ||
	(b.val[1] == b.val[4] && b.val[4] == b.val[7]) ||
	(b.val[2] == b.val[5] && b.val[5] == b.val[8])	{
		return true
	}
	
	// check horizontally
	if (b.val[0] == b.val[1] && b.val[1] == b.val[2]) || 
	 (b.val[3] == b.val[4] && b.val[4] == b.val[5]) || 
	 (b.val[6] == b.val[7] && b.val[7] == b.val[8]){
		return true
	}


	// check diagonally
	if (b.val[0] == b.val[4] && b.val[4] == b.val[8]) || 
	(b.val[2] == b.val[4] && b.val[4] == b.val[6]) {
	   return true
   }

	return false
}
