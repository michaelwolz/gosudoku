package gosudoku

import (
	"github.com/yanzay/tbot"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const GROUPID = -237105392 // GROUP CHAT ID

var foreignBoxRegEx = regexp.MustCompile(`^(BOX_[A,D,G][1,4,7]),(([0-2][0-2]:[1-9],?)+)$`)

var bot *tbot.Server

var connectedCh = make(chan bool)
var connected = false

func StartTelegramBot(token string) {
	var err error

	log.Println("Initializing Telegram Bot!")

	// Create new telegram bot server using token
	bot, err = tbot.NewServer(token)
	checkErr(err)

	bot.HandleFunc("/fieldconfig {fieldconfig}", FieldConfigHandler)
	bot.HandleDefault(DefaultHandler)

	err = bot.ListenAndServe()
	checkErr(err)

	err = bot.Send(GROUPID, "Let's get this puzzle solved ;)")
	checkErr(err)

	connectedCh <- true
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
	if !connected {
		<-connectedCh // Wait until Server is started
		connected = true
	}

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
	var isColumn = false
	var isRow = false

	if foreignBoxRegEx.MatchString(fieldConfig) {
		matches := foreignBoxRegEx.FindStringSubmatch(fieldConfig)
		foreignBoxID := matches[1]
		config := strings.Split(matches[2], ",")

		log.Println("Received fieldConfiguration from " + foreignBoxID)

		if foreignBoxID[:len(matches[1])-1] == myBox.id[:len(myBox.id)-1] {
			isColumn = true
		} else if foreignBoxID[len(matches[1])-1:] == myBox.id[len(myBox.id)-1:] {
			isRow = true
		}

		if isRow || isColumn { // Otherwise it's a box which doesn't affect us.
			for _, el := range config {
				x, y, v := readFieldConfigStr(el)
				if isColumn {
					myBox.setColValue(x, v)
				} else {
					myBox.setRowValue(y, v)
				}
			}
		}
	}
}

func drawResultBox() {
	err := bot.Send(GROUPID, myBox.getResultBoxString())
	checkSoftErr(err)
}
