package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	playersInGame := 4
	cond := sync.NewCond(&sync.Mutex{})
	for playerID := 0; playerID < 4; playerID++ {
		go playerHandler(cond, &playersInGame, playerID)
		time.Sleep(1 * time.Second)
	}
}

func playerHandler(cond *sync.Cond, playersRemaining *int, playerID int) {
	cond.L.Lock()
	fmt.Println(playerID, ": Connected")
	*playersRemaining--
	if *playersRemaining == 0 {
		cond.Broadcast()
	}
	for *playersRemaining > 0 {
		fmt.Println(playerID, ": Waiting for more players")
		cond.Wait()
	}
	cond.L.Unlock()
	fmt.Println("All players connected. Ready player", playerID)
	//Game started
}
