package sdk

import (
	"github.com/valyala/fastjson"
)

var (
	actionMap = make(map[string]ActionHandler)
)

func RegisterAction(actionName string, f ActionHandler) {
	actionMap[actionName] = f
}

func handleAction(action, context string, payload *fastjson.Value, deviceId string) {
	if handler, ok := actionMap[action]; ok {
		handler(action, context, payload, deviceId)
	}
}