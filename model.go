package sdk

import "github.com/valyala/fastjson"

// Sent on socket open to identify the plugin/register it
type openMessage struct {
	Event string `json:"event"`
	UUID  string `json:"uuid"`
}

// Sent to the Stream Deck software as a generic event
type sentEvent struct {
	Event   string      `json:"event"`
	Context string      `json:"context,omitempty"`
	Payload interface{} `json:"payload"`
}

// Payloads for sent events

type openUrlPayload struct {
	URL string `json:"url"`
}

type logMessagePayload struct {
	Message string `json:"message"`
}

type setTitlePayload struct {
	Title  string `json:"title"`
	Target int    `json:"target"`
}

type setStatePayload struct {
	State int `json:"state"`
}

type setImagePayload struct {
	Image  string `json:"image"`
	Target int    `json:"target"`
}

// Events

type KeyDownEvent struct {
	Action   string
	Context  string
	Payload  *fastjson.Value
	DeviceId string
}

type KeyUpEvent struct {
	Action   string
	Context  string
	Payload  *fastjson.Value
	DeviceId string
}

type WillAppearEvent struct {
	Action   string
	Context  string
	Payload  *fastjson.Value
	DeviceId string
}

type WillDisappearEvent struct {
	Action   string
	Context  string
	Payload  *fastjson.Value
	DeviceId string
}

type DeviceConnectEvent struct {
	DeviceId string
	Info     *fastjson.Value
}

type DeviceDisconnectEvent struct {
	DeviceId string
}

type SendToPluginEvent struct {
	Action   string
	Context  string
	Payload  *fastjson.Value
	DeviceId string
}

type ReceiveSettingsEvent struct {
	Action   string
	Context  string
	DeviceId string
	Settings *fastjson.Value
}

type GlobalSettingsEvent struct {
	Settings *fastjson.Value
}

type ApplicationLaunchEvent struct {
	Application string
}

type ApplicationTerminateEvent struct {
	Application string
}
