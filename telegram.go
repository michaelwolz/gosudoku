package gosudoku

import (
	"fmt"
	"github.com/yanzay/tbot"
	"log"
	"strconv"
)

const GROUPID = -237105392

var bot *tbot.Server

var connected = make(chan bool)

func StartTelegramBot(token string) {
	var err error

	log.Println("Initializing Telegram Bot!")

	// Create new telegram bot server using token
	bot, err = tbot.NewServer(token)
	checkErr(err)

	bot.HandleFunc("/fieldconfig {fieldconfig}", FieldConfigHandler)
	bot.HandleDefault(DefaultHandler)

	err = bot.Send(GROUPID, "Let's get this puzzle solved ;)")
	checkErr(err)

	connected <- true
	err = bot.ListenAndServe()
	checkErr(err)
}

func FieldConfigHandler(message *tbot.Message) {
	fieldConfig := message.Vars["fieldconfig"]
	readFieldConfiguration(fieldConfig)
}

func DefaultHandler(message *tbot.Message) {
	message.Reply("I don't know, what you want from me :(")
}

func sendMessage(message string) {
	err := bot.Send(GROUPID, message)
	checkErr(err)
}

func sendFieldConfiguration() {
	<-connected
	conf := myBox.id
	for key, val := range myBox.values {
		if val != 0 {
			x, y := getCoordinatesForIndex(key)
			conf += "," + strconv.Itoa(x) + strconv.Itoa(y) + ":" + strconv.Itoa(val)
		}
	}
	sendMessage("/fieldconfig " + conf)
}

func readFieldConfiguration(fieldConfig string) {
	fmt.Println(fieldConfig)
}
