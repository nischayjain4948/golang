package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Player struct {
	health int
}

func takeDamageFromExplosion(player *Player) {
	fmt.Println("Player is taking damage from explossion")
	explosionDamage := 12
	player.health -= explosionDamage
}

func getHealth(player *Player) {
	fmt.Println("Player will get the health from")
	player.health += 20
}

func comeFlyWithMe() {
	fmt.Println("Come On, come fly with me!")
}

func main() {

	fmt.Println("Welcome to the game of healUp and healDown")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Your code: ")
	input, _ := reader.ReadString('\n')
	trimmedInput := strings.TrimSpace(input)
	myCase, _ := strconv.ParseInt(trimmedInput, 10, 32)
	fmt.Println(myCase)

	playerHealth := Player{
		health: 100,
	}
	switch myCase {
	case 1:
		getHealth(&playerHealth)
		fmt.Println("After Healup the health of Player is: ", playerHealth.health)

	case 2:
		takeDamageFromExplosion(&playerHealth)
		fmt.Println("After explossion the health is: ", playerHealth.health)

	case 3:
		comeFlyWithMe()

	default:
		fmt.Println("Before Explosion the health is: ", playerHealth.health)

	}

}
