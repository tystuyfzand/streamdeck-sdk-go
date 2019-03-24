package sdk

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"os"
	"os/signal"
	"syscall"
)

var (
	flagPort = flag.Int("port", 0, "streamdeck sdk port")
	flagEvent = flag.String("registerEvent", "", "streamdeck sdk register event")
	flagInfo = flag.String("info", "", "streamdeck application info")

	PluginUUID string

	conn *websocket.Conn
	closeSdk chan struct{}
)

func init() {
	flag.StringVar(&PluginUUID, "pluginUUID", "", "streamdeck plugin uuid")
}

func Open() (error) {
	flag.Parse()

	closeSdk = make(chan struct{})

	c, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://localhost:%d", *flagPort), nil)

	if err != nil {
		return err
	}

	conn = c

	go reader()

	c.WriteJSON(&openMessage{Event: *flagEvent, UUID: PluginUUID})

	return nil
}

// Waits for either the socket to close, or SIGINT/SIGTERM.
// Stream Deck closes the connection as a way to close the plugin,
// but will also SIGTERM/SIGINT as a final resort.
func Wait() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <- closeSdk:
	case <- sigs:
		break
	}
}