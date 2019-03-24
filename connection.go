package sdk

import (
	"github.com/valyala/fastjson"
)

func reader() {
	var p fastjson.Parser

	var event, context, action, deviceId string

	var payload *fastjson.Value

	for {
		_, message, err := conn.ReadMessage()

		if err != nil {
			close(closeSdk)
			return
		}


		v, err := p.ParseBytes(message)

		if err != nil {
			continue
		}

		event = JsonStringValue(v, CommonEvent)
		context = JsonStringValue(v, CommonContext)
		action = JsonStringValue(v, CommonAction)
		deviceId = JsonStringValue(v, CommonDevice)

		payload = v.Get(CommonPayload)

		switch event {
		case EventKeyDown:
			handleEvent(EventKeyDown, &KeyDownEvent{action, context, payload, deviceId})
		case EventKeyUp:
			handleEvent(EventKeyUp, &KeyUpEvent{action, context, payload, deviceId})
		case EventWillAppear:
			handleEvent(EventWillAppear, &WillAppearEvent{action, context, payload, deviceId})
		case EventWillDisappear:
			handleEvent(EventWillDisappear, &WillDisappearEvent{action, context, payload, deviceId})
		case EventDeviceDidConnect:
			handleEvent(EventDeviceDidConnect, &DeviceConnectEvent{deviceId, v.Get(CommonDeviceInfo)})
		case EventDeviceDidDisconnect:
			handleEvent(EventDeviceDidDisconnect, &DeviceDisconnectEvent{deviceId})
		case EventSendToPlugin:
			handleEvent(EventSendToPlugin, &SendToPluginEvent{action, context, payload, deviceId})
		case EventDidReceiveGlobalSettings:
			handleEvent(EventDidReceiveGlobalSettings, &GlobalSettingsEvent{payload.Get("settings")})
		case EventApplicationDidLaunch:
			handleEvent(EventApplicationDidLaunch, &ApplicationLaunchEvent{JsonStringValue(payload, "application")})
		case EventApplicationDidTerminate:
			handleEvent(EventApplicationDidTerminate, &ApplicationTerminateEvent{JsonStringValue(payload, "application")})
		}
	}
}