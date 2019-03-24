package sdk

func OpenURL(url string) {
	conn.WriteJSON(&sentEvent{
		Event: EventOpenURL,
		Payload: &openUrlPayload{url},
	})
}

func GetGlobalSettings() {
	conn.WriteJSON(&sentEvent{Event: EventGetGlobalSettings, Context: PluginUUID})
}

func SetGlobalSettings(v interface{}) {
	conn.WriteJSON(&sentEvent{Event: EventSetGlobalSettings, Context: PluginUUID, Payload: v})
}

func Log(message string) {
	conn.WriteJSON(&sentEvent{
		Event: EventLogMessage,
		Context: PluginUUID,
		Payload: &logMessagePayload{message},
	})
}

func SetTitle(context, title string, target int) {
	conn.WriteJSON(&sentEvent{
		Event: EventSetTitle,
		Context: context,
		Payload: &setTitlePayload{Title: title, Target: target},
	})
}

func SetState(context string, state int) {
	conn.WriteJSON(&sentEvent{
		Event: EventSetState,
		Context: context,
		Payload: &setStatePayload{state},
	})
}

func SetImage(context, imageData string, target int) {
	conn.WriteJSON(&sentEvent{
		Event: EventSetImage,
		Context: context,
		Payload: &setImagePayload{imageData, target},
	})
}

func ShowAlert(context string) {
	conn.WriteJSON(&sentEvent{
		Event: EventShowAlert,
		Context: context,
	})
}

func ShowOk(context string) {
	conn.WriteJSON(&sentEvent{
		Event: EventShowOK,
		Context: context,
	})
}