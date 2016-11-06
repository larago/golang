package main

import (
	"code.google.com/p/go.net/websocket"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	listenAddr = "localhost:9527" // server address
)

var (
	pwd, _        = os.Getwd()
	RootTemp      = template.Must(template.ParseFiles(pwd + "/chat.html"))
	JSON          = websocket.JSON              // codec for JSON
	Message       = websocket.Message           // codec for string, []by
	ActiveClients = make(map[ClientConn]string) // map containing clients
	User          = make(map[string]string)
)

// Initialize handlers and websocket handlers
func init() {
	User["aaa"] = "aaa"
	User["bbb"] = "bbb"
	User["test"] = "test"
	User["test2"] = "test2"
	User["test3"] = "test3"
}

//
