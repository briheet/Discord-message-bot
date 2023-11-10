package main

import(
	"fmt"
	"github.com/briheet01/discord-bot/config"
	"github.com/briheet01/discord-bot/bot"
)

func main(){
	err := config.ReadConfig()

	if err != nil{
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
}