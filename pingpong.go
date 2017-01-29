package main

import "fmt"
import "time"
import "sync"
import "math/rand"

var wg sync.WaitGroup
var vollies = 0
var max = 17
var min = 5

func init() {
	rand.Seed(time.Now().Unix())
	vollies = rand.Intn(max-min) + min
}

func main() {

	messages := make(chan string)
	wg.Add(1)
	go func() {
		defer wg.Done()
		i := 0
		for {
			select {

			case msg := <-messages:
				fmt.Println("Pong")
				if i > vollies {
					fmt.Println("Game Over ! I win Mr. Ping")
					messages <- "done"
					return
				} else if msg == "done" {
					return
				}
				messages <- "Pong"
				i++
			case <-time.After(time.Second):
				fmt.Println("You are slow. Let's get this game going!!! Waiting for Mr. Ping")
			}
		}
	}()

	//Waiting here for timeout in pong function
	//
	time.Sleep(time.Second * 1)
	fmt.Println("Ping")
	messages <- "Ping"
	time.Sleep(time.Second * 2)
	for {
		select {
		case msg := <-messages:
			if msg == "Pong" {
				fmt.Println("Ping")
				messages <- "Ping"
			} else if msg == "done" {
				fmt.Println("Will get you next time Mr. Pong")
				wg.Wait()
				return
			}
		case <-time.After(time.Second):
			fmt.Println("Waiting for Mr.  Pong")
		}
	}
}
