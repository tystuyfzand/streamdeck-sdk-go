package main

import (
	"github.com/valyala/fastjson"
	"log"
	"meow.tf/streamdeck/sdk"
)

func main() {
	// Initialize handlers for events
	sdk.RegisterAction("tf.meow.example.doSomething", doSomethingHandler)

	// Open and connect the SDK
	err := sdk.Open()

	if err != nil {
		log.Fatalln(err)
	}

	// Wait until the socket is closed, or SIGTERM/SIGINT is received
	sdk.Wait()
}

func doSomethingHandler(action, context string, payload *fastjson.Value, deviceId string) {
	// Do something as a result of an action (keyDown)
}