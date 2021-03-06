package main

import (
	"os"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

type options struct {
	AccountSid string
	AuthToken  string
	Receiver   string
	Sender     string
}

func main() {
        if len(os.Args) != 5 {
            fmt.Println("Usage:", os.Args[0], "PORT", "SID", "TOKEN", "SENDER")
            return
        }
        
	Port := ":" + os.Args[1]

	opts := options{
        	AccountSid: os.Args[2],
        	AuthToken: os.Args[3],
        	Sender: os.Args[4],
		Receiver:   os.Getenv("RECEIVER"),
	}

	if opts.AccountSid == "" || opts.AuthToken == "" || opts.Sender == "" {
		log.Fatal("'SID', 'TOKEN' and 'SENDER' need to be set")
	}

	o := NewMOptionsWithHandler(&opts)
        err := fasthttp.ListenAndServe(Port, o.HandleFastHTTP)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
