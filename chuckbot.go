package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/thoj/go-ircevent"
)

var j = make(chan string)
var con *irc.Connection

var server string
var channel string
var nick string

func init() {
	flag.StringVar(&nick, "nick", "chuckbot", "the nick for the bot")
	flag.StringVar(&channel, "channel", "", "the channel to join")
	flag.StringVar(&server, "server", "", "the <server-ip>:<port> to connect to")

	flag.Parse()
	
	if server == "" {
		panic("no server defined. use flag '-server <host>:<port>' to set it.")
	}

	if channel == "" {
		panic("no channel defined. use flag '-channel <channelname>' to set it.")
	}
}

func main() {
	con = irc.IRC(nick, nick)
	con.Debug = true
	con.UseTLS = true
	con.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	fmt.Printf("%v\n", con)
	err := con.Connect(server)
	if err != nil {
		panic(err)
	}
	defer con.Quit()

	con.AddCallback("001", func(e *irc.Event) { con.Join(channel) })
	con.AddCallback("PRIVMSG", func(e *irc.Event) { con.Privmsg(e.Arguments[0], fmt.Sprintf("Apropos, %s: %s", e.Nick, <-j)) })

	go getQuote()
	con.Loop()
}

type J struct {
	Id         uint
	Joke       string
	Categories []string
}

type Q struct {
	Type  string
	Value J
}

func getQuote() {
	for {
		resp, err := http.Get("http://api.icndb.com/jokes/random?firstName=Chuckbot&amp;lastName=Norris")
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
		quote := Q{}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
		fmt.Printf("%v\n", body)
		err = json.Unmarshal(body, &quote)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
		resp.Body.Close()
		j <- quote.Value.Joke
	}
}
